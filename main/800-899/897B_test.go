// Generated by copypasta/template/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/897/B
// https://codeforces.com/problemset/status/897/problem/B
func Test_cf897B(t *testing.T) {
	testCases := [][2]string{
		{
			`2 100`,
			`33`,
		},
		{
			`5 30`,
			`15`,
		},
	}
	testutil.AssertEqualStringCase(t, testCases, 0, cf897B)
}