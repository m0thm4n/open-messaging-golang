package structs

import (
	"encoding/json"

	"github.com/mypurecloud/platform-client-sdk-go/v77/platformclientv2"
)

type MessageData struct {
	SenderMessageId  string
	OpenMessageApi   *platformclientv2.ConversationsApi
	NumberOfMessages int
	Body             string
	Transcript       []map[string]string
	ID               int
}

type WebhookResponse struct {
	Id                string          `json:"id"`
	Channel           json.RawMessage `json:"channel"`
	Type              string          `json:"type"`
	Text              string          `json:"text"`
	Status            string          `json:"status"`
	OriginatingEntity string          `json:"originatingEntity"`
	Direction         string          `json:"direction`
}
