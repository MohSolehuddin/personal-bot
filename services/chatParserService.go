package services

import (
	"errors"
	"regexp"
	"strings"
)

// TransactionData adalah hasil terjemahan dari chat WA
type TransactionData struct {
	Type        string // "B" atau "J"
	Category    string // "Makan" atau "Gaji" (Jenis)
	Description string // "Ayam geprek" (Detail)
	Amount      string // "3k", "2jt", "1.5k"
	DateTime    string // "13:00" atau kosong
}

// ParseMessage memecah teks WA menjadi data transaksi
func ParseMessage(text string) (*TransactionData, error) {
	// Regex Explanation:
	// 1. ^([bBjJ])            -> B atau J
	// 2. \s+([^\s]+)          -> Kata pertama setelah B/J (Jenis/Kategori)
	// 3. (?:\s+(.*?))?        -> Kata-kata setelah Jenis sebelum nominal (Detail Opsional)
	// 4. \s+(\d+(?:\.\d+)?[a-zA-Z]*) -> Angka nominal, bisa desimal (1.5) dan bisa diakhiri k, jt, rb
	// 5. (?:\s+(\S+))?$       -> Sisa (Waktu Opsional)
	
	pattern := `^([bBjJ])\s+([^\s]+)(?:\s+(.*?))?\s+(\d+(?:\.\d+)?[a-zA-Z]*)(?:\s+(\S+))?$`
	re := regexp.MustCompile(pattern)
	matches := re.FindStringSubmatch(strings.TrimSpace(text))

	if len(matches) == 0 {
		return nil, errors.New("format tidak dikenali")
	}

	data := &TransactionData{
		Type:        strings.ToUpper(matches[1]),
		Category:    strings.TrimSpace(matches[2]),
		Description: strings.TrimSpace(matches[3]),
		Amount:      strings.ToLower(matches[4]),
		DateTime:    matches[5],
	}

	return data, nil
}