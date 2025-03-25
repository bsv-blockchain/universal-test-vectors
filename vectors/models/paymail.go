package models

// Paymail is a struct that holds paymail information.
type Paymail struct {
	Alias      string `json:"alias"`
	Domain     string `json:"domain"`
	Paymail    string `json:"paymail"`
	PublicName string `json:"public_name"`
}
