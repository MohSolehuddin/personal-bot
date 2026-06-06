package interfaces

// AIProvider mendefinisikan kontrak untuk berinteraksi dengan AI
type AIProvider interface {
	GenerateResponse(prompt string) (string, error)
}
