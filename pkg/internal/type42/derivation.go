package type42

import (
	"fmt"

	primitives "github.com/bsv-blockchain/go-sdk/primitives/ec"
)

var (
	anyonePriv, _ = primitives.PrivateKeyFromBytes([]byte{1})
	anyonePub     = anyonePriv.PubKey()
)

// Derive creates derived public key based on derivation key (type 42 derivation with "anyone" public key)
func Derive(pubKey *primitives.PublicKey, derivationKey string) (*primitives.PublicKey, error) {
	if pubKey == nil {
		return nil, fmt.Errorf("public key is nil")
	}
	derivedPubByRef, err := pubKey.DeriveChild(anyonePriv, derivationKey)
	if err != nil {
		return nil, fmt.Errorf("failed to derive public key: %w", err)
	}
	return derivedPubByRef, nil
}

// DerivePrivateKey created derived private key based on derivation key (type 42 derivation with "anyone" private key)
func DerivePrivateKey(priv *primitives.PrivateKey, derivationKey string) (*primitives.PrivateKey, error) {
	derived, err := priv.DeriveChild(anyonePub, derivationKey)
	if err != nil {
		return nil, fmt.Errorf("failed to derive private key: %w", err)
	}
	return derived, nil
}
