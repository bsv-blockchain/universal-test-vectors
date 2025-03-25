package vectors

// All holds all the testing vectors. (registered by addCategory in one of the `init()` functions)
var All = make(map[string]any)

func addCategory[T any, Dst any](prefix string, specs map[string]T, mapper func(T) Dst) {
	for k, v := range specs {
		All[prefix+k] = mapper(v)
	}
}
