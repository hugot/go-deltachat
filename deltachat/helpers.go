package deltachat

// #include <stdlib.h>
// #include <deltachat.h>
import "C"
import "unsafe"

func freeCString(strings ...*C.char) {
	for _, s := range strings {
		C.free(unsafe.Pointer(s))
	}
}

func freeDCString(strings ...*C.char) {
	for _, s := range strings {
		dcStringUnref(s)
	}
}

func cStringToGo(cStr *C.char) string {
	str := C.GoString(cStr)
	freeCString(cStr)

	return str
}

func dcStringToGo(cStr *C.char) string {
	str := C.GoString(cStr)
	freeDCString(cStr)

	return str
}

func dcStringUnref(str *C.char) {
	C.dc_str_unref(str)
}
