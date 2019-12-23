# go-deltachat
This is go-deltachat, an implementation of go bindings for the C FFI provided with
[deltachat-core-rust](https://github.com/deltachat/deltachat-core-rust). It links with C
through cgo and aims to wrap the entire API that deltachat exposes. All datatypes are
wrapped in go datatypes and all methods are wrapped in go functions. Efford has been made
to abstract away as much memory management as possible.

## Usage example
This library wraps the entire FFI API that deltachat has to offer and can of course be
used however you please. The recommended way is to use deltachat.Client, which is a thin
layer on top of deltachat's context datatype. It will make handeling events that are
emitted by deltachat easier and removes the need for some boilerplate code.

Example:

```go
package main

import (
	"fmt"

	"github.com/hugot/go-deltachat/deltachat"
	"github.com/labstack/gommon/log"
)

func main() {
	client := &deltachat.Client{}

	client.Open("/var/lib/deltabot/stuff.db")

    // Bear in mind that the config parameters are stored in the sqlite database
    // So these config values do not need to be set during each run. They're just here
    // to have a working example.
	client.SetConfig("addr", "test@example.com")
	client.SetConfig("mail_pw", "secret password")
	client.SetConfig("mail_server", "imap.example.com")
	client.SetConfig("mail_user", "test@example.com")
	client.SetConfig("mail_port", "993")

	client.SetConfig("send_pw", "secret password")
	client.SetConfig("send_server", "smtp.example.com")
	client.SetConfig("send_user", "test@example.com")
	client.SetConfig("send_port", "587")

	client.SetConfig(
		"server_flags",
		fmt.Sprintf(
			"%d",
			deltachat.DC_LP_AUTH_NORMAL|
				deltachat.DC_LP_IMAP_SOCKET_SSL|
				deltachat.DC_LP_SMTP_SOCKET_STARTTLS,
		),
	)

	client.Configure()

	addr := "chat@example.com"
    
   	client.On(deltachat.DC_EVENT_IMAP_CONNECTED, func(c *deltachat.Context, e *deltachat.Event) {
		contactID := c.CreateContact(nil, &addr)
		chatID := c.CreateChatByContactID(contactID)

		c.SendTextMessage(chatID, "Hello World!")

		log.Info("Sent message!")
	})

	// Handler for info logs from libdeltachat
	client.On(deltachat.DC_EVENT_INFO, func(c *deltachat.Context, e *deltachat.Event) {
		info, _ := e.Data2.String()

		log.Info(*info)
	})

    // Prevent the program from exiting
	wait := make(chan struct{})

	for {
		<-wait
	}
}
```

## Documentation
Since the API is pretty much the go-equivalent of the C API, it should suffice to read the
documentation at [c.delta.chat](https://c.delta.chat) for most datatypes and functions.
