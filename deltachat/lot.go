package deltachat

// #include <deltachat.h>
import "C"

// Lot wraps dc_lot_t
type Lot struct {
	lot *C.dc_lot_t
}
