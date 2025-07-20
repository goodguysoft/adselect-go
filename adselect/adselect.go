package adselect

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
