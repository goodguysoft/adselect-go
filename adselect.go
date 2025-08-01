package _go

import (
	"encoding/json"
	"fmt"
)

func SendChatHistory(apiId string, apiKey string, userID string, conversationID string, messages []ChatMessage) error {
	payload := SendChatHistoryRequest{
		ApiID:          apiId,
		ApiKey:         apiKey,
		UserID:         userID,
		ConversationID: conversationID,
		Messages:       messages,
	}
	_, err := PostAPIRequest("/chatMessages", payload)
	return err
}

func GetChatAd(apiId string, apiKey string, userID string, conversationID string, adType string, jsFunc string) (string, error) {
	var reqType string
	switch adType {
	case AdTypeJsonText, AdTypeHtmlText, AdTypeJavaScriptText:
		reqType = AdTypeText
	case AdTypeJsonImage, AdTypeHtmlImage, AdTypeJavaScriptImage:
		reqType = AdTypeImage
	case AdTypeBannerMediumRectJson, AdTypeBannerMediumRectHtml:
		reqType = AdTypeBannerMediumRect
	case AdTypeBannerLeaderboardJson, AdTypeBannerLeaderboardHtml:
		reqType = AdTypeBannerLeaderboard
	case AdTypeBannerWideSkyJson, AdTypeBannerWideSkyHtml:
		reqType = AdTypeBannerWideSky
	default:
		return "", fmt.Errorf("unsupported adType: %s", adType)
	}

	payload := GetChatAdRequest{
		ApiID:          apiId,
		ApiKey:         apiKey,
		UserID:         userID,
		ConversationID: conversationID,
		Type:           reqType,
	}

	data, err := PostAPIRequest("/chatAd", payload)
	if err != nil {
		return "", fmt.Errorf("API request failed: %w", err)
	}

	adBytes, err := json.Marshal(data)
	if err != nil {
		return "", fmt.Errorf("failed to marshal ad data: %w", err)
	}

	var ad Ad
	if err := json.Unmarshal(adBytes, &ad); err != nil {
		return "", fmt.Errorf("failed to unmarshal ad: %w", err)
	}

	return GenerateAd(&ad, adType, jsFunc)
}

func GetPageAd(apiId string, apiKey string, pageContent string, adType string, jsFunc string) (string, error) {
	var reqType string
	switch adType {
	case AdTypeJsonText, AdTypeHtmlText, AdTypeJavaScriptText:
		reqType = AdTypeText
	case AdTypeJsonImage, AdTypeHtmlImage, AdTypeJavaScriptImage:
		reqType = AdTypeImage
	case AdTypeBannerMediumRectJson, AdTypeBannerMediumRectHtml:
		reqType = AdTypeBannerMediumRect
	case AdTypeBannerLeaderboardJson, AdTypeBannerLeaderboardHtml:
		reqType = AdTypeBannerLeaderboard
	case AdTypeBannerWideSkyJson, AdTypeBannerWideSkyHtml:
		reqType = AdTypeBannerWideSky
	default:
		return "", fmt.Errorf("unsupported adType: %s", adType)
	}

	payload := GetPageAdRequest{
		ApiID:        apiId,
		ApiKey:       apiKey,
		Type:         reqType,
		PageContents: pageContent,
	}

	data, err := PostAPIRequest("/pageAd", payload)
	if err != nil {
		return "", fmt.Errorf("API request failed: %w", err)
	}

	adBytes, err := json.Marshal(data)
	if err != nil {
		return "", fmt.Errorf("failed to marshal ad data: %w", err)
	}

	var ad Ad
	if err := json.Unmarshal(adBytes, &ad); err != nil {
		return "", fmt.Errorf("failed to unmarshal ad: %w", err)
	}

	return GenerateAd(&ad, adType, jsFunc)
}
