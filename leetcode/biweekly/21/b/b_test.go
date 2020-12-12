// Code generated by copypasta/template/leetcode/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/leetcode/testutil"
	"testing"
)

func Test(t *testing.T) {
	t.Log("Current test is [b]")
	examples := [][]string{
		{
			`"eleetminicoworoep"`, 
			`13`,
		},
		{
			`"leetcodeisgreat"`, 
			`5`,
		},
		{
			`"bcbcbc"`, 
			`6`,
		},
		
	}
	targetCaseNum := 0 // -1
	if err := testutil.RunLeetCodeFuncWithExamples(t, findTheLongestSubstring, examples, targetCaseNum); err != nil {
		t.Fatal(err)
	}
}
// https://leetcode-cn.com/contest/biweekly-contest-21/problems/find-the-longest-substring-containing-vowels-in-even-counts/
