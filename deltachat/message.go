package deltachat

// #include <deltachat.h>
import "C"

// Message is a struct that wraps dc_msg_t
type Message struct {
	msg *C.dc_msg_t
}

func (m *Message) GetChatID() uint32 {
	return uint32(C.dc_msg_get_chat_id(m.msg))
}

func (m *Message) GetDuration() int {
	return int(C.dc_msg_get_duration(m.msg))
}

func (m *Message) GetFile() string {
	return cStringToGo(C.dc_msg_get_file(m.msg))
}

func (m *Message) GetFileBytes() uint64 {
	return uint64(C.dc_msg_get_filebytes(m.msg))
}

func (m *Message) GetFileMIME() string {
	return cStringToGo(C.dc_msg_get_filemime(m.msg))
}

func (m *Message) GetFromID() uint32 {
	return uint32(C.dc_msg_get_from_id(m.msg))
}

func (m *Message) GetHeight() int {
	return int(C.dc_msg_get_height(m.msg))
}

func (m *Message) GetID() uint32 {
	return uint32(C.dc_msg_get_id(m.msg))
}

func (m *Message) GetReceivedTimestamp() int64 {
	return int64(C.dc_msg_get_received_timestamp(m.msg))
}

func (m *Message) GetSetupCodeBegin() string {
	return cStringToGo(C.dc_msg_get_setupcodebegin(m.msg))
}

func (m *Message) GetShowPadlock() bool {
	return int(C.dc_msg_get_showpadlock(m.msg)) > 0
}

func (m *Message) GetSortTimestamp() int64 {
	return int64(C.dc_msg_get_sort_timestamp(m.msg))
}

func (m *Message) GetState() int {
	return int(C.dc_msg_get_state(m.msg))
}

func (m *Message) GetSummary(chat *Chat) *Lot {
	cLot := C.dc_msg_get_summary(m.msg, chat.getCChat())

	return &Lot{
		lot: cLot,
	}
}

func (m *Message) GetSummaryText(characters int) string {
	return cStringToGo(C.dc_msg_get_summarytext(m.msg, C.int(characters)))
}

func (m *Message) GetText() string {
	return cStringToGo(C.dc_msg_get_text(m.msg))
}

func (m *Message) GetTimestamp() int64 {
	return int64(C.dc_msg_get_timestamp(m.msg))
}

func (m *Message) GetViewType() int {
	return int(C.dc_msg_get_viewtype(m.msg))
}

func (m *Message) GetWidth() int {
	return int(C.dc_msg_get_width(m.msg))
}

func (m *Message) HasDeviatingTimestamp() bool {
	return int(C.dc_msg_has_deviating_timestamp(m.msg)) > 0
}

func (m *Message) HasLocation() bool {
	return int(C.dc_msg_has_location(m.msg)) > 0
}

func (m *Message) IsForwarded() bool {
	return int(C.dc_msg_is_forwarded(m.msg)) > 0
}

func (m *Message) IsInCreation() bool {
	return int(C.dc_msg_is_increation(m.msg)) > 0
}

func (m *Message) IsInfo() bool {
	return int(C.dc_msg_is_info(m.msg)) > 0
}

func (m *Message) IsSent() bool {
	return int(C.dc_msg_is_sent(m.msg)) > 0
}

func (m *Message) IsSetupMessage() bool {
	return int(C.dc_msg_is_setupmessage(m.msg)) > 0
}

func (m *Message) IsStarred() bool {
	return int(C.dc_msg_is_starred(m.msg)) > 0
}

func (m *Message) LateFilingMediaSize(width int, height int, duration int) {
	C.dc_msg_latefiling_mediasize(m.msg, C.int(width), C.int(height), C.int(duration))
}

func (m *Message) SetDimension(width int, height int) {
	C.dc_msg_set_dimension(m.msg, C.int(width), C.int(height))
}

func (m *Message) SetDuration(duration int) {
	C.dc_msg_set_duration(m.msg, C.int(duration))
}

func (m *Message) SetFile(file string, fileMIME string) {
	cFile, cFileMIME := C.CString(file), C.CString(fileMIME)

	C.dc_msg_set_file(m.msg, cFile, cFileMIME)

	freeCString(cFile, cFileMIME)
}

func (m *Message) SetLocation(latitude float64, longitude float64) {
	C.dc_msg_set_location(m.msg, C.double(latitude), C.double(longitude))
}

func (m *Message) SetText(text string) {
	cText := C.CString(text)

	C.dc_msg_set_text(m.msg, cText)

	freeCString(cText)
}

func (m *Message) Unref() {
	C.dc_msg_unref(m.msg)
}
