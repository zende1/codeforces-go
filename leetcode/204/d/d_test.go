// Code generated by copypasta/template/leetcode/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/leetcode/testutil"
	"testing"
)

func Test(t *testing.T) {
	t.Log("Current test is [d]")
	examples := [][]string{
		{
			`[2,1,3]`, 
			`1`,
		},
		{
			`[3,4,5,1,2]`, 
			`5`,
		},
		{
			`[1,2,3]`, 
			`0`,
		},
		{
			`[3,1,2,5,4,6]`, 
			`19`,
		},
		{
			`[9,4,2,1,3,6,5,7,8,14,11,10,12,13,16,15,17,18]`, 
			`216212978`,
		},
		// TODO 测试参数的下界和上界
		
	}
	targetCaseNum := 0
	if err := testutil.RunLeetCodeFuncWithExamples(t, numOfWays, examples, targetCaseNum); err != nil {
		t.Fatal(err)
	}
}
// https://leetcode-cn.com/contest/weekly-contest-204/problems/number-of-ways-to-reorder-array-to-get-same-bst/
