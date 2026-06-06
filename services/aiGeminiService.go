package services

import (
	"context"
	"fmt"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"

	"waha-bot/config"
)

type GeminiService struct{}

func NewGeminiService() *GeminiService {
	return &GeminiService{}
}

func (g *GeminiService) GenerateResponse(prompt string) (string, error) {
	ctx := context.Background()

	if config.AppConfig.GeminiApiKey == "" {
		return "", fmt.Errorf("GEMINI_API_KEY belum dikonfigurasi")
	}

	client, err := genai.NewClient(ctx, option.WithAPIKey(config.AppConfig.GeminiApiKey))
	if err != nil {
		return "", fmt.Errorf("gagal inisialisasi Gemini Client: %w", err)
	}
	defer client.Close()

	// Menggunakan model Gemini 1.5 Flash yang cepat
	model := client.GenerativeModel("gemini-1.5-flash")
	model.SetTemperature(0.7)

	resp, err := model.GenerateContent(ctx, genai.Text(prompt))
	if err != nil {
		return "", fmt.Errorf("gagal generate content: %w", err)
	}

	return extractGeminiText(resp), nil
}

// extractGeminiText mengambil teks pertama yang valid dari respons Gemini
func extractGeminiText(resp *genai.GenerateContentResponse) string {
	if len(resp.Candidates) == 0 {
		return "Maaf, AI tidak memberikan respons."
	}
	
	candidate := resp.Candidates[0]
	if candidate.Content == nil || len(candidate.Content.Parts) == 0 {
		return "Maaf, AI merespons dengan format kosong."
	}

	part := candidate.Content.Parts[0]
	
	if textPart, ok := part.(genai.Text); ok {
		return string(textPart)
	}

	return "Maaf, AI merespons dengan format yang tidak didukung."
}
