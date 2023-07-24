package leetcode

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
)

//编写一个函数来查找字符串数组中的最长公共前缀。
//
// 如果不存在公共前缀，返回空字符串 ""。
//
//
//
// 示例 1：
//
//
//输入：strs = ["flower","flow","flight"]
//输出："fl"
//
//
// 示例 2：
//
//
//输入：strs = ["dog","racecar","car"]
//输出：""
//解释：输入不存在公共前缀。
//
//
//
// 提示：
//
//
// 1 <= strs.length <= 200
// 0 <= strs[i].length <= 200
// strs[i] 仅由小写英文字母组成
//
//
// Related Topics 字典树 字符串 👍 2835 👎 0

// leetcode submit region begin(Prohibit modification and deletion)
func longestCommonPrefix(strs []string) string {
	res := strs[0]
	for i := 1; i < len(strs); i++ {
		if len(res) == 0 {
			break
		}
		cur := strs[i]
		index := 0
		short := min(len(cur), len(res))
		for (index < short) && cur[index] == res[index] {
			index++
		}
		res = res[:index]
	}
	return res
}

func min(x, y int) int {
	if x > y {
		return y
	} else {
		return x
	}
}

// leetcode submit region end(Prohibit modification and deletion)
func tricks(strs []string) string {
	res := strs[0]
	for i := 1; i < len(strs); i++ {
		if len(res) == 0 {
			return ""
		}
		cur := strs[i]
		for strings.Index(cur, res) != 0 {
			if len(cur) == 0 {
				return ""
			}
			res = res[:len(res)-1]
		}
	}
	return res
}

func TestLongestCommonPrefix(t *testing.T) {
	prefix := longestCommonPrefix([]string{"c", "acc", "ccc"})
	fmt.Println(prefix)
	reflect.DeepEqual(prefix, "")
	prefix = longestCommonPrefix([]string{"", "b"})
	fmt.Println(prefix)
	reflect.DeepEqual(prefix, "")
	prefix = longestCommonPrefix([]string{"ab", "a"})
	fmt.Println(prefix)
	reflect.DeepEqual(prefix, "a")
}
