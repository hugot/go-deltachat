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
	str      string
	i        int
}

func uintPtrToString(ptr C.uintptr_t) string {
	return C.GoString((*C.char)(unsafe.Pointer(uintptr(ptr))))
}

func NewData1(event int, data1 C.uintptr_t) *Data {
	return newData(event, data1, Data1TypeForEvent)
}

func NewData2(event int, data2 C.uintptr_t) *Data {
	return newData(event, data2, Data2TypeForEvent)
}

func newData(event int, data C.uintptr_t, typeProvider func(int) uint8) *Data {
	dataType := typeProvider(event)

	if dataType == DATA_TYPE_STRING {
		return &Data{
			DataType: dataType,
			str:      uintPtrToString(data),
		}
	}

	if dataType == DATA_TYPE_INT {
		return &Data{
			DataType: dataType,
			i:        int(data),
		}
	}

	return &Data{
		DataType: dataType,
	}
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

	return &d.str, nil
}

func (d *Data) Int() (*int, error) {
	if d.DataType != DATA_TYPE_INT {
		return nil, dataTypeError(DATA_TYPE_INT, d.DataType)
	}

	return &d.i, nil
}
