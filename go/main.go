package main

import (
	"bufio"
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

func main() {
	defer out.Flush()

}
