package deltachat

// #cgo CFLAGS: -I../deltachat-ffi/include
// #cgo LDFLAGS: -L../deltachat-ffi/lib -ldeltachat -ldl -lm
// #include <deltachat.h>
// #include <godeltachat.h>
import "C"

import (
	"sync"
)

var deltachatCbMutex sync.RWMutex

type deltachatCallback func(
	context *C.dc_context_t,
	event C.int,
	data1 C.uintptr_t,
	data2 C.uintptr_t,
) C.uintptr_t

var deltachatCallbacks map[*C.dc_context_t]deltachatCallback = make(
	map[*C.dc_context_t]deltachatCallback,
)

//export godeltachat_eventhandler_proxy
func godeltachat_eventhandler_proxy(
	context *C.dc_context_t,
	event C.int,
	data1 C.uintptr_t,
	data2 C.uintptr_t,
) C.uintptr_t {
	deltachatCbMutex.RLock()

	callback, ok := deltachatCallbacks[context]

	deltachatCbMutex.RUnlock()

	if !ok {
		panic("dc_context_t callback was called but not set")
	}

	return callback(context, event, data1, data2)
}

func NewClient() *Client {
	context := C.godeltachat_create_context()

	client := &Client{
		context: context,
	}

	deltachatCbMutex.Lock()
	deltachatCallbacks[context] = client.handleEvent
	deltachatCbMutex.Unlock()

	return client
}

type Client struct {
	context *C.dc_context_t
}

func (c *Client) SetConfig(key string, value string) {
	cKey, cValue := C.CString(key), C.CString(value)
	C.dc_set_config(c.context, cKey, cValue)
	freeCString(cKey, cValue)
}

func (c *Client) GetConfig(key string) string {
	return cStringToGo(C.dc_get_config(c.context, C.CString(key)))
}

func (c *Client) Open(databaseLocation string) {
	cDatabaseLocation := C.CString(databaseLocation)
	C.dc_open(c.context, cDatabaseLocation, nil)
	freeCString(cDatabaseLocation)
}

func (c *Client) Configure() {
	C.dc_configure(c.context)
}

func (c *Client) handleEvent(
	context *C.dc_context_t,
	event C.int,
	data1 C.uintptr_t,
	data2 C.uintptr_t,
) C.uintptr_t {
	return 0
}

func (c *Client) imapRoutine() {
	for {
		C.dc_perform_imap_jobs(c.context)
		C.dc_perform_imap_fetch(c.context)
		C.dc_perform_imap_idle(c.context)
	}
}

func (c *Client) smtpRoutine() {
	for {
		C.dc_perform_smtp_jobs(c.context)
		C.dc_perform_smtp_idle(c.context)
	}
}

func (c *Client) PerformIMAPJobs() {
	C.dc_perform_imap_jobs(c.context)
}

func (c *Client) PerformIMAPFetch() {
	C.dc_perform_imap_fetch(c.context)
}

func (c *Client) PerformIMAPIdle() {
	C.dc_perform_imap_idle(c.context)
}

func (c *Client) PerformSMTPJobs() {
	C.dc_perform_smtp_jobs(c.context)
}

func (c *Client) PerformSMTPIdle() {
	C.dc_perform_smtp_idle(c.context)
}

func (c *Client) StartWorkers() {
	go c.imapRoutine()
	go c.smtpRoutine()
}

func (c *Client) CreateChatByContactID(ID uint32) uint32 {
	return uint32(C.dc_create_chat_by_contact_id(c.context, C.uint32_t(ID)))
}

func (c *Client) SendTextMessage(chatID uint32, message string) {
	cMessage := C.CString(message)
	C.dc_send_text_msg(c.context, C.uint32_t(chatID), cMessage)
	freeCString(cMessage)
}

func (c *Client) CreateContact(name *string, address *string) uint32 {
	var nameString *C.char

	if name != nil {
		nameString = C.CString(*name)
	}

	var addressString *C.char

	if address != nil {
		addressString = C.CString(*address)
	}

	contactID := C.dc_create_contact(c.context, nameString, addressString)

	freeCString(nameString, addressString)

	return uint32(contactID)
}

func (c *Client) NewMessage(viewType int) *Message {
	cMsg := C.dc_msg_new(c.context, C.int(viewType))

	return &Message{
		msg: cMsg,
	}
}

func (c *Client) AddAddressBook(addrBook string) int {
	cAddrBook := C.CString(addrBook)
	addedContacts := int(C.dc_add_address_book(c.context, cAddrBook))
	freeCString(cAddrBook)

	return addedContacts
}

