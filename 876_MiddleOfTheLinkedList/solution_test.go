package leetcode

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"lc_p/structures"
	"testing"
)

/*
*
给你单链表的头结点 head ，请你找出并返回链表的中间结点。

	如果有两个中间结点，则返回第二个中间结点。



	示例 1：

输入：head = [1,2,3,4,5]
输出：[3,4,5]
解释：链表只有一个中间结点，值为 3 。

	示例 2：

输入：head = [1,2,3,4,5,6]
输出：[4,5,6]
解释：该链表有两个中间结点，值分别为 3 和 4 ，返回第二个结点。

	提示：


	链表的结点数范围是 [1, 100]
	1 <= Node.val <= 100


	Related Topics 链表 双指针 👍 1047 👎 0
*/
type ListNode = structures.ListNode

// leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func middleNode(head *ListNode) *ListNode {

	slow, fast := head, head

	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}

	return slow
}

// leetcode submit region end(Prohibit modification and deletion)
type question876 struct {
	input  []int
	output []int
}

var qs = []question876{
	{
		[]int{1, 2, 3, 4, 5},
		[]int{3, 4, 5},
	},
	{
		[]int{1, 2, 3, 4, 5, 6},
		[]int{4, 5, 6},
	},
}

func TestMiddleOfTheLinkedList(t *testing.T) {
	for i, q := range qs {
		testName := fmt.Sprintf("%s%v", "test", i)
		t.Run(testName, func(t *testing.T) {
			input := structures.Ints2List(q.input)
			fmt.Printf("input: %v \n", q.input)

			output := middleNode(input)
			outputArray := structures.List2Ints(output)
			fmt.Printf("middleNode output: %v \n", outputArray)
			assert.Equal(t, q.output, outputArray, "The two output should be the same.")
		})
	}
}
