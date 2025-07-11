package testabilities

import (
	"fmt"

	"github.com/bsv-blockchain/go-sdk/chainhash"
	"github.com/bsv-blockchain/go-sdk/script"
	trx "github.com/bsv-blockchain/go-sdk/transaction"
	"github.com/bsv-blockchain/universal-test-vectors/pkg/defs"
)

// grandparentTXIDs are used to indicate prevTXID for parentTXs(source transactions)
// [grandparentTX] -> [parentTX] -> [tx]
// tx is the actual transaction that is being created
// parentTX contains merkle proof
// parentOfSource txs are just placeholders (IDs only)
var grandparentTXIDs = []string{
	"56acf2ec5ae4906e8814369849f60d07ba339671e57f52200ab473e183463f2e",
	"6b0fc7403ffa214357f0326224903e612acf5c3fc5b88dfc175a2be81e343609",
	"6bed3bef2bb8b41289bca3e7f92fa5e7714decb404590ccbbc6ff7dcabf0c725",
	"e4fb03a7eae2f3766f76d52719633a65c7882c7734ae7afe603e97d193f42c0e",
	"53e784cb876751e114c9e4c1921240f184b6dff8d167715227e3708d0f7bb26d",
	"27a53423aa3e5d5c46bf30be53a9998dd247daf758847f244f82d430be71de6e",
}

// TransactionSpec is a builder for creating MOCK! transactions
type TransactionSpec interface {
	WithSender(sender User) TransactionSpec
	WithRecipient(recipient User) TransactionSpec
	WithoutSigning() TransactionSpec
	WithInput(satoshis uint64) TransactionSpec
	WithInputFromUTXO(tx *trx.Transaction, vout uint32, customInstructions ...defs.CustomInstruction) TransactionSpec
	WithSingleSourceInputs(satoshis ...uint64) TransactionSpec
	WithOPReturn(dataStr string) TransactionSpec
	WithOutputScriptParts(parts ...ScriptPart) TransactionSpec
	WithOutputScript(satoshis uint64, script *script.Script) TransactionSpec
	WithP2PKHOutput(satoshis uint64) TransactionSpec

	TX() *trx.Transaction
	InputSourceTX(inputID int) *trx.Transaction
	ID() *chainhash.Hash
	BEEF() Hexer
	BEEFv2() Hexer
	AtomicBEEF() Hexer
	RawTX() Hexer
	EF() Hexer
	LockingScripts() []string

	Sender() User
	Recipient() User
}

type Hexer interface {
	Bytes() []byte
	Hex() string
}

func GivenTX() TransactionSpec {
	spec := &txSpec{
		sourceTransactions: make(map[string]*trx.Transaction),
		sender:             Alice,
		recipient:          Bob,
		blockHeight:        1000,
		grandparentTXIndex: 0,
	}

	return spec
}

type txSpec struct {
	utxos          []*trx.UTXO
	outputs        []*trx.TransactionOutput
	disableSigning bool

	sourceTransactions map[string]*trx.Transaction
	sender             User
	recipient          User

	blockHeight        uint32
	grandparentTXIndex int
}

// WithSender sets the sender for the transaction (default is Sender)
func (spec *txSpec) WithSender(sender User) TransactionSpec {
	spec.sender = sender
	return spec
}

// WithRecipient sets the recipient for the transaction (default is RecipientInternal)
func (spec *txSpec) WithRecipient(recipient User) TransactionSpec {
	spec.recipient = recipient
	return spec
}

// Sender returns the sender for the transaction
func (spec *txSpec) Sender() User {
	return spec.sender
}

// Recipient returns the recipient for the transaction
func (spec *txSpec) Recipient() User {
	return spec.recipient
}

// WithoutSigning disables signing of the transaction (default is false)
func (spec *txSpec) WithoutSigning() TransactionSpec {
	spec.disableSigning = true
	return spec
}

// WithInput adds an input to the transaction with the specified satoshis
// it automatically creates a parent tx (sourceTX) with P2PKH UTXO with provided satoshis
func (spec *txSpec) WithInput(satoshis uint64) TransactionSpec {
	return spec.WithSingleSourceInputs(satoshis)
}

