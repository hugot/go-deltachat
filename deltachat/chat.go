package deltachat

// #include <deltachat.h>
import "C"

type Chat struct {
	chat *C.dc_chat_t
}

func (c *Chat) getCChat() *C.dc_chat_t {
	return c.chat
}

func (c *Chat) CanSent() bool {
	return int(C.dc_chat_can_send(c.chat)) > 0
}

func (c *Chat) GetArchived() int {
	return int(C.dc_chat_get_archived(c.chat))
}

func (c *Chat) GetColor() uint32 {
	return uint32(C.dc_chat_get_color(c.chat))
}

func (c *Chat) GetID() uint32 {
	return uint32(C.dc_chat_get_id(c.chat))
}

func (c *Chat) GetName() string {
	return dcStringToGo(C.dc_chat_get_name(c.chat))
}

func (c *Chat) GetProfileImage() string {
	return dcStringToGo(C.dc_chat_get_profile_image(c.chat))
}

func (c *Chat) GetType() int {
	return int(C.dc_chat_get_type(c.chat))
}

func (c *Chat) IsDeviceTalk() bool {
	return int(C.dc_chat_is_device_talk(c.chat)) > 0
}

func (c *Chat) IsSelfTalk() bool {
	return int(C.dc_chat_is_self_talk(c.chat)) > 0
}

func (c *Chat) IsSendingLocations() bool {
	return int(C.dc_chat_is_sending_locations(c.chat)) > 0
}

func (c *Chat) IsUnpromoted() bool {
	return int(C.dc_chat_is_unpromoted(c.chat)) > 0
}

func (c *Chat) IsVerified() bool {
	return int(C.dc_chat_is_verified(c.chat)) > 0
}

func (c *Chat) Unref() {
	C.dc_chat_unref(c.chat)
}
