package deltachat

// #include <deltachat.h>
import "C"

type Chat struct {
	chat *C.dc_chat_t
}

func (c *Chat) getCChat() *C.dc_chat_t {
	return c.chat
}
