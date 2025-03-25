package mapper

import (
	"github.com/4chain-ag/universal-test-vectors/pkg/testabilities"
	"github.com/4chain-ag/universal-test-vectors/vectors/models"
)

func FromTxSpec(spec testabilities.TransactionSpec) models.TxSpec {
	return models.TxSpec{
		Sender:    FromUser(spec.Sender()),
		Recipient: FromUser(spec.Recipient()),
		TxID:      spec.ID(),
		RawHex:    spec.RawTX(),
		BeefHex:   spec.BEEF(),
		EfHex:     spec.EF(),
	}
}
