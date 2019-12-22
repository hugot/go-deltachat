package deltachat

// #include <deltachat.h>
import "C"

type Contact struct {
	contact *C.dc_contact_t
}

func (c *Contact) GetAddress() string {
	return dcStringToGo(C.dc_contact_get_addr(c.contact))
}

func (c *Contact) GetColor() uint32 {
	return uint32(C.dc_contact_get_color(c.contact))
}

func (c *Contact) GetDisplayName() string {
	return dcStringToGo(C.dc_contact_get_display_name(c.contact))
}

func (c *Contact) GetFirstName() string {
	return dcStringToGo(C.dc_contact_get_first_name(c.contact))
}

func (c *Contact) GetID() uint32 {
	return uint32(C.dc_contact_get_id(c.contact))
}

func (c *Contact) GetName() string {
	return dcStringToGo(C.dc_contact_get_name(c.contact))
}

func (c *Contact) GetNameAndAddress() string {
	return dcStringToGo(C.dc_contact_get_name_n_addr(c.contact))
}

func (c *Contact) GetProfileImage() string {
	return dcStringToGo(C.dc_contact_get_profile_image(c.contact))
}

func (c *Contact) IsBlocked() bool {
	return int(C.dc_contact_is_blocked(c.contact)) > 0
}

func (c *Contact) Unref() {
	C.dc_contact_unref(c.contact)
}
