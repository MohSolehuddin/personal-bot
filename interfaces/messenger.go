package interfaces

// Messenger mendefinisikan kontrak untuk mengirim pesan balasan ke pengguna
type Messenger interface {
	SendMessage(to string, text string) error
}
