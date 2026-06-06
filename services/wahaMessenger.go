package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"waha-bot/config"
)

// WAHAMessenger implementasi dari interfaces.Messenger
type WAHAMessenger struct{}

// NewWAHAMessenger membuat instance WAHAMessenger
func NewWAHAMessenger() *WAHAMessenger {
	return &WAHAMessenger{}
}

// SendMessage mengirim pesan ke nomor WA menggunakan API WAHA
func (w *WAHAMessenger) SendMessage(to string, text string) error {
	url := config.AppConfig.WahaApiUrl

	// Payload untuk WAHA API (format standar WAHA untuk sendText)
	payload := map[string]interface{}{
		"chatId":  to,
		"text":    text,
		"session": "default", // Sesuaikan dengan session ID WAHA Anda jika berbeda
	}

	jsonData, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("gagal marshal pesan WAHA: %w", err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return fmt.Errorf("gagal buat request WAHA: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	if config.AppConfig.WahaApiKey != "" {
		req.Header.Set("X-Api-Key", config.AppConfig.WahaApiKey) // Sesuaikan header auth jika ada
	}

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("gagal mengirim pesan: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		bodyBytes, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("error dari WAHA server (status %d): %s", resp.StatusCode, string(bodyBytes))
	}

	return nil
}
