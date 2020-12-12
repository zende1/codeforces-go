// Code generated by copypasta/template/leetcode/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/leetcode/testutil"
	"testing"
)

func Test(t *testing.T) {
	t.Log("Current test is [c]")
	examples := [][]string{
		{
			`[1,null,1,1,1,null,null,1,1,null,1,null,null,null,1,null,1]`, 
			`3`,
		},
		{
			`[1,1,1,null,1,null,null,1,1,null,1]`, 
			`4`,
		},
		{
			`[1]`, 
			`0`,
		},
		
	}
	targetCaseNum := 0 // -1
	if err := testutil.RunLeetCodeFuncWithExamples(t, longestZigZag, examples, targetCaseNum); err != nil {
		t.Fatal(err)
	}
}
// https://leetcode-cn.com/contest/biweekly-contest-21/problems/longest-zigzag-path-in-a-binary-tree/
