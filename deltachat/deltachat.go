package deltachat

// #cgo CFLAGS: -I../deltachat-ffi/include
// #cgo LDFLAGS: -L../deltachat-ffi/lib -ldeltachat
// #include <deltachat.h>
// #include <godeltachat.h>
import "C"

import (
	"fmt"
	"sync"
)

const DC_LP_AUTH_NORMAL int = C.DC_LP_AUTH_NORMAL
const DC_LP_AUTH_OAUTH2 int = C.DC_LP_AUTH_OAUTH2

const DC_LP_IMAP_SOCKET_PLAIN int = C.DC_LP_IMAP_SOCKET_PLAIN
const DC_LP_IMAP_SOCKET_SSL int = C.DC_LP_IMAP_SOCKET_SSL
const DC_LP_IMAP_SOCKET_STARTTLS int = C.DC_LP_IMAP_SOCKET_STARTTLS

const DC_LP_SMTP_SOCKET_PLAIN int = C.DC_LP_SMTP_SOCKET_PLAIN
const DC_LP_SMTP_SOCKET_SSL int = C.DC_LP_SMTP_SOCKET_SSL
const DC_LP_SMTP_SOCKET_STARTTLS int = C.DC_LP_SMTP_SOCKET_STARTTLS

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
	C.dc_set_config(c.context, C.CString(key), C.CString(value))
}

func (c *Client) GetConfig(key string) string {
	return C.GoString(C.dc_get_config(c.context, C.CString(key)))
}

func (c *Client) Open(databaseLocation string) {
	C.dc_open(c.context, C.CString(databaseLocation), nil)
}

func (c *Client) handleEvent(
	context *C.dc_context_t,
	event C.int,
	data1 C.uintptr_t,
	data2 C.uintptr_t,
) C.uintptr_t {
	fmt.Println("I was Called!!!")
	return 0
}
