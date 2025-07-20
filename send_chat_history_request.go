package _go

type SendChatHistoryRequest struct {
	ApiID          string        `json:"apiId"`
	ApiKey         string        `json:"apiKey"`
	UserID         string        `json:"userID"`
	ConversationID string        `json:"conversationID"`
	Messages       []ChatMessage `json:"messages"`
}
