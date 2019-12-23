package deltachat

const EVENT_TYPES_DATA1_IS_INT = DC_EVENT_CHAT_MODIFIED |
	DC_EVENT_CONFIGURE_PROGRESS |
	DC_EVENT_CONTACTS_CHANGED |
	DC_EVENT_ERROR_NETWORK |
	DC_EVENT_IMEX_PROGRESS |
	DC_EVENT_INCOMING_MSG |
	DC_EVENT_LOCATION_CHANGED |
	DC_EVENT_MSG_DELIVERED |
	DC_EVENT_MSG_FAILED |
	DC_EVENT_MSG_READ |
	DC_EVENT_MSGS_CHANGED |
	DC_EVENT_SECUREJOIN_INVITER_PROGRESS |
	DC_EVENT_SECUREJOIN_JOINER_PROGRESS

const EVENT_TYPES_DATA2_IS_INT = DC_EVENT_INCOMING_MSG |
	DC_EVENT_MSG_DELIVERED |
	DC_EVENT_MSG_FAILED |
	DC_EVENT_SECUREJOIN_INVITER_PROGRESS |
	DC_EVENT_SECUREJOIN_JOINER_PROGRESS |
	DC_EVENT_MSG_READ |
	DC_EVENT_MSGS_CHANGED

const EVENT_TYPES_ERROR = DC_EVENT_ERROR |
	DC_EVENT_ERROR_NETWORK |
	DC_EVENT_ERROR_SELF_NOT_IN_GROUP

const (
	DATA_TYPE_INT uint8 = iota
	DATA_TYPE_STRING
	DATA_TYPE_NIL
)

var errorTypeNames = map[int]string{
	DC_EVENT_ERROR:                   "DC_EVENT_ERROR",
	DC_EVENT_ERROR_NETWORK:           "DC_EVENT_ERROR_NETWORK",
	DC_EVENT_ERROR_SELF_NOT_IN_GROUP: "DC_EVENT_ERROR_SELF_NOT_IN_GROUP",
}

var dataTypeNames = map[uint8]string{
	DATA_TYPE_INT:    "int",
	DATA_TYPE_STRING: "string",
	DATA_TYPE_NIL:    "nil",
}

func Data1TypeForEvent(event int) uint8 {
	if event == DC_EVENT_IMEX_FILE_WRITTEN || event == DC_EVENT_FILE_COPIED {
		return DATA_TYPE_STRING
	}

	if (EVENT_TYPES_DATA1_IS_INT & event) == event {
		return DATA_TYPE_INT
	}

	return DATA_TYPE_NIL
}

func Data2TypeForEvent(event int) uint8 {
	if event >= 100 && event <= 499 {
		return DATA_TYPE_STRING
	}

	if (EVENT_TYPES_DATA2_IS_INT & event) == event {
		return DATA_TYPE_INT
	}

	return DATA_TYPE_NIL
}

type Event struct {
	EventType int
	Data1     Data
	Data2     Data
}
