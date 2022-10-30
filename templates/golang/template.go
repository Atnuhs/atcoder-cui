package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

// ### IO === === === === === ===

type Io struct {
	spaceSc bufio.Scanner
	out     bufio.Writer
	debug   bufio.Writer
}

func NewIo() *Io {
	io := Io{
		spaceSc: *bufio.NewScanner(os.Stdin),
		out:     *bufio.NewWriter(os.Stdout),
		debug:   *bufio.NewWriter(os.Stderr),
	}

	io.spaceSc.Split(bufio.ScanWords)
	return &io
}

func FinalizeIo(io *Io) {
	io.out.Flush()
	io.debug.Flush()
}

func (io *Io) ReadStr() string {
	io.spaceSc.Scan()
	return io.spaceSc.Text()
}

func (io *Io) ReadInt() int {
	val, _ := strconv.Atoi(io.ReadStr())
	return val
}

func (io *Io) ReadFloat() float64 {
	val, _ := strconv.ParseFloat(io.ReadStr(), 64)
	return val
}

func (io *Io) ReadInts(N int) []int {

	retSlice := make([]int, N)
	for i := 0; i < N; i++ {
		retSlice[i] = io.ReadInt()
	}
	return retSlice
}

func (io *Io) WriteStr(val string) {
	fmt.Println(&io.out, val)
}

func (io *Io) WriteInt(val int) {
	fmt.Fprintln(&io.out, val)
}

func (io *Io) WriteInts(vals ...int) {
	for _, v := range vals {
		fmt.Fprintf(&io.out, "%d ", v)
	}
	fmt.Fprintf(&io.out, "\n")
}

func (io *Io) DebugInt(val int) {
	fmt.Fprintf(&io.debug, "[Debug]: %d\n", val)
}

func (io *Io) DebugInts(vals ...int) {
	fmt.Fprintf(&io.debug, "[Debug]: %v\n", vals)
}

// ### val

func Sum(arr []int) int {
	ans := 0
	for _, v := range arr {
		ans += v
	}
	return ans
}

func Max(vals ...int) int {
	if vals == nil {
		return 0
	}

	ret := vals[0]
	for _, v := range vals {
		if ret < v {
			ret = v
		}
	}
	return ret
}

func Min(vals ...int) int {
	if vals == nil {
		return 0
	}

	ret := vals[0]
	for _, v := range vals {
		if ret > v {
			ret = v
		}
	}
	return ret
}

var io = NewIo()

type Point struct {
	x float64
	y float64
}

func NewPoint(x float64, y float64) Point {
	return Point{x, y}
}

func ReadPoint() Point {
	x, y := io.ReadFloat(), io.ReadFloat()
	return NewPoint(x, y)
}

func (p Point) minus() Point {
	return NewPoint(-p.x, -p.y)
}

func (p Point) inv() Point {
	return NewPoint(1/p.x, 1/p.y)
}

func (p Point) rotateL90() Point {
	return NewPoint(-p.y, p.x)
}

func (p1 Point) add(p2 Point) Point {
	return NewPoint(p1.x+p2.x, p1.y+p2.y)
}

func (p1 Point) mulA(p2 Point) Point {
	return NewPoint(p1.x*p2.x, p1.y*p2.x)
}
func (p1 Point) mul(val float64) Point {
	return NewPoint(p1.x*val, p1.y*val)
}

func dist(p Point) float64 {
	return math.Sqrt(p.x*p.x + p.y*p.y)
}

func main() {
	defer FinalizeIo(io)
}
