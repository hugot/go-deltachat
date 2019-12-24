package deltachat

// #cgo CFLAGS: -I../deltachat-ffi/include
// #cgo LDFLAGS: -L../deltachat-ffi/lib -ldeltachat -ldl -lm
// #include <deltachat.h>
// #include <godeltachat.h>
import "C"

import (
	"log"
	"sync"
)

var deltachatCbMutex sync.RWMutex

type deltachatCallback func(
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

	return callback(event, data1, data2)
}

func NewContext() *Context {
	context := C.godeltachat_create_context()

	client := &Context{
		context: context,
	}

	deltachatCbMutex.Lock()
	deltachatCallbacks[context] = client.handleEvent
	deltachatCbMutex.Unlock()

	return client
}

type Context struct {
	context      *C.dc_context_t
	eventHandler EventHandler
	imapQuit     chan struct{}
	smtpQuit     chan struct{}
}

type EventHandler func(
	event int,
	data1 C.uintptr_t,
	data2 C.uintptr_t,
) int

func (c *Context) SetConfig(key string, value string) {
	cKey, cValue := C.CString(key), C.CString(value)
	C.dc_set_config(c.context, cKey, cValue)
	freeCString(cKey, cValue)
}

func (c *Context) GetConfig(key string) string {
	return dcStringToGo(C.dc_get_config(c.context, C.CString(key)))
}

func (c *Context) Open(databaseLocation string) bool {
	cDatabaseLocation := C.CString(databaseLocation)
	defer freeCString(cDatabaseLocation)

	return int(C.dc_open(c.context, cDatabaseLocation, nil)) > 0
}

func (c *Context) Configure() {
	C.dc_configure(c.context)
}

func (c *Context) SetHandler(handler EventHandler) {
	c.eventHandler = handler
}

func (c *Context) handleEvent(
	event C.int,
	data1 C.uintptr_t,
	data2 C.uintptr_t,
) C.uintptr_t {
	if c.eventHandler == nil {
		return 0
	}

	return C.uintptr_t(c.eventHandler(int(event), data1, data2))
}

func (c *Context) imapRoutine(quit chan struct{}) {
	for {
		select {
		case <-quit:
			log.Println("Quitting IMAP routine")
			return
		default:
			C.dc_perform_imap_jobs(c.context)
			C.dc_perform_imap_fetch(c.context)
			C.dc_perform_imap_idle(c.context)
		}
	}
}

func (c *Context) smtpRoutine(quit chan struct{}) {
	for {
		select {
		case <-quit:
			log.Println("Quitting SMTP routine")
			return
		default:
			C.dc_perform_smtp_jobs(c.context)
			C.dc_perform_smtp_idle(c.context)
		}
	}
}

func (c *Context) PerformIMAPJobs() {
	C.dc_perform_imap_jobs(c.context)
}

func (c *Context) PerformIMAPFetch() {
	C.dc_perform_imap_fetch(c.context)
}

func (c *Context) PerformIMAPIdle() {
	C.dc_perform_imap_idle(c.context)
}

func (c *Context) PerformSMTPJobs() {
	C.dc_perform_smtp_jobs(c.context)
}

func (c *Context) PerformSMTPIdle() {
	C.dc_perform_smtp_idle(c.context)
}

func (c *Context) StartWorkers() {
	c.imapQuit = make(chan struct{})
	c.smtpQuit = make(chan struct{})

	go c.imapRoutine(c.imapQuit)
	go c.smtpRoutine(c.smtpQuit)
}

func (c *Context) StopWorkers() {
	close(c.imapQuit)
	close(c.smtpQuit)
}

func (c *Context) CreateChatByContactID(ID uint32) uint32 {
	return uint32(C.dc_create_chat_by_contact_id(c.context, C.uint32_t(ID)))
}

func (c *Context) SendTextMessage(chatID uint32, message string) {
	cMessage := C.CString(message)
	C.dc_send_text_msg(c.context, C.uint32_t(chatID), cMessage)
	freeCString(cMessage)
}

