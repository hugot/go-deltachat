package deltachat

// #include <deltachat.h>
import "C"

// Lot wraps dc_lot_t
type Lot struct {
	lot *C.dc_lot_t
}

func (l *Lot) GetID() uint32 {
	return uint32(C.dc_lot_get_id(l.lot))
}

func (l *Lot) GetState() int {
	return int(C.dc_lot_get_state(l.lot))
}

func (l *Lot) GetText1() string {
	return cStringToGo(C.dc_lot_get_text1(l.lot))
}

func (l *Lot) GetText1Meaning() int {
	return int(C.dc_lot_get_text1_meaning(l.lot))
}

func (l *Lot) GetText2() string {
	return cStringToGo(C.dc_lot_get_text2(l.lot))
}

func (l *Lot) GetTimestamp() int64 {
	return int64(C.dc_lot_get_timestamp(l.lot))
}

func (l *Lot) Unref() {
	C.dc_lot_unref(l.lot)
}
