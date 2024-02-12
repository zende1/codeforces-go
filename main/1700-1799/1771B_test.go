// Code generated by copypasta/template/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1771/B
// https://codeforces.com/problemset/status/1771/problem/B
func Test_cf1771B(t *testing.T) {
	testCases := [][2]string{
		{
			`2
3 2
1 3
2 3
4 2
1 2
2 3`,
			`4
5`,
		},
	}
	testutil.AssertEqualStringCase(t, testCases, 0, cf1771B)
}