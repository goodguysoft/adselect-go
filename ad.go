package _go

const (
	AdTypeText              = "text"
	AdTypeImage             = "image"
	AdTypeBannerMediumRect  = "Medium Rectangle Banner"
	AdTypeBannerLeaderboard = "Leaderboard Banner"
	AdTypeBannerWideSky     = "Wide Skyscraper Banner"
)

// Ad represents an advertisement with basic properties.
type Ad struct {
	ID          int    `json:"id"`
	URL         string `json:"url"`
	Title       string `json:"title"`
	Description string `json:"description"`
	ImageURL    string `json:"imageUrl"`
}
