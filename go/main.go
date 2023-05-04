package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

var in = bufio.NewScanner(os.Stdin)
var out = bufio.NewWriter(os.Stdout)

func init() {
	in.Split(bufio.ScanWords)
	in.Buffer([]byte{}, math.MaxInt64)
}

func reads() string {
	in.Scan()
	return in.Text()
}

func readi() int {
	in.Scan()
	ret, _ := strconv.Atoi(in.Text())
	return ret
}

func primes(n int) []int {
	isPrime := make([]bool, n)
	for i := range isPrime {
		isPrime[i] = true
	}
	isPrime[0] = false
	isPrime[1] = false
	for i := range isPrime {
		if !isPrime[i] {
			continue
		}

		for j := 2 * i; j < n; j += i {
			isPrime[j] = false
		}
	}

	primes := make([]int, 0)
	for i := range isPrime {
		if isPrime[i] {
			primes = append(primes, i)
		}
	}
	return primes
}

func modPow(x, e, mod int) int {
	i := 1
	ret := 1

	for i <= e {
		if i&e > 0 {
			ret = (ret * x) % mod
		}
		i <<= 1
		x = (x * x) % mod
		fmt.Fprintln(out, i, x, ret)
	}
	return ret
}

func inv(x, mod int) int {
	return modPow(x, mod-2, mod)
}

func main() {
	defer out.Flush()
}
