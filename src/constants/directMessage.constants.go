package constants

const SEND_DIRECT_MESSAGE_ENDPOINT = "/2/dm_conversations/with/%s/messages"

const LOOKUP_DIRECT_MESSAGES = "/2/dm_events"
const DM_EVENT_FIELDS_QUERY = "dm_event.fields"
const DM_EVENT_FIELDS_VALUE = "id,text,event_type,sender_id"

const LOOKUP_DIRECT_MESSAGES_BY_ID = "/2/dm_conversations/with/%s/dm_events"
