package mapper

import (
	"github.com/bsv-blockchain/universal-test-vectors/pkg/testabilities"
	"github.com/bsv-blockchain/universal-test-vectors/vectors/models"
)

// FromTxSpec converts a testabilities.TransactionSpec to JSON-serializable models.TxSpec
func FromTxSpec(spec testabilities.TransactionSpec) models.TxSpec {
	return models.TxSpec{
		Sender:         FromUser(spec.Sender()),
		Recipient:      FromUser(spec.Recipient()),
		TxID:           spec.ID(),
		RawHex:         spec.RawTX(),
		BeefHex:        spec.BEEF(),
		EfHex:          spec.EF(),
		LockingScripts: spec.LockingScripts(),
	}
}
