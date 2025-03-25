package mapper

import (
	"github.com/4chain-ag/universal-test-vectors/pkg/testabilities"
	"github.com/4chain-ag/universal-test-vectors/vectors/models"
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
