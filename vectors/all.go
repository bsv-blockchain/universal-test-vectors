package vectors

// All holds all the testing vectors. (registered by addCategory in one of the `init()` functions)
var All = make(map[string]TestVector)

// TestVector Single test vector with category name.
type TestVector struct {
	Category string
	Vector   any
}

func addCategory[T any, Dst any](category string, specs map[string]T, mapper func(T) Dst) {
	for k, v := range specs {
		All[k] = TestVector{
			Category: category,
			Vector:   mapper(v),
		}
	}
}
