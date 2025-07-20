package _go

const (
	AdTypeText  = "text"
	AdTypeImage = "image"
)

// Ad represents an advertisement with basic properties.
type Ad struct {
	ID          int    `json:"id"`
	URL         string `json:"url"`
	Title       string `json:"title"`
	Description string `json:"description"`
	ImageURL    string `json:"imageUrl"`
}
