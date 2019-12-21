package deltachat

// #include <stdlib.h>
import "C"
import "unsafe"

func freeCString(strings ...*C.char) {
	for _, s := range strings {
		C.free(unsafe.Pointer(s))
	}
}

func cStringToGo(cStr *C.char) string {
	str := C.GoString(cStr)
	freeCString(cStr)

	return str
}
