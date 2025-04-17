package mapper

import (
	"github.com/bsv-blockchain/universal-test-vectors/pkg/testabilities"
	"github.com/bsv-blockchain/universal-test-vectors/vectors/models"
)

// FromPaymail converts a testabilities.Paymail to JSON-serializable models.Paymail.
func FromPaymail(paymail testabilities.Paymail) models.Paymail {
	return models.Paymail{
		Paymail:    paymail.Address(),
		Alias:      paymail.Alias(),
		Domain:     paymail.Domain(),
		PublicName: paymail.PublicName(),
	}
}
