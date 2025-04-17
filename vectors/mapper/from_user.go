package mapper

import (
	"slices"

	"github.com/bsv-blockchain/universal-test-vectors/pkg/testabilities"
	"github.com/bsv-blockchain/universal-test-vectors/vectors/models"
)

// FromUser converts a testabilities.User to JSON-serializable models.User.
func FromUser(user testabilities.User) models.User {
	return models.User{
		Paymails: slices.Collect(func(yield func(models.Paymail) bool) {
			for _, paymail := range user.Paymails {
				yield(FromPaymail(paymail))
			}
		}),
		PrivateKey: user.PrivateKeyHex(),
		PublicKey:  user.PublicKey().ToDERHex(),
		XPriv:      user.PrivKey,
		XPub:       user.XPub(),
		Address:    user.Address().AddressString,
	}
}
