package leetcode

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

/**
给你一个整数数组 nums，请你将该数组升序排列。

 你必须在 不使用任何内置函数 的情况下解决问题，时间复杂度为 O(nlog(n))，并且空间复杂度尽可能小。






 示例 1：


输入：nums = [5,2,3,1]
输出：[1,2,3,5]


 示例 2：


输入：nums = [5,1,1,2,0,0]
输出：[0,0,1,1,2,5]




 提示：


 1 <= nums.length <= 5 * 10⁴
 -5 * 10⁴ <= nums[i] <= 5 * 10⁴


 Related Topics 数组 分治 桶排序 计数排序 基数排序 排序 堆（优先队列） 归并排序 👍 1079 👎 0

*/

// leetcode submit region begin(Prohibit modification and deletion)
func sortArray(nums []int) []int {
	n := len(nums)
	// 每轮子数组size
	for size := 1; size < n; size *= 2 {
		// 分组合并, left为子数组起始下标, 每次合并两个子数组
		for left := 0; left < n-size; left += size * 2 {
			mid := left + size - 1
			// 避免数组越界
			right := min(left+2*size-1, n-1)
			mergeSort(nums, left, mid, right)
		}
	}
	return nums
}

func mergeSort(arr []int, left, mid, right int) {
	var tmp []int
	l := left
	r := mid + 1

	for l <= mid && r <= right {
		if arr[l] < arr[r] {
			tmp = append(tmp, arr[l])
			l++
		} else {
			tmp = append(tmp, arr[r])
			r++
		}
	}

	for ; l <= mid; l++ {
		tmp = append(tmp, arr[l])
	}
	for ; r <= right; r++ {
		tmp = append(tmp, arr[r])
	}

	k := 0
	for i := left; i <= right; i++ {
		arr[i] = tmp[k]
		k++
	}
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

// leetcode submit region end(Prohibit modification and deletion)

type question912 struct {
	input  []int
	output []int
}

var qs = []question912{
	{
		[]int{5, 2, 3, 1},
		[]int{1, 2, 3, 5},
	},
	{
		[]int{5, 1, 1, 2, 0, 0},
		[]int{0, 0, 1, 1, 2, 5},
	},
}

func TestSortAnArray(t *testing.T) {
	for i, q := range qs {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			fmt.Printf("param: %#v \n", q.input)

			result := sortArray(q.input)
			fmt.Printf("sortArray: %#v \n", result)
			assert.Equal(t, q.output, result, "The two result should be the same.")
		})
	}
}