func (c *Client) AddContactToChat(chatID uint32, contactID uint32) int {
	return int(C.dc_add_contact_to_chat(c.context, C.uint32_t(chatID), C.uint32_t(contactID)))
}

func (c *Client) AddDeviceMessage(label *string, message *Message) uint32 {
	var cLabel *C.char

	if label != nil {
		cLabel = C.CString(*label)
	}

	msgID := uint32(C.dc_add_device_msg(c.context, cLabel, message.getCMessage()))

	freeCString(cLabel)

	return msgID
}

func (c *Client) ArchiveChat(chatID uint32, archive int) {
	C.dc_archive_chat(c.context, C.uint32_t(chatID), C.int(archive))
}

func (c *Client) BlockContact(contactID uint32, block int) {
	C.dc_block_contact(c.context, C.uint32_t(contactID), C.int(block))
}

func (c *Client) CheckQR(QR string) *Lot {
	cQR := C.CString(QR)
	lot := C.dc_check_qr(c.context, cQR)
	freeCString(cQR)

	return &Lot{
		lot: lot,
	}
}

func (c *Client) Close() {
	C.dc_close(c.context)
}

func (c *Client) Unref() {
	C.dc_context_unref(c.context)
}

func (c *Client) ContinueKeyTransfer(msgID uint32, setupCode string) bool {
	cSetupCode := C.CString(setupCode)

	success := int(C.dc_continue_key_transfer(c.context, C.uint32_t(msgID), cSetupCode)) > 0

	freeCString(cSetupCode)

	return success
}

func (c *Client) CreateChatByMessageID(msgID uint32) uint32 {
	return uint32(C.dc_create_chat_by_msg_id(c.context, C.uint32_t(msgID)))
}

func (c *Client) CreateGroupChat(verified bool, name string) uint32 {
	cVerified := 0
	if verified {
		cVerified = 1
	}

	cName := C.CString(name)

	chatID := uint32(C.dc_create_group_chat(c.context, C.int(cVerified), cName))

	freeCString(cName)

	return chatID
}

func (c *Client) DeleteAllLocations() {
	C.dc_delete_all_locations(c.context)
}

func (c *Client) DeleteChat(chatID uint32) {
	C.dc_delete_chat(c.context, C.uint32_t(chatID))
}

func (c *Client) DeleteContact(contactID uint32) bool {
	return int(C.dc_delete_contact(c.context, C.uint32_t(contactID))) > 0
}

func (c *Client) DeleteMessages(messageIDs []uint32, messageCount int) {
	if messageCount > 0 {
		C.dc_delete_msgs(c.context, (*C.uint32_t)(&messageIDs[0]), C.int(messageCount))
	}
}

func (c *Client) EmptyServer(flags uint32) {
	C.dc_empty_server(c.context, C.uint32_t(flags))
}

func (c *Client) ForwardMessages(messageIDs []uint32, messageCount int, chatID uint32) {
	if messageCount > 0 {
		C.dc_forward_msgs(
			c.context,
			(*C.uint32_t)(&messageIDs[0]),
			C.int(messageCount),
			C.uint32_t(chatID),
		)
	}
}

func (c *Client) GetBlobDir() string {
	return cStringToGo(C.dc_get_blobdir(c.context))
}

func (c *Client) GetBlockedCount() int {
	return int(C.dc_get_blocked_cnt(c.context))
}

func (c *Client) GetBlockedContacts() *Array {
	cArray := C.dc_get_blocked_contacts(c.context)

	return &Array{
		array: cArray,
	}
}

func (c *Client) GetChat(chatID uint32) *Chat {
	cChat := C.dc_get_chat(c.context, C.uint32_t(chatID))

	return &Chat{
		chat: cChat,
	}
}

func (c *Client) GetChatContacts(chatID uint32) *Array {
	cArray := C.dc_get_chat_contacts(c.context, C.uint32_t(chatID))

	return &Array{
		array: cArray,
	}
}

func (c *Client) GetChatIDByContactID(contactID uint32) uint32 {
	return uint32(C.dc_get_chat_id_by_contact_id(c.context, C.uint32_t(contactID)))
}

func (c *Client) GetChatMedia(chatID uint32, msgType1 int, msgType2 int, msgType3 int) *Array {
	cArray := C.dc_get_chat_media(c.context, C.int(msgType1), C.int(msgType2), C.int(msgType3))

	return &Array{
		array: cArray,
	}
}

func GetChatMessages(chatID uint32, flags uint32, marker1Before uint32) *Array {
	cArray := C.dc_get_chat_msgs(
		c.context,
		C.uint32_t(chatID),
		C.uint32_t(flags),
		C.uint32_t(marker1Before),
	)

	return &Array{
		array: cArray,
	}
}
