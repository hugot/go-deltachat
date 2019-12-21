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
	cFile := C.dc_msg_get_file(m.msg)
	file := C.GoString(cFile)

	freeCString(cFile)

	return file
}

func (m *Message) GetFileBytes() uint64 {
	return uint64(C.dc_msg_get_filebytes(m.msg))
}

func (m *Message) GetFileMIME() string {
	cMIME := C.dc_msg_get_filemime(m.msg)
	MIME := C.GoString(cMIME)

	freeCString(cMIME)

	return MIME
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
	cCodeBegin := C.dc_msg_get_setupcodebegin(m.msg)
	codeBegin := C.GoString(cCodeBegin)

	freeCString(cCodeBegin)

	return codeBegin
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

// func (m *Message) GetSummary(chat *Chat)  {
// 	return
// }
