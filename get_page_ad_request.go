package _go

// GetPageAdRequest represents a request to get an ad for a page.
type GetPageAdRequest struct {
	ApiID        string `json:"apiId"`
	ApiKey       string `json:"apiKey"`
	Type         string `json:"type"`
	PageContents string `json:"pageContents"` // Page HTML or plain text contents
}
