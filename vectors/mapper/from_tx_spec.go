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
		TxID:           spec.ID().Hex(),
		RawHex:         spec.RawTX().Hex(),
		BeefHex:        spec.BEEF().Hex(),
		BeefV2Hex:      spec.BEEFv2().Hex(),
		AtomicBeefHex:  spec.AtomicBEEF().Hex(),
		EfHex:          spec.EF().Hex(),
		LockingScripts: spec.LockingScripts(),
	}
}
