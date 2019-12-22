package deltachat

// #include <deltachat.h>
import "C"

// Array wraps dc_array_t
type Array struct {
	array *C.dc_array_t
}

func (a *Array) GetAccuracy(index uint) float64 {
	return float64(C.dc_array_get_accuracy(a.array, C.size_t(index)))
}

func (a *Array) GetChatID(index uint) uint32 {
	return uint32(C.dc_array_get_chat_id(a.array, C.size_t(index)))
}

func (a *Array) GetCount() uint {
	return uint(C.dc_array_get_cnt(a.array))
}

func (a *Array) GetID(index uint) uint32 {
	return uint32(C.dc_array_get_id(a.array, C.size_t(index)))
}

func (a *Array) GetLatitude(index uint) float64 {
	return float64(C.dc_array_get_latitude(a.array, C.size_t(index)))
}

func (a *Array) GetLongitude(index uint) float64 {
	return float64(C.dc_array_get_longitude(a.array, C.size_t(index)))
}

// This string proably shouldn't be free'd as that might destroy part of whatever is in
// the array
func (a *Array) GetMarker(index uint) string {
	return C.GoString(C.dc_array_get_marker(a.array, C.size_t(index)))
}

func (a *Array) GetMessageID(index uint) uint32 {
	return uint32(C.dc_array_get_msg_id(a.array, C.size_t(index)))
}

// Returns a pointer to the raw C array data. intentionally not exposed outside of this
// library.
func (a *Array) getRaw() *C.uint32_t {
	return C.dc_array_get_raw(a.array)
}

func (a *Array) GetTimestamp(index uint) int64 {
	return int64(C.dc_array_get_timestamp(a.array, C.size_t(index)))
}

func (a *Array) IsIndependent(index uint) bool {
	return int(C.dc_array_is_independent(a.array, C.size_t(index))) > 0
}

func (a *Array) Unref() {
	C.dc_array_unref(a.array)
}
