package leetcode

import (
	"math"
	"testing"
)

//给你一个整数 x ，如果 x 是一个回文整数，返回 true ；否则，返回 false 。
//
// 回文数是指正序（从左向右）和倒序（从右向左）读都是一样的整数。
//
//
// 例如，121 是回文，而 123 不是。
//
//
//
//
// 示例 1：
//
//
//输入：x = 121
//输出：true
//
//
// 示例 2：
//
//
//输入：x = -121
//输出：false
//解释：从左向右读, 为 -121 。 从右向左读, 为 121- 。因此它不是一个回文数。
//
//
// 示例 3：
//
//
//输入：x = 10
//输出：false
//解释：从右向左读, 为 01 。因此它不是一个回文数。
//
//
//
//
// 提示：
//
//
// -2³¹ <= x <= 2³¹ - 1
//
//
//
//
// 进阶：你能不将整数转为字符串来解决这个问题吗？
//
// Related Topics 数学 👍 2944 👎 0

// leetcode submit region begin(Prohibit modification and deletion)
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	b := x == reserve(x)
	return b
}

func reserve(x int) int {
	var result int64 = 0
	for x != 0 {
		result = result*10 + int64(x%10)
		x = x / 10
	}

	if result > math.MaxInt32 {
		return 0
	}
	return int(result)
}

//leetcode submit region end(Prohibit modification and deletion)

func TestPalindromeNumber(t *testing.T) {
	case1 := 123
	expected1 := false
	palindrome := isPalindrome(case1)
	if palindrome != expected1 {
		t.Errorf("case: %v expected %v, but got %v", case1, expected1, palindrome)
	}
}
