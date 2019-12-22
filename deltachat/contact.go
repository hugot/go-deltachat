package deltachat

// #include <deltachat.h>
import "C"

type Contact struct {
	contact *C.dc_contact_t
}
