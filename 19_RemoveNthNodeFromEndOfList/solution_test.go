package leetcode

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"lc_p/structures"
	"testing"
)

// 给你一个链表，删除链表的倒数第 n 个结点，并且返回链表的头结点。
//
// 示例 1：
//
// 输入：head = [1,2,3,4,5], n = 2
// 输出：[1,2,3,5]
//
// 示例 2：
//
// 输入：head = [1], n = 1
// 输出：[]
//
// 示例 3：
//
// 输入：head = [1,2], n = 1
// 输出：[1]
//
// 提示：
//
// 链表中结点的数目为 sz
// 1 <= sz <= 30
// 0 <= Node.val <= 100
// 1 <= n <= sz
//
// 进阶：你能尝试使用一趟扫描实现吗？
//
// Related Topics 链表 双指针 👍 3037 👎 0
type ListNode = structures.ListNode

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	// 哨兵节点处理边界
	var dummyNode = &ListNode{}
	dummyNode.Next = head
	// 快慢指针，保证快慢的间隔是n，fast=nil时，slow=倒数n+1节点
	slow, fast := dummyNode, dummyNode
	counter := 0
	for fast != nil {
		if counter > n {
			slow = slow.Next
		}
		fast = fast.Next
		counter++
	}

	// 删除倒数n+1节点的后继节点
	slow.Next = slow.Next.Next
	// 返回头节点
	return dummyNode.Next
}

//leetcode submit region end(Prohibit modification and deletion)

type question19 struct {
	para19
	result19 []int
}

type para19 struct {
	list       []int
	nthFromEnd int
}

var qs = []question19{
	{
		para19{[]int{1, 2, 3, 4, 5}, 2},
		[]int{1, 2, 3, 5},
	},
	//{
	//	para19{[]int{1}, 1},
	//	[]int{},
	//},
	//{
	//	para19{[]int{1, 2}, 1},
	//	[]int{1},
	//},
}

func TestRemoveNthNodeFromEndOfList(t *testing.T) {

	for i, q := range qs {
		testName := fmt.Sprintf("%s%v", "test", i)
		t.Run(testName, func(t *testing.T) {
			fmt.Printf("param: %#v \n", q.para19)

			result := removeNthFromEnd(structures.Ints2List(q.list), q.nthFromEnd)
			formattedResult := structures.List2Ints(result)
			fmt.Printf("isPalindrome: %#v \n", formattedResult)
			assert.Equal(t, q.result19, formattedResult, "The two result should be the same.")
		})
	}
}
