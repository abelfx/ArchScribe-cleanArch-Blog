package infrastructure

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"time"
)

type MistralAIService struct{}

func NewMistralAIService() *MistralAIService {
    return &MistralAIService{}
}

func (s *MistralAIService) Suggest(prompt string) (string, error) {
   
    apiKey := os.Getenv("M_API_KEY")
    if apiKey == "" {
        return "(AI not configured) Try writing: Intro, Body, Conclusion.", nil
    }

    reqBody := map[string]interface{}{
        "model": "mistral-medium",
        "messages": []map[string]string{
            {"role": "user", "content": prompt},
        },
        "max_tokens": 500,
    }
    b, _ := json.Marshal(reqBody)

    client := &http.Client{Timeout: 15 * time.Second}
    req, err := http.NewRequest("POST", "https://api.mistral.ai/v1/chat/completions", bytes.NewBuffer(b))
    if err != nil {
        return "", err
    }
    req.Header.Set("Authorization", "Bearer "+apiKey)
    req.Header.Set("Content-Type", "application/json")

    resp, err := client.Do(req)
    if err != nil {
        return "", err
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        return "", fmt.Errorf("AI provider returned status %d", resp.StatusCode)
    }

    var parsed struct {
        Choices []struct {
            Message struct {
                Content string `json:"content"`
            } `json:"message"`
        } `json:"choices"`
    }
    if err := json.NewDecoder(resp.Body).Decode(&parsed); err != nil {
        return "", err
    }

    if len(parsed.Choices) == 0 || parsed.Choices[0].Message.Content == "" {
        return "", errors.New("no AI response")
    }

    return parsed.Choices[0].Message.Content, nil
}
