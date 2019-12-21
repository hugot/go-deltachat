package deltachat

// #include <deltachat.h>
import "C"

// Array wraps dc_array_t
type Array struct {
	array *C.dc_array_t
}
