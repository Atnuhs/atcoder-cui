package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// ### IO === === === === === ===

type Io struct {
	intSc bufio.Scanner
	out   bufio.Writer
	debug bufio.Writer
}

func NewIo() *Io {
	io := Io{
		intSc: *bufio.NewScanner(os.Stdin),
		out:   *bufio.NewWriter(os.Stdout),
		debug: *bufio.NewWriter(os.Stderr),
	}

	io.intSc.Split(bufio.ScanWords)
	return &io
}

func FinalizeIo(io *Io) {
	io.out.Flush()
	io.debug.Flush()
}

func (io *Io) ReadInt() int {
	io.intSc.Scan()
	val, _ := strconv.Atoi(io.intSc.Text())

	return val
}

func (io *Io) ReadInts(N int) []int {

	retSlice := make([]int, N)
	for i := 0; i < N; i++ {
		retSlice[i] = io.ReadInt()
	}
	return retSlice
}

func (io *Io) WriteInt(num int) {
	fmt.Fprintln(&io.out, num)
}

func (io *Io) WriteInts(nums ...int) {
	for _, v := range nums {
		fmt.Fprintf(&io.out, "%d ", v)
	}
	fmt.Fprintf(&io.out, "\n")
}

func (io *Io) DebugInt(num int) {
	fmt.Fprintf(&io.debug, "[Debug]: %d\n", num)
}

func (io *Io) DebugInts(nums ...int) {
	fmt.Fprintf(&io.debug, "[Debug]: %v\n", nums)
}

// ### Num

func Sum(arr []int) int {
	ans := 0
	for _, v := range arr {
		ans += v
	}
	return ans
}

func Max(nums ...int) int {
	if nums == nil {
		return 0
	}

	ret := nums[0]
	for _, v := range nums {
		if ret < v {
			ret = v
		}
	}
	return ret
}

func Min(nums ...int) int {
	if nums == nil {
		return 0
	}

	ret := nums[0]
	for _, v := range nums {
		if ret > v {
			ret = v
		}
	}
	return ret
}

var io = NewIo()

func ScanInt(sc *bufio.Scanner) int {
	sc.Scan()
	val, err := strconv.Atoi(sc.Text())

	if err != nil {
		panic(err)
	}
	return val
}

func main() {
	defer FinalizeIo(io)

}
