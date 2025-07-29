# adselect-go

**Detailed documentation is available on the wiki:** [http://wiki.vektorone.goodguysoft.com/](http://wiki.vektorone.goodguysoft.com/)

The Go SDK from **vektorone.co** lets you integrate intelligent ad generation into your Go-based chat or web systems. It enables you to:

- Send full chat history
- Retrieve contextual ads based on conversations or web page content
- Render ads in multiple formats (JSON, HTML, JavaScript)

All SDK functions use idiomatic Go error handling.

## Installation

Install the SDK:

```bash
go get github.com/goodguysoft/adselect-go
```

Import in your application:

```go
import adselect "github.com/goodguysoft/adselect-go"
```

## Sending Chat History

Use `SendChatHistory()` to send complete conversation data to the API.

**Function Signature:**
```go
func SendChatHistory(apiId string, apiKey string, userID string, conversationID string, messages []adselect.ChatMessage) error
```

**Parameters:**
- `apiId` — Your VektorOne API ID (provided by support)
- `apiKey` — Your VektorOne API Key
- `userID` — Unique identifier for the human user receiving the ad
- `conversationID` — Unique identifier for the chat session
- `messages` — Slice of message objects containing roles and text

**Message Format:**
```go
type ChatMessage struct {
    Role string // "User" or "Bot"
    Text string // Message content
}
```
- `Role` — `"User"` (end user) or `"Bot"` (assistant, system, or agent)
- `Text` — Message text content

**Example:**
```go
messages := []adselect.ChatMessage{
    {Role: "User", Text: "I'm looking for dog food."},
    {Role: "Bot", Text: "We have natural dog food options."},
}

err := adselect.SendChatHistory("YOUR_API_ID", "YOUR_API_KEY", "user-123", "conv-456", messages)
if err != nil {
    log.Fatalf("failed to send chat history: %v", err)
}
```

## Getting Ad for a Chat Session

Use `GetChatAd()` to fetch an ad based on chat context.

**Function Signature:**
```go
func GetChatAd(apiId string, apiKey string, userID string, conversationID string, adType string, jsFunc string) (string, error)
```

**Parameters:**
- `apiId`, `apiKey` — Credentials provided by VektorOne support
- `userID` — Unique user identifier
- `conversationID` — Current chat session ID
- `adType` — Ad format:
  - `adselect.AdTypeJsonText`
  - `adselect.AdTypeJsonImage`
  - `adselect.AdTypeHtmlTextAd`
  - `adselect.AdTypeHtmlImageAd`
  - `adselect.AdTypeJavaScriptText`
  - `adselect.AdTypeJavaScriptImage`
  - `adselect.AdTypeBannerMediumRectJson`
  - `adselect.AdTypeBannerMediumRectHtml`
  - `adselect.AdTypeBannerLeaderboardJson`
  - `adselect.AdTypeBannerLeaderboardHtml`
  - `adselect.AdTypeBannerWideSkyJson`
  - `adselect.AdTypeBannerWideSkyHtml`
- `jsFunc` — JavaScript callback name (for `JavaScript` ad type only)

**Example:**
```go
adHtml, err := adselect.GetChatAd("YOUR_API_ID", "YOUR_API_KEY", "user-123", "conv-456", adselect.AdTypeHtmlTextAd, "")
if err != nil {
    log.Printf("failed to get ad: %v", err)
} else {
    fmt.Println("Ad HTML:", adHtml)
}
```

## Getting Ad for a Web Page

Use `GetPageAd()` to retrieve an ad based on HTML page content.

**Function Signature:**
```go
func GetPageAd(apiId string, apiKey string, pageContent string, adType adsdk.AdType, jsFunc string) (string, error)
```

**Parameters:**
- `apiId`, `apiKey` — Your API credentials
- `pageContent` — Raw HTML or text content of the page
- `adType` — Desired ad output format:
  - `adsdk.AdTypeJsonText`
  - `adsdk.AdTypeJsonImage`
  - `adsdk.AdTypeHtmlTextAd`
  - `adsdk.AdTypeHtmlImageAd`
  - `adsdk.AdTypeJavaScriptText`
  - `adsdk.AdTypeJavaScriptImage`
  - `adsdk.AdTypeBannerMediumRectJson`
  - `adsdk.AdTypeBannerMediumRectHtml`
  - `adsdk.AdTypeBannerLeaderboardJson`
  - `adsdk.AdTypeBannerLeaderboardHtml`
  - `adsdk.AdTypeBannerWideSkyJson`
  - `adsdk.AdTypeBannerWideSkyHtml`
- `jsFunc` — JavaScript callback name (for `JavaScript` ads only)

**Example:**
```go
html := "<html><body><h1>Best dog food 2024</h1><p>Our guide covers natural ingredients...</p></body></html>"

adScript, err := adsdk.GetPageAd("YOUR_API_ID", "YOUR_API_KEY", html, adsdk.AdTypeJavaScript, "renderAd")
if err != nil {
    log.Printf("ad error: %v", err)
} else {
    fmt.Println("<script>" + adScript + "</script>")
}
```

## AdType Constants

```go
const (
    AdTypeJsonText              = "JsonText"
    AdTypeJsonImage             = "JsonImage"
    AdTypeHtmlText              = "HtmlTextAd"
    AdTypeHtmlImage             = "HtmlImageAd"
    AdTypeJavaScriptText        = "JavaScriptTextAd"
    AdTypeJavaScriptImage       = "JavaScriptImageAd"
    AdTypeBannerMediumRectJson  = "BannerMediumRectJson"
    AdTypeBannerMediumRectHtml  = "BannerMediumRectHtml"
    AdTypeBannerLeaderboardJson = "BannerLeaderboardJson"
    AdTypeBannerLeaderboardHtml = "BannerLeaderboardHtml"
    AdTypeBannerWideSkyJson     = "BannerWideSkyJson"
    AdTypeBannerWideSkyHtml     = "BannerWideSkyHtml"
)
```

## Error Handling

All functions return a standard Go `error` type.

Always validate return values using:
```go
if err != nil {
    // handle failure
}
```

## See Also

- [Getting Started](http://wiki.vektorone.goodguysoft.com/doku.php?id=getting_started)
- [API Reference](http://wiki.vektorone.goodguysoft.com/doku.php?id=api_reference)
- [Ad Formats](http://wiki.vektorone.goodguysoft.com/doku.php?id=ad_formats)
- [Use Cases](http://wiki.vektorone.goodguysoft.com/doku.php?id=use_cases)
