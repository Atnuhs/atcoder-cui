package main

type Operator[T any] func(x1, x2 T) T
type Monoid[T any] struct {
    Op Operator[T]
    E T
}

func MoMax() *Monoid[int] {
    return &Monoid[int]{
        Op: func(x1, x2 int) int {
            return Max(x1, x2)
        },
        E: -INF,
    }
}

func MoMin() *Monoid[int] {
    return &Monoid[int]{
        Op: func(x1, x2 int) int {
            return Min(x1, x2)
        },
        E: INF,
    }
}

func MoSum[T int | float64]() *Monoid[T] {
    return &Monoid[T]{
        Op: func(x1, x2 T) T {
            return x1+x2
        },
        E: 0,
    }
}

func MoXOR() *Monoid[int] {
    return &Monoid[int]{
        Op: func(x1, x2 int) int {
            return x1 ^ x2
        },
        E: 0,
    }
}

func MoMODMul(mod int) *Monoid[int] {
    return &Monoid[int]{
        Op: func(x1, x2 int) int {
            return (x1*x2) % mod
        },
        E: 1,
    }
}


