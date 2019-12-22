package deltachat

// #include <deltachat.h>
import "C"

type ChatList struct {
	list *C.dc_chatlist_t
}

func (c *ChatList) GetChatID(index uint) uint32 {
	return uint32(C.dc_chatlist_get_chat_id(c.list, C.size_t(index)))
}

func (c *ChatList) GetCount() int64 {
	return int64(C.dc_chatlist_get_cnt(c.list))
}

// dc_chatlist_get_context should return the original Context instance that was
// instantiated to wrap the dc_context_t. Not Implemented for now.
// func (c *ChatList) GetContext() *Context {
// }

func (c *ChatList) GetMessageID(index uint) uint32 {
	return uint32(C.dc_chatlist_get_msg_id(c.list, C.size_t(index)))
}

func (c *ChatList) GetSummary(index uint, chat *Chat) *Lot {
	cLot := C.dc_chatlist_get_summary(c.list, C.size_t(index), chat.getCChat())

	return &Lot{
		lot: cLot,
	}
}

func (c *ChatList) Unref() {
	C.dc_chatlist_unref(c.list)
}