func (c *Context) CreateContact(name *string, address *string) uint32 {
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

func (c *Context) NewMessage(viewType int) *Message {
	cMsg := C.dc_msg_new(c.context, C.int(viewType))

	return &Message{
		msg: cMsg,
	}
}

func (c *Context) AddAddressBook(addrBook string) int {
	cAddrBook := C.CString(addrBook)
	addedContacts := int(C.dc_add_address_book(c.context, cAddrBook))
	freeCString(cAddrBook)

	return addedContacts
}

func (c *Context) AddContactToChat(chatID uint32, contactID uint32) int {
	return int(C.dc_add_contact_to_chat(c.context, C.uint32_t(chatID), C.uint32_t(contactID)))
}

func (c *Context) AddDeviceMessage(label *string, message *Message) uint32 {
	var cLabel *C.char

	if label != nil {
		cLabel = C.CString(*label)
	}

	msgID := uint32(C.dc_add_device_msg(c.context, cLabel, message.getCMessage()))

	freeCString(cLabel)

	return msgID
}

func (c *Context) ArchiveChat(chatID uint32, archive int) {
	C.dc_archive_chat(c.context, C.uint32_t(chatID), C.int(archive))
}

func (c *Context) BlockContact(contactID uint32, block int) {
	C.dc_block_contact(c.context, C.uint32_t(contactID), C.int(block))
}

func (c *Context) CheckQR(QR string) *Lot {
	cQR := C.CString(QR)
	lot := C.dc_check_qr(c.context, cQR)
	freeCString(cQR)

	return &Lot{
		lot: lot,
	}
}

func (c *Context) Close() {
	c.StopWorkers()

	// The idle jobs can sometimes hang for a bit too long, interrupt them so that the
	// routines can receive the stop signal through their channels.
	c.InterruptIMAPIdle()
	c.InterruptSMTPIdle()

	C.dc_close(c.context)
}

func (c *Context) Unref() {
	C.dc_context_unref(c.context)
}

func (c *Context) ContinueKeyTransfer(msgID uint32, setupCode string) bool {
	cSetupCode := C.CString(setupCode)

	success := int(C.dc_continue_key_transfer(c.context, C.uint32_t(msgID), cSetupCode)) > 0

	freeCString(cSetupCode)

	return success
}

func (c *Context) CreateChatByMessageID(msgID uint32) uint32 {
	return uint32(C.dc_create_chat_by_msg_id(c.context, C.uint32_t(msgID)))
}

func (c *Context) CreateGroupChat(verified bool, name string) uint32 {
	cVerified := 0
	if verified {
		cVerified = 1
	}

	cName := C.CString(name)

	chatID := uint32(C.dc_create_group_chat(c.context, C.int(cVerified), cName))

	freeCString(cName)

	return chatID
}

func (c *Context) DeleteAllLocations() {
	C.dc_delete_all_locations(c.context)
}

func (c *Context) DeleteChat(chatID uint32) {
	C.dc_delete_chat(c.context, C.uint32_t(chatID))
}

func (c *Context) DeleteContact(contactID uint32) bool {
	return int(C.dc_delete_contact(c.context, C.uint32_t(contactID))) > 0
}

func (c *Context) DeleteMessages(messageIDs []uint32, messageCount int) {
	if messageCount > 0 {
		C.dc_delete_msgs(c.context, (*C.uint32_t)(&messageIDs[0]), C.int(messageCount))
	}
}

func (c *Context) EmptyServer(flags uint32) {
	C.dc_empty_server(c.context, C.uint32_t(flags))
}

func (c *Context) ForwardMessages(messageIDs []uint32, messageCount int, chatID uint32) {
	if messageCount > 0 {
		C.dc_forward_msgs(
			c.context,
			(*C.uint32_t)(&messageIDs[0]),
			C.int(messageCount),
			C.uint32_t(chatID),
		)
	}
}

func (c *Context) GetBlobDir() string {
	return dcStringToGo(C.dc_get_blobdir(c.context))
}

func (c *Context) GetBlockedCount() int {
	return int(C.dc_get_blocked_cnt(c.context))
}

func (c *Context) GetBlockedContacts() *Array {
	cArray := C.dc_get_blocked_contacts(c.context)

	return &Array{
		array: cArray,
	}
}

func (c *Context) GetChat(chatID uint32) *Chat {
	cChat := C.dc_get_chat(c.context, C.uint32_t(chatID))

	return &Chat{
		chat: cChat,
	}
}

func (c *Context) GetChatContacts(chatID uint32) *Array {
	cArray := C.dc_get_chat_contacts(c.context, C.uint32_t(chatID))

	return &Array{
		array: cArray,
	}
}

func (c *Context) GetChatIDByContactID(contactID uint32) uint32 {
	return uint32(C.dc_get_chat_id_by_contact_id(c.context, C.uint32_t(contactID)))
}

func (c *Context) GetChatMedia(chatID uint32, msgType1 int, msgType2 int, msgType3 int) *Array {
	cArray := C.dc_get_chat_media(
		c.context,
		C.uint32_t(chatID),
		C.int(msgType1),
		C.int(msgType2),
		C.int(msgType3),
	)

	return &Array{
		array: cArray,
	}
}

func (c *Context) GetChatMessages(chatID uint32, flags uint32, marker1Before uint32) *Array {
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

func (c *Context) GetChatlist(flags int, queryString string, queryID uint32) *ChatList {
	cQueryString := C.CString(queryString)

	cList := C.dc_get_chatlist(c.context, C.int(flags), cQueryString, C.uint32_t(queryID))

	freeCString(cQueryString)

	return &ChatList{
		list: cList,
	}
}

func (c *Context) GetContact(contactID uint32) *Contact {
	cContact := C.dc_get_contact(c.context, C.uint32_t(contactID))

	return &Contact{
		contact: cContact,
	}
}

func (c *Context) GetContactEncryptionInfo(contactID uint32) string {
	return dcStringToGo(C.dc_get_contact_encrinfo(c.context, C.uint32_t(contactID)))
}

func (c *Context) GetContacts(flags uint32, query *string) *Array {
	var cQuery *C.char
	if query != nil {
		cQuery = C.CString(*query)
	}

	cArray := C.dc_get_contacts(c.context, C.uint32_t(flags), cQuery)

	freeCString(cQuery)

	return &Array{
		array: cArray,
	}
}

func (c *Context) GetDraft(chatID uint32) *Message {
	cMsg := C.dc_get_draft(c.context, C.uint32_t(chatID))

	return &Message{
		msg: cMsg,
	}
}

func (c *Context) GetFreshMessageCount(chatID uint32) int {
	return int(C.dc_get_fresh_msg_cnt(c.context, C.uint32_t(chatID)))
}

func (c *Context) GetFreshMessages() *Array {
	return &Array{
		array: C.dc_get_fresh_msgs(c.context),
	}
}

func (c *Context) GetInfo() string {
	return dcStringToGo(C.dc_get_info(c.context))
}

func (c *Context) GetLocations(
	chatID uint32,
	contactID uint32,
	timestampBegin int64,
	timestampEnd int64,
) *Array {
	cArray := C.dc_get_locations(
		c.context,
		C.uint32_t(chatID),
		C.uint32_t(contactID),
		C.int64_t(timestampBegin),
		C.int64_t(timestampEnd),
	)

	return &Array{
		array: cArray,
	}
}

func (c *Context) GetMIMEHeaders(messageID uint32) string {
	return dcStringToGo(C.dc_get_mime_headers(c.context, C.uint32_t(messageID)))
}

func (c *Context) GetMessage(messageID uint32) *Message {
	msg := C.dc_get_msg(c.context, C.uint32_t(messageID))

	return &Message{
		msg: msg,
	}
}

func (c *Context) GetMessageCount(chatID uint32) int {
	return int(C.dc_get_msg_cnt(c.context, C.uint32_t(chatID)))
}

func (c *Context) GetMessageInfo(msgID uint32) string {
	return dcStringToGo(C.dc_get_msg_info(c.context, C.uint32_t(msgID)))
}

func (c *Context) GetNextMedia(
	msgID uint32,
	dir int,
	msgType int,
	msgType2 int,
	msgType3 int,
) uint32 {
	return uint32(
		C.dc_get_next_media(
			c.context,
			C.uint32_t(msgID),
			C.int(dir),
			C.int(msgType),
			C.int(msgType2),
			C.int(msgType3),
		),
	)
}

func (c *Context) GetOauth2URL(address string, redirectURI string) string {
	cAddr, cURI := C.CString(address), C.CString(redirectURI)
	defer freeCString(cAddr, cURI)

	return dcStringToGo(C.dc_get_oauth2_url(c.context, cAddr, cURI))
}

func (c *Context) GetSecurejoinQR(chatID uint32) string {
	return dcStringToGo(C.dc_get_securejoin_qr(c.context, C.uint32_t(chatID)))
}

func (c *Context) ImportExport(what int, param1 *string, param2 *string) {
	var cParam1 *C.char
	var cParam2 *C.char

	if param1 != nil {
		cParam1 = C.CString(*param1)
	}

	if param2 != nil {
		cParam2 = C.CString(*param2)
	}

	defer freeCString(cParam1, cParam2)

	C.dc_imex(c.context, C.int(what), cParam1, cParam2)
}

func (c *Context) ImportExportHasBackup(dir string) *string {
	cDir := C.CString(dir)
	defer freeCString(cDir)

	cExportDir := C.dc_imex_has_backup(c.context, cDir)

	if cExportDir != nil {
		exportDir := dcStringToGo(cExportDir)
		return &exportDir
	}

	return nil
}

func (c *Context) InitiateKeyTransfer() {
	C.dc_initiate_key_transfer(c.context)
}

func (c *Context) InterruptIMAPIdle() {
	C.dc_interrupt_imap_idle(c.context)
}

func (c *Context) InterruptMvboxIdle() {
	C.dc_interrupt_mvbox_idle(c.context)
}

func (c *Context) InterruptSentboxIdle() {
	C.dc_interrupt_sentbox_idle(c.context)
}

func (c *Context) InterruptSMTPIdle() {
	C.dc_interrupt_smtp_idle(c.context)
}

func (c *Context) IsConfigured() bool {
	return int(C.dc_is_configured(c.context)) > 0
}

func (c *Context) IsContactInChat(chatID uint32, contactID uint32) bool {
	return int(C.dc_is_contact_in_chat(c.context, C.uint32_t(chatID), C.uint32_t(contactID))) > 0
}

func (c *Context) IsOpen() bool {
	return int(C.dc_is_open(c.context)) > 0
}

func (c *Context) IsSendingLocationsToChat(chatID uint32) bool {
	return int(C.dc_is_sending_locations_to_chat(c.context, C.uint32_t(chatID))) > 0
}

func (c *Context) JoinSecurejoin(QR string) uint32 {
	cQR := C.CString(QR)
	defer freeCString(cQR)

	return uint32(C.dc_join_securejoin(c.context, cQR))
}

func (c *Context) LookupContactIDByAddress(address string) uint32 {
	cAddr := C.CString(address)
	defer freeCString(cAddr)

	return uint32(C.dc_lookup_contact_id_by_addr(c.context, cAddr))
}

func (c *Context) MarkNoticedAllChats() {
	C.dc_marknoticed_all_chats(c.context)
}

func (c *Context) MarkNoticedChat(chatID uint32) {
	C.dc_marknoticed_chat(c.context, C.uint32_t(chatID))
}

func (c *Context) MarkNoticedContact(contactID uint32) {
	C.dc_marknoticed_contact(c.context, C.uint32_t(contactID))
}

func (c *Context) MarkSeenMessages(messageIDs []uint32, messageCount int) {
	if messageCount > 0 {
		C.dc_markseen_msgs(
			c.context,
			(*C.uint32_t)(&messageIDs[0]),
			C.int(messageCount),
		)
	}
}

func MayBeValidAddress(address string) bool {
	cAddr := C.CString(address)
	defer freeCString(cAddr)

	return int(C.dc_may_be_valid_addr(cAddr)) > 0
}

func (c *Context) MaybeNetwork() {
	C.dc_maybe_network(c.context)
}

func (c *Context) PerformMvboxFetch() {
	C.dc_perform_mvbox_fetch(c.context)
}

func (c *Context) PerformMvboxIdle() {
	C.dc_perform_mvbox_idle(c.context)
}

func (c *Context) PerformMvboxJobs() {
	C.dc_perform_mvbox_jobs(c.context)
}

func (c *Context) PerformSentboxFetch() {
	C.dc_perform_sentbox_fetch(c.context)
}

func (c *Context) PerformSentboxIdle() {
	C.dc_perform_sentbox_idle(c.context)
}

func (c *Context) PerformSentboxJobs() {
	C.dc_perform_sentbox_jobs(c.context)
}

func (c *Context) PrepareMessage(chatID uint32, message *Message) uint32 {
	return uint32(C.dc_prepare_msg(c.context, C.uint32_t(chatID), message.getCMessage()))
}

func (c *Context) RemoveContactFromChat(chatID uint32, contactID uint32) bool {
	removed := C.dc_remove_contact_from_chat(c.context, C.uint32_t(chatID), C.uint32_t(contactID))

	return int(removed) > 0
}

func (c *Context) SearchMessages(chatID uint32, query string) *Array {
	cQuery := C.CString(query)
	defer freeCString(cQuery)

	cArray := C.dc_search_msgs(c.context, C.uint32_t(chatID), cQuery)

	if cArray != nil {
		return &Array{
			array: cArray,
		}
	}

	return nil
}

func (c *Context) SendLocationsToChat(chatID uint32, seconds int) {
	C.dc_send_locations_to_chat(c.context, C.uint32_t(chatID), C.int(seconds))
}

func (c *Context) SendMessage(chatID uint32, message *Message) uint32 {
	return uint32(C.dc_send_msg(c.context, C.uint32_t(chatID), message.getCMessage()))
}

func (c *Context) SetChatName(chatID uint32, name string) bool {
	cName := C.CString(name)
	defer freeCString(cName)

	return int(C.dc_set_chat_name(c.context, C.uint32_t(chatID), cName)) > 0
}

func (c *Context) SetChatProfileImage(chatID uint32, image string) bool {
	cImage := C.CString(image)
	defer freeCString(cImage)

	return int(C.dc_set_chat_profile_image(c.context, C.uint32_t(chatID), cImage)) > 0
}

func (c *Context) SetDraft(chatID uint32, message *Message) {
	C.dc_set_draft(c.context, C.uint32_t(chatID), message.getCMessage())
}

func (c *Context) SetLocation(latitude float64, longitude float64, accuracy float64) bool {
	success := C.dc_set_location(
		c.context,
		C.double(latitude),
		C.double(longitude),
		C.double(accuracy),
	)

	return int(success) > 0
}

func (c *Context) SetStockTranslation(stockID uint32, message string) bool {
	cMessage := C.CString(message)
	defer freeCString(cMessage)

	return int(C.dc_set_stock_translation(c.context, C.uint32_t(stockID), cMessage)) > 0
}

func (c *Context) StarMessages(messageIDs []uint32, messageCount int, star int) {
	if messageCount > 0 {
		C.dc_star_msgs(
			c.context,
			(*C.uint32_t)(&messageIDs[0]),
			C.int(messageCount),
			C.int(star),
		)
	}
}

func (c *Context) StopOngoingProcess() {
	C.dc_stop_ongoing_process(c.context)
}

func (c *Context) WasDeviceMessageEverAdded(label string) bool {
	cLabel := C.CString(label)
	defer freeCString(cLabel)

	return int(C.dc_was_device_msg_ever_added(c.context, cLabel)) > 0
}
