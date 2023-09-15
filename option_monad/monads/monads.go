package monads

// Wrapper Type => The contract type between the monad ecosystem and the external ecosystem

type Option[T any] struct {
	Result  T
	IsValid bool
}

// Wrapper Function =>The entry to the monad ecosystem
func WrapperOption[T any](value T) Option[T] {
	return Option[T]{Result: value, IsValid: true}
}

// Run Function (Transformer) => Runs transformations into monadic values
func RunOption[T any, R any](value Option[T], transform func(T) Option[R]) Option[R] {
	if value.IsValid {
		return transform(value.Result)
	}
	var result R
	return Option[R]{Result: result, IsValid: false}
}
