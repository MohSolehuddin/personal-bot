# personal-bot - WhatsApp Budget Bot

**Language:** Go (Golang)  
**Framework:** Fiber v3  
**AI Provider:** Google Gemini  
**Budget System:** Actual Budget API  
**WhatsApp:** WAHA (WhatsApp HTTP API)

---

## Overview

**personal-bot** adalah WhatsApp bot yang mengintegrasikan:
- **AI responses** menggunakan Google Gemini
- **Transaction logging** ke Actual Budget
- **WhatsApp integration** via WAHA

---

## Architecture

```
┌─────────────┐     ┌──────────────────┐     ┌──────────────────┐
│  WhatsApp   │────▶│  personal-bot    │────▶│  Actual Budget   │
│   (WAHA)    │     │   (Go Fiber)     │     │  (Docker)        │
│             │     │                  │     │                  │
│ 3000 (API)  │     │ 3001 (API)       │     │ 5006 (API)       │
└─────────────┘     └──────────────────┘     └──────────────────┘
                        │
                        ▼
                ┌──────────────┐
                │   Gemini AI  │
                │  (API Key)   │
                └──────────────┘
```

---

## Features

### ✅ Implemented
- WhatsApp webhook handling (WAHA)
- AI chat responses (Gemini)
- Transaction parsing (B/J format)
- Actual Budget integration
- Async message processing

### 🔄 In Progress
- Multi-user support
- Transaction history
- Budget sync

### 📋 Proposed
- User authentication
- Expense categorization
- Reports & analytics

---

## Quick Start (Docker)

### 1. Clone & Setup

```bash
cd ~/server-app/personal-bot
cp .env.example .env
```

### 2. Configure `.env`

```env
# WAHA API
WAHA_API_URL=http://172.17.0.1:3000/api/sendText
WAHA_API_KEY=your_waha_api_key

# Actual Budget
ACTUAL_BUDGET_URL=http://actual_budget:5006
ACTUAL_BUDGET_PASSWORD=your_password
ACTUAL_BUDGET_ID=your_budget_id
ACTUAL_ACCOUNT_ID=your_account_id

# Gemini AI
GEMINI_API_KEY=your_gemini_api_key
```

### 3. Run with Docker Compose

```bash
docker-compose up -d
```

---

## Docker Compose (Full Stack)

Jalankan semua services dengan satu perintah:

```bash
cd ~/server-app
docker-compose -f personal-bot/docker-compose.yaml \
               -f integrate-actual-budget-service/docker-compose.yaml \
               -f integrate-ollama/docker-compose.yaml \
               up -d
```

---

## API Endpoints

### Webhook
```
POST /webhook/waha
Content-Type: application/json

{
  "event": "message",
  "payload": {
    "from": "6281234567890@c.us",
    "body": "B Makan Nasi Goreng 15k"
  }
}
```

**Response:**
```json
{
  "status": "success"
}
```

---

## Message Format

### Transactions
- **Expense:** `B [Kategori] [Deskripsi] [Nominal]`
  - Contoh: `B Makan Nasi Goreng 15k`
  - Contoh: `B Transport Ojol 25k`
- **Income:** `J [Kategori] [Deskripsi] [Nominal]`
  - Contoh: `J Gaji Gaji Bulanan 5jt`

### AI Chat
- **Format:** `AI [prompt]`
- **Contoh:** `AI Apa cuaca hari ini?`

---

## Configuration

| Variable | Default | Description |
|----------|---------|-------------|
| `WAHA_API_URL` | `http://172.17.0.1:3000/api/sendText` | WAHA API endpoint |
| `WAHA_API_KEY` | *(required)* | WAHA API key |
| `ACTUAL_BUDGET_URL` | `http://192.168.0.2:5006` | Actual Budget URL |
| `ACTUAL_BUDGET_PASSWORD` | *(required)* | Actual Budget password |
| `ACTUAL_BUDGET_ID` | *(required)* | Budget ID |
| `ACTUAL_ACCOUNT_ID` | *(required)* | Account ID |
| `GEMINI_API_KEY` | *(required)* | Gemini API key |

---

## Project Structure

```
personal-bot/
├── config/          # Configuration (env, structs)
├── controllers/     # Webhook handlers
├── routes/          # API routes
├── services/        # Business logic
│   ├── actualBudgetService.go
│   ├── aiGeminiService.go
│   ├── chatParserService.go
│   ├── messageRouter.go
│   └── wahaMessenger.go
├── models/          # Data models
├── interfaces/      # Interfaces (Messenger, AIProvider)
└── bot.go           # Main entry point
```

---

## Development

### Run Locally

```bash
cd ~/server-app/personal-bot
go mod download
go run bot.go
```

### Build Docker Image

```bash
docker build -t personal-bot:latest .
```

---

## Troubleshooting

### WAHA Not Responding
- Check WAHA is running: `docker ps | grep waha`
- Verify WAHA_API_URL in `.env`
- Check network bridge: `docker network ls`

### Actual Budget Connection
- Verify budget ID and account ID
- Check API password matches
- Test with curl: `curl http://localhost:5006`

### Gemini AI Issues
- Check API key is valid
- Verify quota and billing
- Test: `curl -X POST https://generativelanguage.googleapis.com/v1beta/models:generateContent`

---

## Production Deployment

1. Update `.env` dengan production credentials
2. Run Docker Compose:
   ```bash
   cd ~/server-app
   docker-compose -f personal-bot/docker-compose.yaml up -d
   ```
3. Monitor logs:
   ```bash
   docker-compose -f personal-bot/docker-compose.yaml logs -f
   ```
4. Set up auto-restart (already in docker-compose.yaml)

---

## Next Steps

- [ ] Add multi-user support
- [ ] Implement transaction history
- [ ] Add expense categories
- [ ] Create reports & analytics
- [ ] User authentication
- [ ] Web dashboard

---

**Author:** Moh Solehuddin  
**Last Updated:** 2026-06-06
