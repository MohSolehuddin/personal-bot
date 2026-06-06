package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"

	"waha-bot/config"
	"waha-bot/models"
)

func parseAmount(amountStr string, transType string) (int, error) {
	amountStr = strings.ToLower(amountStr)
	multiplier := 1.0

	if strings.HasSuffix(amountStr, "k") || strings.HasSuffix(amountStr, "rb") {
		multiplier = 1000.0
		amountStr = strings.TrimRight(amountStr, "krb")
	} else if strings.HasSuffix(amountStr, "jt") || strings.HasSuffix(amountStr, "m") {
		multiplier = 1000000.0
		amountStr = strings.TrimRight(amountStr, "jtm")
	}

	val, err := strconv.ParseFloat(amountStr, 64)
	if err != nil {
		return 0, err
	}
	
	// Convert ke nominal sebenarnya (menangani 1.5 * 1000 = 1500)
	finalAmountFloat := val * multiplier
	finalAmount := int(finalAmountFloat)
	
	// Actual Budget dikali 100 (standar 2 decimal)
	finalAmount = finalAmount * 100

	if transType == "B" {
		finalAmount = -finalAmount
	}

	return finalAmount, nil
}

// SendToActualBudget mengirim data transaksi ke Actual Budget API
func SendToActualBudget(data *TransactionData) error {
	if config.AppConfig.ActualBudgetId == "" || config.AppConfig.ActualAccountId == "" {
		return fmt.Errorf("ACTUAL_BUDGET_ID atau ACTUAL_ACCOUNT_ID belum diatur di .env")
	}

	amount, err := parseAmount(data.Amount, data.Type)
	if err != nil {
		return fmt.Errorf("gagal parsing amount: %w", err)
	}

	dateStr := time.Now().Format("2006-01-02")
	
	// Gabungkan Category dan Description
	notes := data.Category
	if data.Description != "" {
		notes += " - " + data.Description
	}

	reqData := models.ActualTransactionRequest{
		Transactions: []models.ActualTransaction{
			{
				AccountId: config.AppConfig.ActualAccountId,
				Date:      dateStr,
				Amount:    amount,
				Notes:     notes,
			},
		},
	}

	jsonData, err := json.Marshal(reqData)
	if err != nil {
		return fmt.Errorf("gagal marshal JSON: %w", err)
	}

	// Asumsi endpoint menggunakan actual-http-api wrapper
	url := fmt.Sprintf("%s/budgets/%s/transactions", strings.TrimRight(config.AppConfig.ActualBudgetURL, "/"), config.AppConfig.ActualBudgetId)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return fmt.Errorf("gagal membuat request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	if config.AppConfig.ActualBudgetPassword != "" {
		// Sesuaikan header otentikasi berdasarkan wrapper yang digunakan
		req.Header.Set("x-api-key", config.AppConfig.ActualBudgetPassword) 
	}

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("gagal request ke server Actual: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		bodyBytes, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("error server Actual (status %d): %s", resp.StatusCode, string(bodyBytes))
	}

	return nil
}
