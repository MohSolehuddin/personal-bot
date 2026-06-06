package models

type WAHAPayload struct {
	Event   string `json:"event"`
	Session string `json:"session"`
	Payload struct {
		ID     string `json:"id"`
		From   string `json:"from"`
		Body   string `json:"body"`
		FromMe bool   `json:"fromMe"`
	} `json:"payload"`
}