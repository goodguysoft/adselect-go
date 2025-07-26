package _go

import (
	"encoding/json"
	"fmt"
	"html"
)

func GenerateAd(ad *Ad, adType string, jsFunc string) (string, error) {
	switch adType {
	case AdTypeJsonText:
		return generateJsonTextAd(ad)
	case AdTypeJsonImage:
		return generateJsonImageAd(ad)
	case AdTypeHtmlText:
		return generateHtmlTextAd(ad)
	case AdTypeHtmlImage:
		return generateHtmlImageAd(ad)
	case AdTypeJavaScriptText:
		return generateJavaScriptTextAd(ad, jsFunc)
	case AdTypeJavaScriptImage:
		return generateJavaScriptImageAd(ad, jsFunc)
	case AdTypeBannerMediumRectJson, AdTypeBannerLeaderboardJson, AdTypeBannerWideSkyJson:
		return generateJsonImageAd(ad)
	case AdTypeBannerMediumRectHtml, AdTypeBannerLeaderboardHtml, AdTypeBannerWideSkyHtml:
		return generateHtmlBannerAd(ad)
	default:
		return "", fmt.Errorf("unsupported ad type: %s", adType)
	}
}

type adTextOnly struct {
	ID          int    `json:"id"`
	URL         string `json:"url"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

func generateJsonTextAd(ad *Ad) (string, error) {
	textAd := adTextOnly{
		ID:          ad.ID,
		URL:         ad.URL,
		Title:       ad.Title,
		Description: ad.Description,
	}
	b, err := json.Marshal(textAd)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func generateJsonImageAd(ad *Ad) (string, error) {
	b, err := json.Marshal(ad)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func generateHtmlTextAd(ad *Ad) (string, error) {
	style := `<style>
.ad-box.ad-text {
  border: 1px solid #e0e0e0;
  border-radius: 6px;
  padding: 12px 16px;
  background: #fafbfc;
  max-width: 400px;
  font-family: Arial, sans-serif;
  margin: 12px 0;
}
.ad-label {
  font-size: 12px;
  color: #5f6368;
  margin-bottom: 4px;
  display: block;
  font-weight: bold;
  letter-spacing: 0.5px;
}
.ad-label a {
  color: #5f6368;
  text-decoration: none;
}
.ad-label a:hover {
  text-decoration: underline;
}
.ad-body {
  font-size: 15px;
  color: #202124;
  margin-top: 6px;
}
</style>
`
	htmlStr := fmt.Sprintf(
		`%s<div class="ad-box ad-text">
  <span class="ad-label"><a href="%s" target="_blank" rel="noopener noreferrer">%s</a></span>
  <div class="ad-body">%s</div>
</div>`,
		style,
		html.EscapeString(ad.URL),
		html.EscapeString(ad.Title),
		html.EscapeString(ad.Description),
	)
	return htmlStr, nil
}

func generateHtmlImageAd(ad *Ad) (string, error) {
	style := `<style>
.ad-box.ad-image {
  border: 1px solid #e0e0e0;
  border-radius: 6px;
  padding: 12px 16px;
  background: #fafbfc;
  max-width: 400px;
  font-family: Arial, sans-serif;
  margin: 12px 0;
  overflow: hidden;
  display: flex;
  align-items: stretch;
}
.ad-box.ad-image a {
  display: flex;
  text-decoration: none;
  color: inherit;
  width: 100%;
}
.ad-icon {
  float: left;
  height: 100%;
  max-height: 80px;
  width: 80px;
  object-fit: cover;
  border-radius: 4px;
  margin-right: 14px;
  background: #fff;
  flex-shrink: 0;
  align-self: stretch;
}
.ad-body {
  display: flex;
  flex-direction: column;
  justify-content: center;
  flex: 1;
}
.ad-label {
  font-size: 14px;
  color: #5f6368;
  font-weight: bold;
  margin-bottom: 4px;
  letter-spacing: 0.5px;
}
.ad-body span:last-child {
  font-size: 15px;
  color: #202124;
}
</style>
`
	htmlStr := fmt.Sprintf(
		`%s<div class="ad-box ad-image">
  <a href="%s" target="_blank" rel="noopener noreferrer">
    <img class="ad-icon" src="%s" alt="%s"/>
    <div class="ad-body">
      <span class="ad-label">%s</span>
      <span>%s</span>
    </div>
  </a>
</div>`,
		style,
		html.EscapeString(ad.URL),
		html.EscapeString(ad.ImageURL),
		html.EscapeString(ad.Title),
		html.EscapeString(ad.Title),
		html.EscapeString(ad.Description),
	)
	return htmlStr, nil
}

func generateJavaScriptTextAd(ad *Ad, jsFunc string) (string, error) {
	jsonStr, err := generateJsonTextAd(ad)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s(%s);", jsFunc, jsonStr), nil
}

func generateJavaScriptImageAd(ad *Ad, jsFunc string) (string, error) {
	jsonStr, err := generateJsonImageAd(ad)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s(%s);", jsFunc, jsonStr), nil
}

func generateHtmlBannerAd(ad *Ad) (string, error) {
	htmlStr := fmt.Sprintf(
		`<a href="%s" target="_blank" rel="noopener noreferrer"><img src="%s" alt="%s" style="max-width:100%%;height:auto;display:block;"></a>`,
		html.EscapeString(ad.URL),
		html.EscapeString(ad.ImageURL),
		html.EscapeString(ad.Title),
	)
	return htmlStr, nil
}
