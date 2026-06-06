package controllers

import (
	"fmt"
	"log"
	"strings"

	"waha-bot/models"   // Import struct dari folder models
	"waha-bot/services" // <-- INI IMPORT UNTUK SERVICES PARSER

	"github.com/gofiber/fiber/v3"
)

// WebhookHandler adalah fungsi utama yang memproses pesan dari WAHA
func WebhookHandler(c fiber.Ctx) error {
	fmt.Println("\n========== WEBHOOK MASUK ==========")

	rawBody := c.Body()
	fmt.Printf("RAW JSON PAYLOAD:\n%s\n\n", string(rawBody))

	req := new(models.WAHAPayload)

	if err := c.Bind().JSON(req); err != nil {
		log.Println("Gagal memproses JSON:", err)
		return c.SendStatus(fiber.StatusBadRequest)
	}

	teksPesan := strings.TrimSpace(req.Payload.Body)
	fmt.Printf("Event: %s | Pesan dari %s: %s\n", req.Event, req.Payload.From, teksPesan)

	if req.Event != "message" {
		return c.SendStatus(fiber.StatusOK)
	}

	// Matikan dulu filter nomor kalau lagi iseng ngetes dari banyak nomor
	// allowedIDs := map[string]bool{
	// 	"6281234567890@c.us":  true,
	// 	"266133823291556@lid": true,
	// }
	// senderID := req.Payload.From
	// if req.Payload.FromMe || !allowedIDs[senderID] {
	// 	fmt.Printf("-> Pesan dari %s diabaikan\n", senderID)
	// 	return c.SendStatus(fiber.StatusOK)
	// }

	fmt.Printf("-> Mencoba memproses pesan: %s\n", teksPesan)

	// Inisialisasi interface untuk platform WAHA dan provider AI
	wahaMessenger := services.NewWAHAMessenger()
	geminiAI := services.NewGeminiService()

	// Panggil Router untuk memproses pesan
	router := services.NewMessageRouter(wahaMessenger, geminiAI)
	
	// Kita jalankan Process secara asynchronous (goroutine) 
	// agar webhook bisa langsung membalas 200 OK ke WAHA tanpa menunggu AI/Actual Budget
	go router.Process(req.Payload.From, teksPesan)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "success",
	})
}