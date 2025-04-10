package vectors

import (
	"github.com/4chain-ag/universal-test-vectors/pkg/testabilities"
	"github.com/4chain-ag/universal-test-vectors/vectors/mapper"
	"github.com/bsv-blockchain/go-sdk/script"
)

const shortData = "hello world"

var opReturnTxs = map[string]testabilities.TransactionSpec{
	"simple": testabilities.GivenTX().
		WithInput(1000).
		WithP2PKHOutput(999).
		WithOPReturn(shortData),
	"no-opfalse": testabilities.GivenTX().
		WithInput(1000).
		WithP2PKHOutput(999).
		WithOutputScriptParts(testabilities.OpCode(script.OpRETURN), testabilities.PushData(shortData)),
}

func init() {
	addCategory("opreturn", opReturnTxs, mapper.FromTxSpec)
}
