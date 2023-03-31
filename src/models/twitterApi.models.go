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

type LookUpDirectMessageResponse struct {
	Data []dataUpDirectMessageResponse `json:"data"`
	Meta struct {
		Result_count int `json:"result_count"`
	} `json:"meta"`
}

type dataUpDirectMessageResponse struct {
	Text               string `json:"text"`
	Id                 string `json:"id"`
	Event_type         string `json:"event_type"`
	Created_at         string `json:"created_at"`
	Sender_id          string `json:"sender_id"`
	Dm_conversation_id string `json:"dm_conversation_id"`
}
