// Code generated by copypasta/template/leetcode/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/leetcode/testutil"
	"testing"
)

func Test(t *testing.T) {
	t.Log("Current test is [a]")
	examples := [][]string{
		{
			`"aaaabbbbcccc"`, 
			`"abccbaabccba"`,
		},
		{
			`"rat"`, 
			`"art"`,
		},
		{
			`"leetcode"`, 
			`"cdelotee"`,
		},
		{
			`"ggggggg"`, 
			`"ggggggg"`,
		},
		{
			`"spo"`, 
			`"ops"`,
		},
		
	}
	targetCaseNum := 0 // -1
	if err := testutil.RunLeetCodeFuncWithExamples(t, sortString, examples, targetCaseNum); err != nil {
		t.Fatal(err)
	}
}
// https://leetcode-cn.com/contest/biweekly-contest-21/problems/increasing-decreasing-string/
