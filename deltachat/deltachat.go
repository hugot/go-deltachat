package deltachat

// #cgo CFLAGS: -I../deltachat-ffi/include
// #cgo LDFLAGS: -L../deltachat-ffi/lib -ldeltachat -ldl -lm
// #include <deltachat.h>
// #include <godeltachat.h>
// #include <stdlib.h>
import "C"

import (
	"sync"
	"unsafe"
)

var deltachatCbMutex sync.RWMutex

type deltachatCallback func(
	context *C.dc_context_t,
	event C.int,
	data1 C.uintptr_t,
	data2 C.uintptr_t,
) C.uintptr_t

var deltachatCallbacks map[*C.dc_context_t]deltachatCallback = make(
	map[*C.dc_context_t]deltachatCallback,
)

//export godeltachat_eventhandler_proxy
func godeltachat_eventhandler_proxy(
	context *C.dc_context_t,
	event C.int,
	data1 C.uintptr_t,
	data2 C.uintptr_t,
) C.uintptr_t {
	deltachatCbMutex.RLock()

	callback, ok := deltachatCallbacks[context]

	deltachatCbMutex.RUnlock()

	if !ok {
		panic("dc_context_t callback was called but not set")
	}

	return callback(context, event, data1, data2)
}

func NewClient() *Client {
	context := C.godeltachat_create_context()

	client := &Client{
		context: context,
	}

	deltachatCbMutex.Lock()
	deltachatCallbacks[context] = client.handleEvent
	deltachatCbMutex.Unlock()

	return client
}

type Client struct {
	context *C.dc_context_t
}

func (c *Client) SetConfig(key string, value string) {
	cKey, cValue := C.CString(key), C.CString(value)
	C.dc_set_config(c.context, cKey, cValue)
	freeCString(cKey, cValue)
}

func (c *Client) GetConfig(key string) string {
	cConfigValue := C.dc_get_config(c.context, C.CString(key))
	configValue := C.GoString(cConfigValue)
	freeCString(cConfigValue)

	return configValue
}

func (c *Client) Open(databaseLocation string) {
	cDatabaseLocation := C.CString(databaseLocation)
	C.dc_open(c.context, cDatabaseLocation, nil)
	freeCString(cDatabaseLocation)
}

func (c *Client) Configure() {
	C.dc_configure(c.context)
}

func (c *Client) handleEvent(
	context *C.dc_context_t,
	event C.int,
	data1 C.uintptr_t,
	data2 C.uintptr_t,
) C.uintptr_t {
	return 0
}

func (c *Client) imapRoutine() {
	for {
		C.dc_perform_imap_jobs(c.context)
		C.dc_perform_imap_fetch(c.context)
		C.dc_perform_imap_idle(c.context)
	}
}

func (c *Client) smtpRoutine() {
	for {
		C.dc_perform_smtp_jobs(c.context)
		C.dc_perform_smtp_idle(c.context)
	}
}

func (c *Client) StartWorkers() {
	go c.imapRoutine()
	go c.smtpRoutine()
}

func (c *Client) CreateChatByContactID(ID uint32) uint32 {
	return uint32(C.dc_create_chat_by_contact_id(c.context, C.uint32_t(ID)))
}

func (c *Client) SendTextMessage(chatID uint32, message string) {
	cMessage := C.CString(message)
	C.dc_send_text_msg(c.context, C.uint32_t(chatID), cMessage)
	freeCString(cMessage)
}

func freeCString(strings ...*C.char) {
	for _, s := range strings {
		C.free(unsafe.Pointer(s))
	}
}

func (c *Client) CreateContact(name *string, address *string) uint32 {
	var nameString *C.char

	if name != nil {
		nameString = C.CString(*name)
	}

	var addressString *C.char

	if address != nil {
		addressString = C.CString(*address)
	}

	contactID := C.dc_create_contact(c.context, nameString, addressString)

	freeCString(nameString, addressString)

	return uint32(contactID)
}
