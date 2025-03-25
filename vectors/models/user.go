package models

type User struct {
	PrivateKey string `json:"private_key"`
	PublicKey  string `json:"public_key"`
	Address    string `json:"address"`

	XPriv  string `json:"xpriv"`
	XPub   string `json:"xpub"`
	XPubID string `json:"xpub_id"`

	Paymails []Paymail `json:"paymails"`
}
