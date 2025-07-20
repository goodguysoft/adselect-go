package adselect

const (
	ChatMessageRoleUser = "User"
	ChatMessageRoleBot  = "Bot"
)

type ChatMessage struct {
	Role string `json:"Role"` // "User" or "Bot"
	Text string `json:"Text"` // Message content
}
