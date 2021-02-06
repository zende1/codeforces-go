package main

import (
	. "fmt"
	"io"
	"math"
	"os"
)

// github.com/EndlessCheng/codeforces-go
func run(in io.Reader, out io.Writer) {
	const m = 1e4
	rf := func() int {
		var f float64
		Fscan(in, &f)
		return int(math.Round(f * m))
	}

	ans := 0
	x := rf()%m + 2e9 // 避免负数
	y := rf()%m + 2e9
	r := rf()
	x0, x1 := ((x-r-1)/m+1)*m, (x+r)/m*m
	for i := x0; i <= x1; i += m {
		d := r*r - (x-i)*(x-i)
		dy := int(math.Sqrt(float64(d))) - 1
		for ; (dy+1)*(dy+1) <= d; dy++ {
		}
		ans += (y+dy)/m - (y-dy-1)/m
	}
	Fprint(out, ans)
}

func main() { run(os.Stdin, os.Stdout) }
