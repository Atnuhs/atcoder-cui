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

func readss(n int) []string {
	ret := make([]string, n)
	for i := range ret {
		ret[i] = reads()
	}
	return ret
}

func readrs(n int) [][]rune {
	ret := make([][]rune, n)
	for i := range ret {
		ret[i] = []rune(reads())
	}
	return ret
}

func readi() int {
	in.Scan()
	ret, _ := strconv.Atoi(in.Text())
	return ret
}

func readis(n int) []int {
	ret := make([]int, n)
	for i := range ret {
		ret[i] = readi()
	}
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

func popBack(a *[]int) int {
	ret := (*a)[len(*a)-1]
	*a = (*a)[:len(*a)-1]
	return ret
}

func lcm(a, b int) int {
	if b == 0 {
		return a
	}
	return lcm(b, a%b)
}

func popFront(a *[]int) int {
	ret := (*a)[0]
	*a = (*a)[1:]
	return ret
}

func nextPerm(a []int) bool {
	// search i
	i := len(a) - 2
	for i >= 0 && a[i] >= a[i+1] {
		i--
	}
	if i < 0 {
		return false
	}
	j := len(a) - 1
	for j >= 0 && a[j] <= a[i] {
		j--
	}

	a[i], a[j] = a[j], a[i]

	l := i + 1
	r := len(a) - 1
	for l < r {
		a[l], a[r] = a[r], a[l]
		l++
		r--
	}
	return true
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func abs(x int) int {
	if x < 0 {
		x = -x
	}
	return x
}

func main() {
	defer out.Flush()
}
