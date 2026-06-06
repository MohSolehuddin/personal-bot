package models

// ActualTransactionRequest struktur request ke REST wrapper Actual Budget
type ActualTransactionRequest struct {
	Transactions []ActualTransaction `json:"transactions"`
}

type ActualTransaction struct {
	AccountId string `json:"account_id"`
	Date      string `json:"date"`   // Format: YYYY-MM-DD
	Amount    int    `json:"amount"` // Integer, biasanya dikali 100 dari nilai asli (tergantung konfigurasi currency)
	PayeeName string `json:"payee_name,omitempty"`
	Notes     string `json:"notes,omitempty"`
}
