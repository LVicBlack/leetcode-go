package leetcode

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

/**
给你一个字符串 s ，一个分割被称为 「好分割」 当它满足：将 s 分割成 2 个字符串 p 和 q ，它们连接起来等于 s 且 p 和 q 中不同字符的数目相
同。

 请你返回 s 中好分割的数目。



 示例 1：

 输入：s = "aacaba"
输出：2
解释：总共有 5 种分割字符串 "aacaba" 的方法，其中 2 种是好分割。
("a", "acaba") 左边字符串和右边字符串分别包含 1 个和 3 个不同的字符。
("aa", "caba") 左边字符串和右边字符串分别包含 1 个和 3 个不同的字符。
("aac", "aba") 左边字符串和右边字符串分别包含 2 个和 2 个不同的字符。这是一个好分割。
("aaca", "ba") 左边字符串和右边字符串分别包含 2 个和 2 个不同的字符。这是一个好分割。
("aacab", "a") 左边字符串和右边字符串分别包含 3 个和 1 个不同的字符。


 示例 2：

 输入：s = "abcd"
输出：1
解释：好分割为将字符串分割成 ("ab", "cd") 。


 示例 3：

 输入：s = "aaaaa"
输出：4
解释：所有分割都是好分割。

 示例 4：

 输入：s = "acbadbaada"
输出：2




 提示：


 s 只包含小写英文字母。
 1 <= s.length <= 10^5


 👍 59 👎 0

*/

// leetcode submit region begin(Prohibit modification and deletion)
func numSplits(s string) int {
	left, right, ans := 0, 0, 0
	chars := []rune(s)
	m1 := make(map[rune]int) // 预先存放每个char的个数
	for _, char := range chars {
		if m1[char] == 0 {
			right++
		}
		m1[char]++
	}

	m2 := make(map[rune]int)
	for _, char := range chars {
		if m2[char] == 0 { // 初次见面
			left++
		}
		m2[char]++
		m1[char]--
		if m1[char] == 0 { // char已消耗完毕
			right--
		}
		if left == right {
			ans++
		}
	}
	return ans
}

// leetcode submit region end(Prohibit modification and deletion)
type question1525 struct {
	in  string
	ans int
}

var qs = []question1525{
	{
		"abab",
		1,
	},
	{
		"aacaba",
		2,
	},
	{
		"aaaaa",
		4,
	},
}

func TestNumberOfGoodWaysToSplitAString(t *testing.T) {
	for i, q := range qs {
		testName := fmt.Sprintf("%s%v", "test", i)
		t.Run(testName, func(t *testing.T) {
			fmt.Printf("param: %#v \n", q.in)

			ans := numSplits(q.in)
			fmt.Printf("numSplits ans: %#v \n", ans)

			assert.Equal(t, q.ans, ans, "The two result should be the same.")
		})
	}
}
