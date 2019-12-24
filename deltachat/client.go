package deltachat

// #include <deltachat.h>
import "C"
import (
	"fmt"
	"log"
)

type ClientEventHandler func(context *Context, event *Event)

type Client struct {
	context           *Context
	eventChan         chan *Event
	eventReceiverQuit chan struct{}
	handlerMap        map[int]ClientEventHandler
}

func (c *Client) On(event int, handler ClientEventHandler) {
	if c.handlerMap == nil {
		c.handlerMap = make(map[int]ClientEventHandler)
	}

	c.handlerMap[event] = handler
}

func (c *Client) queueEvent(event int, data1 C.uintptr_t, data2 C.uintptr_t) int {
	data1Wrapper := NewData1(event, data1)
	data2Wrapper := NewData2(event, data2)

	c.eventChan <- &Event{
		EventType: event,
		Data1:     *data1Wrapper,
		Data2:     *data2Wrapper,
	}

	return 0
}

// Goroutine that listens for incoming events. Should be started for callbacks to be
// executed.
func (c *Client) startEventReceiver() {
	go func() {
		if c.eventChan == nil {
			c.eventChan = make(chan *Event)
		}

		c.eventReceiverQuit = make(chan struct{})

		for {
			select {
			case <-c.eventReceiverQuit:
				log.Println("Quitting event receiver")
				return
			case event := <-c.eventChan:
				go c.handleEvent(event)
			}
		}
	}()
}

func (c *Client) stopEventReceiver() {
	close(c.eventReceiverQuit)
}

// Default error handler
func handleError(event *Event) {
	log.Println(dcErrorString(event))
}

func dcErrorString(event *Event) string {
	name := eventNames[event.EventType]

	str, err := event.Data2.String()

	if err != nil {
		log.Println(
			fmt.Sprintf(
				"Unexpected data type while handeling %s:",
				name,
				err.Error(),
			),
		)

		return ""
	}

	return fmt.Sprintf("%s: %s", name, *str)
}

func (c *Client) handleEvent(event *Event) {
	eventType := event.EventType

	handler, ok := c.handlerMap[eventType]

	if !ok {
		if (EVENT_TYPES_ERROR&eventType) == eventType || eventType == DC_EVENT_WARNING {
			handleError(event)
			return
		}

		log.Printf("Got unhandled event: %s", eventNames[eventType])

		return
	}

	handler(c.context, event)
}

func (c *Client) Open(dbLocation string) {
	context := NewContext()

	context.StartWorkers()
	context.Open(dbLocation)

	context.SetHandler(c.queueEvent)
	c.context = context

	c.startEventReceiver()
}

func (c *Client) Configure() {
	(*c.context).Configure()
}

func (c *Client) SetConfig(key string, value string) {
	(*c.context).SetConfig(key, value)
}

func (c *Client) Context() *Context {
	return c.context
}

func (c *Client) GetConfig(key string) string {
	return (*c.context).GetConfig(key)
}

func (c *Client) Close() {
	c.stopEventReceiver()
	(*c.context).Close()
	(*c.context).Unref()
}
