package vectors

import (
	"github.com/4chain-ag/universal-test-vectors/pkg/testabilities"
	"github.com/4chain-ag/universal-test-vectors/vectors/mapper"
)

var users = map[string]testabilities.User{
	"alice":   testabilities.Alice,
	"bob":     testabilities.Bob,
	"charlie": testabilities.Charlie,
}

func init() {
	addCategory("user-", users, mapper.FromUser)
}
