package deltachat

// #include <deltachat.h>
import "C"

type ChatList struct {
	list *C.dc_chatlist_t
}
