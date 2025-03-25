package models

type Paymail struct {
	Alias      string `json:"alias"`
	Domain     string `json:"domain"`
	Paymail    string `json:"paymail"`
	PublicName string `json:"public_name"`
}
