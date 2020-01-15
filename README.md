# go-deltachat
This is go-deltachat, an implementation of go bindings for the C FFI provided with
[deltachat-core-rust](https://github.com/deltachat/deltachat-core-rust). It links with C
through cgo and aims to wrap the entire API that deltachat exposes. All datatypes are
wrapped in go datatypes and all methods are wrapped in go functions. Efford has been made
to abstract away as much memory management as possible.

## Installation

```bash
go get github.com/hugot/go-deltachat/deltachat
```

## Current version of libdeltachat
The libdeltachat binary that is distributed with go-deltachat is the C FFI to:

`deltachat v1.0.0-beta.22`

## Usage example
This library wraps the entire FFI API that deltachat has to offer and can of course be
used however you please. The recommended way is to use deltachat.Client, which is a thin
layer on top of deltachat's context datatype. It will make handeling events that are
emitted by deltachat easier and removes the need for some boilerplate code.

The following example will send every line you type as a message to chat@example.com:
```go
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/hugot/go-deltachat/deltachat"
)

func main() {
	client := &deltachat.Client{}

	addr := "chat@example.com"

	client.On(deltachat.DC_EVENT_IMAP_CONNECTED, func(c *deltachat.Context, e *deltachat.Event) {
		contactID := c.CreateContact(nil, &addr)
		chatID := c.CreateChatByContactID(contactID)

		c.SendTextMessage(chatID, "Hello World!")

		log.Println("Sent hello world message!")
	})

	// Handler for info logs from libdeltachat
	client.On(deltachat.DC_EVENT_INFO, func(c *deltachat.Context, e *deltachat.Event) {
		info, _ := e.Data2.String()

		log.Println(*info)
	})

	client.Open("/var/lib/deltabot/stuff.db")


	// Bear in mind that the config parameters are stored in the sqlite database
	// So these config values do not need to be set during each run.
	if !client.IsConfigured() {
		log.Println("Configuring")

		// Connection parameters, change as necessary
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
	}

	wait := make(chan os.Signal, 1)
	signal.Notify(wait, os.Interrupt)

	c := client.Context()
	contactID := c.CreateContact(nil, &addr)
	chatID := c.CreateChatByContactID(contactID)

	reader := bufio.NewReader(os.Stdin)
	messageChan := make(chan string)

	go func() {
		for {
			fmt.Print("Enter text: ")
			text, _ := reader.ReadString('\n')
			messageChan <- text
		}
	}()

	for {
		select {
		case sig := <-wait:
			log.Println(sig)

			// Give dc an opportunity to perform some shutdown logic
			// and close it's db.
			client.Close()
			return
		case text := <-messageChan:
			c.SendTextMessage(chatID, text)
		}
	}
}
```

## Documentation
Since the API is pretty much the go-equivalent of the C API, it should suffice to read the
documentation at [c.delta.chat](https://c.delta.chat) for most datatypes and functions.

## Platform support
This library should work on all platforms that are supported by Go and Deltachat. The
included libdeltachat file has been compiled on an amd64 system, so no extra setup is
required on that architecture. On other CPU architectures you will need to compile the
deltachat core FFI and place it in the deltachat-ffi folder before compiling your go code.
