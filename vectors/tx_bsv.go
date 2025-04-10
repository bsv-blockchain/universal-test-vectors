package vectors

import (
	"github.com/4chain-ag/universal-test-vectors/pkg/testabilities"
	"github.com/4chain-ag/universal-test-vectors/vectors/mapper"
)

var bsvTxs = map[string]testabilities.TransactionSpec{
	"1-in-1-out": testabilities.GivenTX().
		WithInput(1000).
		WithP2PKHOutput(999),
	"2-in-1-out": testabilities.GivenTX().
		WithInput(500).
		WithInput(500).
		WithP2PKHOutput(999),
	"1-in-2-out": testabilities.GivenTX().
		WithInput(1000).
		WithP2PKHOutput(499).
		WithP2PKHOutput(500),
	"3-single-source-inputs": testabilities.GivenTX().
		WithSingleSourceInputs(1000, 500, 500).
		WithP2PKHOutput(1999),
}

func init() {
	addCategory("bsv-tx", bsvTxs, mapper.FromTxSpec)
}
