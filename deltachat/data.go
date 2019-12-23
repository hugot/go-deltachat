package deltachat

// #include <stdint.h>
import "C"
import (
	"errors"
	"fmt"
	"unsafe"
)

type Data struct {
	DataType uint8
	data     C.uintptr_t
}

func dataTypeError(expected uint8, actual uint8) error {
	return errors.New(
		fmt.Sprintf(
			"Attempted to unwrap Data to %s, but type is %s",
			dataTypeNames[expected],
			dataTypeNames[actual],
		),
	)
}

func (d *Data) String() (*string, error) {
	if d.DataType != DATA_TYPE_STRING {
		return nil, dataTypeError(DATA_TYPE_STRING, d.DataType)
	}

	// Now there's a lot of stuff going on, not sure if this can be done in a prettier
	// way..
	str := C.GoString((*C.char)(unsafe.Pointer(uintptr(d.data))))

	return &str, nil
}

func (d *Data) Int() int {
	return int(d.data)
}
