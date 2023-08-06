package main

type Operator func(x1, x2 int) int
type E int
type Monoid struct {
    Op Operator
    E int
}

func MoMax() *Monoid {
    return &Monoid{
        Op: func(x1, x2 int) int {
            return Max(x1, x2)
        },
        E: -INF,
    }
}

func MoMin() *Monoid {
    return &Monoid{
        Op: func(x1, x2 int) int {
            return Min(x1, x2)
        },
        E: INF,
    }
}

func MoSum() *Monoid {
    return &Monoid{
        Op: func(x1, x2 int) int {
            return x1+x2
        },
        E: 0,
    }
}

func MoXOR() *Monoid {
    return &Monoid{
        Op: func(x1, x2 int) int {
            return x1 ^ x2
        },
        E: 0,
    }
}

func MoMODMul(mod int) *Monoid {
    return &Monoid{
        Op: func(x1, x2 int) int {
            return (x1*x2) % mod
        },
        E: 1,
    }
}


