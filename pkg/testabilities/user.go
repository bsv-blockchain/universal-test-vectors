package testabilities

import (
	"encoding/hex"
	"github.com/4chain-ag/universal-test-vectors/pkg/defs"
	"github.com/4chain-ag/universal-test-vectors/pkg/internal/interpreter"
	bip32 "github.com/bsv-blockchain/go-sdk/compat/bip32"
	ec "github.com/bsv-blockchain/go-sdk/primitives/ec"
	"github.com/bsv-blockchain/go-sdk/script"
	sdk "github.com/bsv-blockchain/go-sdk/transaction"
)

// User is a fixture that is representing a user of the system.
type User struct {
	Paymails []Paymail
	PrivKey  string
}

// XPub returns the xpub of this user.
func (f *User) XPub() string {
	if f.PrivKey == "" {
		return ""
	}
	key, err := bip32.NewKeyFromString(f.PrivKey)
	if err != nil {
		panic("Invalid setup of user fixture, cannot restore xpriv: " + err.Error())
	}
	pubkey, err := key.Neuter()
	if err != nil {
		panic("Invalid setup of user fixture, cannot calculate xpub: " + err.Error())
	}
	return pubkey.String()
}

// XPubID returns hash of the xpub of this user.
func (f *User) XPubID() string {
	xpub := f.XPub()
	if xpub == "" {
		return ""
	}
	return hexHash(xpub)
}

func (f *User) XPrivHD() *bip32.ExtendedKey {
	xpriv, err := bip32.GenerateHDKeyFromString(f.PrivKey)
	if err != nil {
		panic("Invalid setup of user fixture, cannot restore xpriv: " + err.Error())
	}
	return xpriv
}

// PrivateKey returns the private key of this user.
func (f *User) PrivateKey() *ec.PrivateKey {
	priv, err := bip32.GetPrivateKeyFromHDKey(f.XPrivHD())
	if err != nil {
		panic("Invalid setup of user fixture, cannot restore private key: " + err.Error())
	}
	return priv
}

func (f *User) PrivateKeyHex() string {
	b := f.PrivateKey().Serialize()
	return hex.EncodeToString(b)
}

// PublicKey returns the public key of this user.
func (f *User) PublicKey() *ec.PublicKey {
	return f.PrivateKey().PubKey()
}

func (f *User) Address() *script.Address {
	addr, err := script.NewAddressFromPublicKey(f.PublicKey(), true)
	if err != nil {
		panic("Invalid setup of user fixture, cannot restore address: " + err.Error())
	}
	return addr
}

// P2PKHLockingScript returns the locking script of this user.
func (f *User) P2PKHLockingScript(instructions ...defs.CustomInstruction) *script.Script {
	res, err := interpreter.NewLockingScriptInterpreter().
		Process(f.PublicKey(), instructions)

	if err != nil {
		panic("Err returned from LockingScriptInterpreter: " + err.Error())
	}

	return res.LockingScript
}

// P2PKHUnlockingScriptTemplate returns the unlocking script template of this user.
func (f *User) P2PKHUnlockingScriptTemplate(instructions ...defs.CustomInstruction) sdk.UnlockingScriptTemplate {
	res, err := interpreter.NewUnlockingScriptInterpreter().
		Process(f.PrivateKey(), instructions)

	if err != nil {
		panic("Err returned from UnlockingTemplateResolver: " + err.Error())
	}

	return res.Template
}
