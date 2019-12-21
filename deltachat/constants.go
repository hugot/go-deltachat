package deltachat

// #include <deltachat.h>
import "C"

const (
	DC_GCL_ARCHIVED_ONLY                 = int(C.DC_GCL_ARCHIVED_ONLY)
	DC_GCL_NO_SPECIALS                   = int(C.DC_GCL_NO_SPECIALS)
	DC_GCL_ADD_ALLDONE_HINT              = int(C.DC_GCL_ADD_ALLDONE_HINT)
	DC_GCL_VERIFIED_ONLY                 = int(C.DC_GCL_VERIFIED_ONLY)
	DC_GCL_ADD_SELF                      = int(C.DC_GCL_ADD_SELF)
	DC_QR_ASK_VERIFYCONTACT              = int(C.DC_QR_ASK_VERIFYCONTACT)
	DC_QR_ASK_VERIFYGROUP                = int(C.DC_QR_ASK_VERIFYGROUP)
	DC_QR_FPR_OK                         = int(C.DC_QR_FPR_OK)
	DC_QR_FPR_MISMATCH                   = int(C.DC_QR_FPR_MISMATCH)
	DC_QR_FPR_WITHOUT_ADDR               = int(C.DC_QR_FPR_WITHOUT_ADDR)
	DC_QR_ADDR                           = int(C.DC_QR_ADDR)
	DC_QR_TEXT                           = int(C.DC_QR_TEXT)
	DC_QR_URL                            = int(C.DC_QR_URL)
	DC_QR_ERROR                          = int(C.DC_QR_ERROR)
	DC_CHAT_ID_DEADDROP                  = int(C.DC_CHAT_ID_DEADDROP)
	DC_CHAT_ID_TRASH                     = int(C.DC_CHAT_ID_TRASH)
	DC_CHAT_ID_MSGS_IN_CREATION          = int(C.DC_CHAT_ID_MSGS_IN_CREATION)
	DC_CHAT_ID_STARRED                   = int(C.DC_CHAT_ID_STARRED)
	DC_CHAT_ID_ARCHIVED_LINK             = int(C.DC_CHAT_ID_ARCHIVED_LINK)
	DC_CHAT_ID_ALLDONE_HINT              = int(C.DC_CHAT_ID_ALLDONE_HINT)
	DC_CHAT_ID_LAST_SPECIAL              = int(C.DC_CHAT_ID_LAST_SPECIAL)
	DC_CHAT_TYPE_UNDEFINED               = int(C.DC_CHAT_TYPE_UNDEFINED)
	DC_CHAT_TYPE_SINGLE                  = int(C.DC_CHAT_TYPE_SINGLE)
	DC_CHAT_TYPE_GROUP                   = int(C.DC_CHAT_TYPE_GROUP)
	DC_CHAT_TYPE_VERIFIED_GROUP          = int(C.DC_CHAT_TYPE_VERIFIED_GROUP)
	DC_MSG_ID_MARKER1                    = int(C.DC_MSG_ID_MARKER1)
	DC_MSG_ID_DAYMARKER                  = int(C.DC_MSG_ID_DAYMARKER)
	DC_MSG_ID_LAST_SPECIAL               = int(C.DC_MSG_ID_LAST_SPECIAL)
	DC_STATE_UNDEFINED                   = int(C.DC_STATE_UNDEFINED)
	DC_STATE_IN_FRESH                    = int(C.DC_STATE_IN_FRESH)
	DC_STATE_IN_NOTICED                  = int(C.DC_STATE_IN_NOTICED)
	DC_STATE_IN_SEEN                     = int(C.DC_STATE_IN_SEEN)
	DC_STATE_OUT_PREPARING               = int(C.DC_STATE_OUT_PREPARING)
	DC_STATE_OUT_DRAFT                   = int(C.DC_STATE_OUT_DRAFT)
	DC_STATE_OUT_PENDING                 = int(C.DC_STATE_OUT_PENDING)
	DC_STATE_OUT_FAILED                  = int(C.DC_STATE_OUT_FAILED)
	DC_STATE_OUT_DELIVERED               = int(C.DC_STATE_OUT_DELIVERED)
	DC_STATE_OUT_MDN_RCVD                = int(C.DC_STATE_OUT_MDN_RCVD)
	DC_CONTACT_ID_SELF                   = int(C.DC_CONTACT_ID_SELF)
	DC_CONTACT_ID_INFO                   = int(C.DC_CONTACT_ID_INFO)
	DC_CONTACT_ID_DEVICE                 = int(C.DC_CONTACT_ID_DEVICE)
	DC_CONTACT_ID_LAST_SPECIAL           = int(C.DC_CONTACT_ID_LAST_SPECIAL)
	DC_MSG_TEXT                          = int(C.DC_MSG_TEXT)
	DC_MSG_IMAGE                         = int(C.DC_MSG_IMAGE)
	DC_MSG_GIF                           = int(C.DC_MSG_GIF)
	DC_MSG_STICKER                       = int(C.DC_MSG_STICKER)
	DC_MSG_AUDIO                         = int(C.DC_MSG_AUDIO)
	DC_MSG_VOICE                         = int(C.DC_MSG_VOICE)
	DC_MSG_VIDEO                         = int(C.DC_MSG_VIDEO)
	DC_MSG_FILE                          = int(C.DC_MSG_FILE)
	DC_LP_AUTH_OAUTH2                    = int(C.DC_LP_AUTH_OAUTH2)
	DC_LP_AUTH_NORMAL                    = int(C.DC_LP_AUTH_NORMAL)
	DC_LP_IMAP_SOCKET_STARTTLS           = int(C.DC_LP_IMAP_SOCKET_STARTTLS)
	DC_LP_IMAP_SOCKET_SSL                = int(C.DC_LP_IMAP_SOCKET_SSL)
	DC_LP_IMAP_SOCKET_PLAIN              = int(C.DC_LP_IMAP_SOCKET_PLAIN)
	DC_LP_SMTP_SOCKET_STARTTLS           = int(C.DC_LP_SMTP_SOCKET_STARTTLS)
	DC_LP_SMTP_SOCKET_SSL                = int(C.DC_LP_SMTP_SOCKET_SSL)
	DC_LP_SMTP_SOCKET_PLAIN              = int(C.DC_LP_SMTP_SOCKET_PLAIN)
	DC_LP_AUTH_FLAGS                     = int(C.DC_LP_AUTH_FLAGS)
	DC_LP_IMAP_SOCKET_FLAGS              = int(C.DC_LP_IMAP_SOCKET_FLAGS)
	DC_LP_SMTP_SOCKET_FLAGS              = int(C.DC_LP_SMTP_SOCKET_FLAGS)
	DC_EMPTY_MVBOX                       = int(C.DC_EMPTY_MVBOX)
	DC_EMPTY_INBOX                       = int(C.DC_EMPTY_INBOX)
	DC_EVENT_INFO                        = int(C.DC_EVENT_INFO)
	DC_EVENT_SMTP_CONNECTED              = int(C.DC_EVENT_SMTP_CONNECTED)
	DC_EVENT_IMAP_CONNECTED              = int(C.DC_EVENT_IMAP_CONNECTED)
	DC_EVENT_SMTP_MESSAGE_SENT           = int(C.DC_EVENT_SMTP_MESSAGE_SENT)
	DC_EVENT_IMAP_MESSAGE_DELETED        = int(C.DC_EVENT_IMAP_MESSAGE_DELETED)
	DC_EVENT_IMAP_MESSAGE_MOVED          = int(C.DC_EVENT_IMAP_MESSAGE_MOVED)
	DC_EVENT_IMAP_FOLDER_EMPTIED         = int(C.DC_EVENT_IMAP_FOLDER_EMPTIED)
	DC_EVENT_NEW_BLOB_FILE               = int(C.DC_EVENT_NEW_BLOB_FILE)
	DC_EVENT_DELETED_BLOB_FILE           = int(C.DC_EVENT_DELETED_BLOB_FILE)
	DC_EVENT_WARNING                     = int(C.DC_EVENT_WARNING)
	DC_EVENT_ERROR                       = int(C.DC_EVENT_ERROR)
	DC_EVENT_ERROR_NETWORK               = int(C.DC_EVENT_ERROR_NETWORK)
	DC_EVENT_ERROR_SELF_NOT_IN_GROUP     = int(C.DC_EVENT_ERROR_SELF_NOT_IN_GROUP)
	DC_EVENT_MSGS_CHANGED                = int(C.DC_EVENT_MSGS_CHANGED)
	DC_EVENT_INCOMING_MSG                = int(C.DC_EVENT_INCOMING_MSG)
	DC_EVENT_MSG_DELIVERED               = int(C.DC_EVENT_MSG_DELIVERED)
	DC_EVENT_MSG_FAILED                  = int(C.DC_EVENT_MSG_FAILED)
	DC_EVENT_MSG_READ                    = int(C.DC_EVENT_MSG_READ)
	DC_EVENT_CHAT_MODIFIED               = int(C.DC_EVENT_CHAT_MODIFIED)
	DC_EVENT_CONTACTS_CHANGED            = int(C.DC_EVENT_CONTACTS_CHANGED)
	DC_EVENT_LOCATION_CHANGED            = int(C.DC_EVENT_LOCATION_CHANGED)
	DC_EVENT_CONFIGURE_PROGRESS          = int(C.DC_EVENT_CONFIGURE_PROGRESS)
	DC_EVENT_IMEX_PROGRESS               = int(C.DC_EVENT_IMEX_PROGRESS)
	DC_EVENT_IMEX_FILE_WRITTEN           = int(C.DC_EVENT_IMEX_FILE_WRITTEN)
	DC_EVENT_SECUREJOIN_INVITER_PROGRESS = int(C.DC_EVENT_SECUREJOIN_INVITER_PROGRESS)
	DC_EVENT_SECUREJOIN_JOINER_PROGRESS  = int(C.DC_EVENT_SECUREJOIN_JOINER_PROGRESS)
	DC_EVENT_SECUREJOIN_MEMBER_ADDED     = int(C.DC_EVENT_SECUREJOIN_MEMBER_ADDED)
	DC_EVENT_FILE_COPIED                 = int(C.DC_EVENT_FILE_COPIED)
	DC_EVENT_IS_OFFLINE                  = int(C.DC_EVENT_IS_OFFLINE)
	DC_EVENT_GET_STRING                  = int(C.DC_EVENT_GET_STRING)
	DC_STR_SELFNOTINGRP                  = int(C.DC_STR_SELFNOTINGRP)
	DC_STR_NOMESSAGES                    = int(C.DC_STR_NOMESSAGES)
	DC_STR_SELF                          = int(C.DC_STR_SELF)
	DC_STR_DRAFT                         = int(C.DC_STR_DRAFT)
	DC_STR_MEMBER                        = int(C.DC_STR_MEMBER)
	DC_STR_CONTACT                       = int(C.DC_STR_CONTACT)
	DC_STR_VOICEMESSAGE                  = int(C.DC_STR_VOICEMESSAGE)
	DC_STR_DEADDROP                      = int(C.DC_STR_DEADDROP)
	DC_STR_IMAGE                         = int(C.DC_STR_IMAGE)
	DC_STR_VIDEO                         = int(C.DC_STR_VIDEO)
	DC_STR_AUDIO                         = int(C.DC_STR_AUDIO)
	DC_STR_FILE                          = int(C.DC_STR_FILE)
	DC_STR_STATUSLINE                    = int(C.DC_STR_STATUSLINE)
	DC_STR_NEWGROUPDRAFT                 = int(C.DC_STR_NEWGROUPDRAFT)
	DC_STR_MSGGRPNAME                    = int(C.DC_STR_MSGGRPNAME)
	DC_STR_MSGGRPIMGCHANGED              = int(C.DC_STR_MSGGRPIMGCHANGED)
	DC_STR_MSGADDMEMBER                  = int(C.DC_STR_MSGADDMEMBER)
	DC_STR_MSGDELMEMBER                  = int(C.DC_STR_MSGDELMEMBER)
	DC_STR_MSGGROUPLEFT                  = int(C.DC_STR_MSGGROUPLEFT)
	DC_STR_GIF                           = int(C.DC_STR_GIF)
	DC_STR_ENCRYPTEDMSG                  = int(C.DC_STR_ENCRYPTEDMSG)
	DC_STR_E2E_AVAILABLE                 = int(C.DC_STR_E2E_AVAILABLE)
	DC_STR_ENCR_TRANSP                   = int(C.DC_STR_ENCR_TRANSP)
	DC_STR_ENCR_NONE                     = int(C.DC_STR_ENCR_NONE)
	DC_STR_CANTDECRYPT_MSG_BODY          = int(C.DC_STR_CANTDECRYPT_MSG_BODY)
	DC_STR_FINGERPRINTS                  = int(C.DC_STR_FINGERPRINTS)
	DC_STR_READRCPT                      = int(C.DC_STR_READRCPT)
	DC_STR_READRCPT_MAILBODY             = int(C.DC_STR_READRCPT_MAILBODY)
	DC_STR_MSGGRPIMGDELETED              = int(C.DC_STR_MSGGRPIMGDELETED)
	DC_STR_E2E_PREFERRED                 = int(C.DC_STR_E2E_PREFERRED)
	DC_STR_CONTACT_VERIFIED              = int(C.DC_STR_CONTACT_VERIFIED)
	DC_STR_CONTACT_NOT_VERIFIED          = int(C.DC_STR_CONTACT_NOT_VERIFIED)
	DC_STR_CONTACT_SETUP_CHANGED         = int(C.DC_STR_CONTACT_SETUP_CHANGED)
	DC_STR_ARCHIVEDCHATS                 = int(C.DC_STR_ARCHIVEDCHATS)
	DC_STR_STARREDMSGS                   = int(C.DC_STR_STARREDMSGS)
	DC_STR_AC_SETUP_MSG_SUBJECT          = int(C.DC_STR_AC_SETUP_MSG_SUBJECT)
	DC_STR_AC_SETUP_MSG_BODY             = int(C.DC_STR_AC_SETUP_MSG_BODY)
	DC_STR_SELFTALK_SUBTITLE             = int(C.DC_STR_SELFTALK_SUBTITLE)
	DC_STR_CANNOT_LOGIN                  = int(C.DC_STR_CANNOT_LOGIN)
	DC_STR_SERVER_RESPONSE               = int(C.DC_STR_SERVER_RESPONSE)
	DC_STR_MSGACTIONBYUSER               = int(C.DC_STR_MSGACTIONBYUSER)
	DC_STR_MSGACTIONBYME                 = int(C.DC_STR_MSGACTIONBYME)
	DC_STR_MSGLOCATIONENABLED            = int(C.DC_STR_MSGLOCATIONENABLED)
	DC_STR_MSGLOCATIONDISABLED           = int(C.DC_STR_MSGLOCATIONDISABLED)
	DC_STR_LOCATION                      = int(C.DC_STR_LOCATION)
	DC_STR_STICKER                       = int(C.DC_STR_STICKER)
	DC_STR_DEVICE_MESSAGES               = int(C.DC_STR_DEVICE_MESSAGES)
	DC_STR_COUNT                         = int(C.DC_STR_COUNT)
)
