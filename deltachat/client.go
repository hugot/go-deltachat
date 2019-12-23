package deltachat

// #include <deltachat.h>
import "C"
import (
	"fmt"

	"github.com/labstack/gommon/log"
)

type ClientEventHandler func(context *Context, event *Event)

type Client struct {
	context           *Context
	rawEventChan      chan *rawEvent
	eventReceiverQuit chan struct{}
	handlerMap        map[int]ClientEventHandler
}

type rawEvent struct {
	EventType int
	Data1     C.uintptr_t
	Data2     C.uintptr_t
}

func (c *Client) On(event int, handler ClientEventHandler) {
	if c.handlerMap == nil {
		c.handlerMap = make(map[int]ClientEventHandler)
	}

	c.handlerMap[event] = handler
}

func (c *Client) queueEvent(event int, data1 C.uintptr_t, data2 C.uintptr_t) uint {
	c.rawEventChan <- &rawEvent{
		EventType: event,
		Data1:     data1,
		Data2:     data2,
	}

	return 0
}

// Goroutine that listens for incoming raw events. Should be started for callbacks to be
// executed.
func (c *Client) startEventReceiver() {
	go func() {
		if c.rawEventChan == nil {
			c.rawEventChan = make(chan *rawEvent)
		}

		c.eventReceiverQuit = make(chan struct{})

		for {
			select {
			case <-c.eventReceiverQuit:
				return
			case event := <-c.rawEventChan:
				c.handleEvent(event.EventType, event.Data1, event.Data2)
			}
		}
	}()
}

func (c *Client) stopEventReceiver() {
	close(c.eventReceiverQuit)
}

// Default error handler
func handleError(event *Event) {
	log.Error(dcErrorString(event))
}

func dcErrorString(event *Event) string {
	name := errorTypeNames[event.EventType]

	str, err := event.Data2.String()

	if err != nil {
		log.Error(
			fmt.Sprintf(
				"Unexpected data type while handeling %s:",
				name,
				err.Error(),
			),
		)

		return ""
	}

	return fmt.Sprintf("%s: %s", name, str)
}

func (c *Client) handleEvent(event int, data1 C.uintptr_t, data2 C.uintptr_t) {
	eventStruct := &Event{
		EventType: event,
		Data1: Data{
			DataType: Data1TypeForEvent(event),
			data:     data1,
		},
		Data2: Data{
			DataType: Data2TypeForEvent(event),
			data:     data2,
		},
	}

	handler, ok := c.handlerMap[event]

	if !ok {
		if (EVENT_TYPES_ERROR & event) == event {
			go handleError(eventStruct)
		}

		return
	}

	go handler(c.context, eventStruct)
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
	(*c.context).Close()
	(*c.context).Unref()

	c.stopEventReceiver()
}
