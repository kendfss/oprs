package oprs

type (
	Op[T any]          func(T) T
	BinOp[L, R, T any] func(L, R) T
	TernOp[T any]      func(bool, T, T) T
	Caster[I, O any]   func(I) O
	Var[T any]         func() T
	BinVar[L, R any]   func() (L, R)
	Option[T any]      BinVar[T, error]
)
