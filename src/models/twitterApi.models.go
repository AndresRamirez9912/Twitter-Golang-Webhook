package models

type DirectMessageRequestBody struct {
	Text string `json:"text"`
}

type DirectMessageResponse struct {
	Data struct {
		Dm_conversation_id string `json:"dm_conversation_id"`
		Dm_event_id        string `json:"dm_event_id"`
	} `json:"data"`
}
