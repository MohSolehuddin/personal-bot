package services

import (
	"fmt"
	"strings"

	"waha-bot/interfaces"
)

// MessageRouter menghubungkan webhook dengan logika bisnis
type MessageRouter struct {
	messenger  interfaces.Messenger
	aiProvider interfaces.AIProvider
}

// NewMessageRouter membuat instance router baru
func NewMessageRouter(m interfaces.Messenger, ai interfaces.AIProvider) *MessageRouter {
	return &MessageRouter{
		messenger:  m,
		aiProvider: ai,
	}
}

// Process mengarahkan teks ke handler yang tepat
func (r *MessageRouter) Process(senderID string, text string) {
	text = strings.TrimSpace(text)
	if text == "" {
		return
	}

	upperText := strings.ToUpper(text)

	// 1. Cek apakah ini chat AI
	if strings.HasPrefix(upperText, "AI ") {
		prompt := strings.TrimSpace(text[3:])
		r.handleAI(senderID, prompt)
		return
	}

	// 2. Cek apakah ini transaksi (B/J)
	if strings.HasPrefix(upperText, "B ") || strings.HasPrefix(upperText, "J ") {
		r.handleTransaction(senderID, text)
		return
	}

	// Jika bukan awalan yang dikenali, bisa diabaikan agar tidak spam
	fmt.Printf("Abaikan pesan: bukan format perintah bot\n")
}

func (r *MessageRouter) handleAI(senderID string, prompt string) {
	if r.aiProvider == nil {
		_ = r.messenger.SendMessage(senderID, "Maaf, layanan AI saat ini belum tersedia.")
		return
	}

	reply, err := r.aiProvider.GenerateResponse(prompt)
	if err != nil {
		fmt.Printf("[ERROR AI] %v\n", err)
		err = r.messenger.SendMessage(senderID, "Maaf, terjadi kesalahan saat memproses permintaan AI.")
		if err != nil {
			fmt.Printf("[ERROR BALAS WAHA] %v\n", err)
		}
		return
	}

	err = r.messenger.SendMessage(senderID, reply)
	if err != nil {
		fmt.Printf("[ERROR BALAS WAHA] %v\n", err)
	}
}

func (r *MessageRouter) handleTransaction(senderID string, text string) {
	parsedData, err := ParseMessage(text)
	if err != nil {
		r.messenger.SendMessage(senderID, "Format salah. Contoh pengeluaran: B Makan Nasi_Goreng 15k")
		return
	}

	err = SendToActualBudget(parsedData)
	if err != nil {
		fmt.Printf("[ERROR BUDGET] %v\n", err)
		r.messenger.SendMessage(senderID, "Gagal mencatat transaksi ke Actual Budget.")
		return
	}

	// Buat pesan sukses
	jenis := "Pengeluaran"
	if parsedData.Type == "J" {
		jenis = "Pemasukan"
	}
	
	msg := fmt.Sprintf("✅ Sukses mencatat %s!\nKategori: %s\nDetail: %s\nNominal: %s", 
		jenis, parsedData.Category, parsedData.Description, parsedData.Amount)

	err = r.messenger.SendMessage(senderID, msg)
	if err != nil {
		fmt.Printf("[ERROR BALAS WAHA] %v\n", err)
	}
}
