package leetcode

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

/**
在股票交易中，如果前一天的股价高于后一天的股价，则可以认为存在一个「交易逆序对」。请设计一个程序，输入一段时间内的股票交易记录 record，返回其中存在的「交
易逆序对」总数。



 示例 1：


输入：record = [9, 7, 5, 4, 6]
输出：8
解释：交易中的逆序对为 (9, 7), (9, 5), (9, 4), (9, 6), (7, 5), (7, 4), (7, 6), (5, 4)。




 提示：

 0 <= record.length <= 50000

 Related Topics 树状数组 线段树 数组 二分查找 分治 有序集合 归并排序 👍 1125 👎 0

*/

// leetcode submit region begin(Prohibit modification and deletion)
func reversePairs(record []int) int {
	return mergeSortCounting(record, 0, len(record)-1, []int{})
}

// mergeSortCounting 归并计数
func mergeSortCounting(record []int, start int, end int, temp []int) int {
	if start >= end {
		return 0
	}

	// tips: 数组越界
	mid := start + (end-start)/2
	leftCount := mergeSortCounting(record, start, mid, temp)
	rightCount := mergeSortCounting(record, mid+1, end, temp)

	// 交叉比较计数
	crossCount := crossCounting(record, start, mid, end, temp)

	return leftCount + rightCount + crossCount
}

// crossCounting 交叉计数
func crossCounting(record []int, start int, mid int, end int, temp []int) int {
	lPtr := start
	rPtr := mid + 1

	// tPtr := 0

	cnt := 0
	for lPtr <= mid && rPtr <= end {
		if record[lPtr] > record[rPtr] {
			cnt += mid - lPtr + 1
			temp = append(temp, record[rPtr])
			rPtr++
		} else {
			temp = append(temp, record[lPtr])
			lPtr++
		}
	}

	// 右侧遍历完，补左侧
	for ; lPtr <= mid; lPtr++ {
		temp = append(temp, record[lPtr])
	}

	// 左侧遍历完，补右侧
	for ; rPtr <= end; rPtr++ {
		temp = append(temp, record[rPtr])
	}

	// 排序
	for i := start; i <= end; i++ {
		record[i] = temp[i-start]
	}

	return cnt
}

// leetcode submit region end(Prohibit modification and deletion)

type qsLCR170 struct {
	input  []int
	output int
}

var qs = []qsLCR170{
	{
		[]int{9, 7, 5, 4, 6},
		8,
	},
}

func TestShuZuZhongDeNiXuDuiLcof(t *testing.T) {

	for i, q := range qs {
		testName := fmt.Sprintf("%s%v", "test", i)
		t.Run(testName, func(t *testing.T) {
			fmt.Printf("param: %#v \n", q.input)

			result := reversePairs(q.input)
			fmt.Printf("reversePairs: %#v \n", result)
			assert.Equal(t, q.output, result, "The two result should be the same.")
		})
	}
}
