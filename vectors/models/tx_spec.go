package models

// TxSpec contains transaction data in various formats and with sender and recipient information.
type TxSpec struct {
	Sender         User     `json:"sender"`
	Recipient      User     `json:"recipient"`
	TxID           string   `json:"tx_id"`
	RawHex         string   `json:"raw_hex"`
	BeefHex        string   `json:"beef_hex"`
	BeefV2Hex      string   `json:"beef_v2_hex"`
	AtomicBeefHex  string   `json:"atomic_beef_hex"`
	EfHex          string   `json:"ef_hex"`
	LockingScripts []string `json:"locking_scripts"`
}