func (spec *txSpec) WithInputFromUTXO(tx *trx.Transaction, vout uint32, customInstructions ...defs.CustomInstruction) TransactionSpec {
	output := tx.Outputs[vout]
	utxo, err := trx.NewUTXO(tx.TxID().String(), vout, output.LockingScript.String(), output.Satoshis)
	spec.requireNoError(err, "creating utxo for input")

	utxo.UnlockingScriptTemplate = spec.sender.P2PKHUnlockingScriptTemplate(customInstructions...)

	spec.utxos = append(spec.utxos, utxo)
	spec.sourceTransactions[tx.TxID().String()] = tx
	return spec
}

// WithSingleSourceInputs adds inputs to the transaction with the specified satoshis
// All the inputs will be sourced from a single parent transaction
func (spec *txSpec) WithSingleSourceInputs(satoshis ...uint64) TransactionSpec {
	sourceTX := spec.makeParentTX(satoshis...)
	for i, s := range satoshis {
		utxo, err := trx.NewUTXO(sourceTX.TxID().String(), uint32(i), spec.sender.P2PKHLockingScript().String(), s)
		spec.requireNoError(err, "creating utxo for input")

		utxo.UnlockingScriptTemplate = spec.sender.P2PKHUnlockingScriptTemplate()

		spec.utxos = append(spec.utxos, utxo)
	}
	spec.sourceTransactions[sourceTX.TxID().String()] = sourceTX

	return spec
}

// WithOPReturn adds an OP_RETURN output to the transaction with the specified data
func (spec *txSpec) WithOPReturn(dataStr string) TransactionSpec {
	data := []byte(dataStr)
	o, err := trx.CreateOpReturnOutput([][]byte{data})
	spec.requireNoError(err, "creating op return")

	spec.outputs = append(spec.outputs, o)

	return spec
}

// WithP2PKHOutput adds a P2PKH output to the transaction with the specified satoshis owned by the recipient
func (spec *txSpec) WithP2PKHOutput(satoshis uint64) TransactionSpec {
	spec.outputs = append(spec.outputs, &trx.TransactionOutput{
		Satoshis:      satoshis,
		LockingScript: spec.recipient.P2PKHLockingScript(),
	})
	return spec
}

// WithOutputScript adds an output to the transaction with the specified satoshis and script
func (spec *txSpec) WithOutputScript(satoshis uint64, script *script.Script) TransactionSpec {
	spec.outputs = append(spec.outputs, &trx.TransactionOutput{
		Satoshis:      satoshis,
		LockingScript: script,
	})
	return spec
}

// ScriptPart is an interface for building script parts
type ScriptPart interface {
	Append(s *script.Script) error
}

// OpCode is an alias for byte to represent an opcode and implements ScriptPart for script building
type OpCode byte

// Append appends the opcode to the script
func (op OpCode) Append(s *script.Script) error {
	return s.AppendOpcodes(script.OpRETURN)
}

// PushData is an alias for []byte to represent push data and implements ScriptPart for script building
type PushData []byte

// Append appends the push data to the script
func (data PushData) Append(s *script.Script) error {
	return s.AppendPushData(data)
}

// WithOutputScriptParts adds an output to the transaction with the specified script parts
func (spec *txSpec) WithOutputScriptParts(parts ...ScriptPart) TransactionSpec {
	s := &script.Script{}
	for _, part := range parts {
		err := part.Append(s)
		spec.requireNoError(err, "appending script part")
	}
	spec.outputs = append(spec.outputs, &trx.TransactionOutput{LockingScript: s})
	return spec
}

// TX creates a go-sdk transaction from the given spec
func (spec *txSpec) TX() *trx.Transaction {
	tx := trx.NewTransaction()
	err := tx.AddInputsFromUTXOs(spec.utxos...)
	spec.requireNoError(err, "adding utxo inputs")

	for _, output := range spec.outputs {
		tx.AddOutput(output)
	}

	for _, input := range tx.Inputs {
		if sourceTX := spec.sourceTransactions[input.SourceTXID.String()]; sourceTX != nil {
			input.SourceTransaction = sourceTX
		}
	}

	if !spec.disableSigning {
		err = tx.Sign()
		spec.requireNoError(err, "signing transaction")
	}
	return tx
}

