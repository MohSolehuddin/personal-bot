package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	ActualBudgetURL      string
	ActualBudgetPassword string
	ActualBudgetId       string
	ActualAccountId      string
	
	WahaApiUrl           string
	WahaApiKey           string
	GeminiApiKey         string
}

var AppConfig Config

// LoadConfig memuat environment variables dari .env
func LoadConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Peringatan: File .env tidak ditemukan, menggunakan variabel lingkungan system")
	}

	AppConfig = Config{
		ActualBudgetURL:      getEnv("ACTUAL_BUDGET_URL", "http://192.168.0.2:5006"), // Default ke server yang diberikan
		ActualBudgetPassword: os.Getenv("ACTUAL_BUDGET_PASSWORD"),
		ActualBudgetId:       os.Getenv("ACTUAL_BUDGET_ID"),
		ActualAccountId:      os.Getenv("ACTUAL_ACCOUNT_ID"),

		WahaApiUrl:           getEnv("WAHA_API_URL", "http://172.17.0.1:3000/api/sendText"),
		WahaApiKey:           os.Getenv("WAHA_API_KEY"),
		GeminiApiKey:         os.Getenv("GEMINI_API_KEY"),
	}
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
