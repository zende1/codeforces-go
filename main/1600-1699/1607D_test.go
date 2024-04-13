// Generated by copypasta/template/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1607/D
// https://codeforces.com/problemset/status/1607/problem/D
func Test_cf1607D(t *testing.T) {
	testCases := [][2]string{
		{
			`8
4
1 2 5 2
BRBR
2
1 1
BB
5
3 1 4 2 5
RBRRB
5
3 1 3 1 3
RBRRB
5
5 1 5 1 5
RBRRB
4
2 2 2 2
BRBR
2
1 -2
BR
4
-2 -1 4 0
RRRR`,
			`YES
NO
YES
YES
NO
YES
YES
YES`,
		},
	}
	testutil.AssertEqualStringCase(t, testCases, 0, cf1607D)
}