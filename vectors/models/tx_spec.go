package models

type TxSpec struct {
	Sender    User   `json:"sender"`
	Recipient User   `json:"recipient"`
	TxID      string `json:"tx_id"`
	RawHex    string `json:"raw_hex"`
	BeefHex   string `json:"beef_hex"`
	EfHex     string `json:"ef_hex"`
}
