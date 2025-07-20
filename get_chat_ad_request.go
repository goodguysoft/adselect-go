package _go

type GetChatAdRequest struct {
	ApiID          string `json:"apiId"`
	ApiKey         string `json:"apiKey"`
	UserID         string `json:"userID"`
	ConversationID string `json:"conversationID"`
	Type           string `json:"type"`
}