// InputSourceTX returns the source transaction for the input with the specified index
func (spec *txSpec) InputSourceTX(inputID int) *trx.Transaction {
	return spec.sourceTransactions[spec.utxos[inputID].TxID.String()]
}

// ID returns the transaction ID
func (spec *txSpec) ID() *chainhash.Hash {
	return spec.TX().TxID()
}

// BEEF returns the BEEF hex of the transaction
func (spec *txSpec) BEEF() Hexer {
	tx := spec.TX()
	beef, err := tx.BEEF()
	spec.requireNoError(err, "beef transaction")

	return newHexer(beef)
}

// BEEFv2 returns a Hexer containing the BEEFv2-encoded bytes of the transaction built from the current txSpec.
func (spec *txSpec) BEEFv2() Hexer {
	tx := spec.TX()
	beefV2, err := trx.NewBeefFromTransaction(tx)
	spec.requireNoError(err, "beefV2 transaction")

	bytes, err := beefV2.Bytes()
	spec.requireNoError(err, "getting beefV2 bytes")

	return newHexer(bytes)
}

// AtomicBEEF returns a Hexer containing the atomic BEEF-encoded bytes of the transaction built from the current txSpec.
func (spec *txSpec) AtomicBEEF() Hexer {
	tx := spec.TX()
	atomicBEEF, err := tx.AtomicBEEF(false)
	spec.requireNoError(err, "atomic beef transaction")

	return newHexer(atomicBEEF)
}

// RawTX returns the raw hex of the transaction
func (spec *txSpec) RawTX() Hexer {
	tx := spec.TX()
	return newHexer(tx.Bytes())
}

// EF returns the EF hex of the transaction
func (spec *txSpec) EF() Hexer {
	tx := spec.TX()
	ef, err := tx.EF()
	spec.requireNoError(err, "getting ef hex")

	return newHexer(ef)
}

// LockingScripts returns the locking scripts of the transaction
func (spec *txSpec) LockingScripts() []string {
	tx := spec.TX()
	ls := make([]string, len(tx.Outputs))
	for i, o := range tx.Outputs {
		ls[i] = o.LockingScript.String()
	}
	return ls
}

func (spec *txSpec) makeParentTX(satoshis ...uint64) *trx.Transaction {
	tx := trx.NewTransaction()

	total := uint64(0)
	for _, s := range satoshis {
		total += s
	}
	withFee := total + 1
	utxo, err := trx.NewUTXO(spec.getNextGrandparentTXID(), 0, spec.sender.P2PKHLockingScript().String(), withFee)
	spec.requireNoError(err, "creating utxo for parent tx")

	utxo.UnlockingScriptTemplate = spec.sender.P2PKHUnlockingScriptTemplate()

	err = tx.AddInputsFromUTXOs(utxo)
	spec.requireNoError(err, "adding utxo inputs")

	for _, s := range satoshis {
		tx.AddOutput(&trx.TransactionOutput{
			Satoshis:      s,
			LockingScript: spec.sender.P2PKHLockingScript(),
		})
	}
	err = tx.Sign()
	spec.requireNoError(err, "signing parent tx")

	// each merkle proof should have a different block height to not collide with each other
	err = tx.AddMerkleProof(trx.NewMerklePath(spec.getNextBlockHeight(), [][]*trx.PathElement{{
		&trx.PathElement{
			Hash:   tx.TxID(),
			Offset: 0,
		},
	}}))
	spec.requireNoError(err, "adding merkle proof to parent tx")

	return tx
}

func (spec *txSpec) getNextGrandparentTXID() string {
	id := grandparentTXIDs[spec.grandparentTXIndex]
	spec.grandparentTXIndex = (spec.grandparentTXIndex + 1) % len(grandparentTXIDs)
	return id
}

func (spec *txSpec) getNextBlockHeight() uint32 {
	h := spec.blockHeight
	spec.blockHeight++
	return h
}

func (spec *txSpec) requireNoError(err error, msg string) {
	if err != nil {
		panic(fmt.Errorf("TXSpec error: %s: %w", msg, err))
	}
}

type hexer struct {
	bytes []byte
}

func (h *hexer) Bytes() []byte {
	return h.bytes
}

func (h *hexer) Hex() string {
	return fmt.Sprintf("%x", h.bytes)
}

func newHexer(b []byte) Hexer {
	return &hexer{bytes: b}
}
