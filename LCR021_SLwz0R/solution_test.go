package leetcode

import (
	"lc_p/structures"
	"testing"
)

/**
给定一个链表，删除链表的倒数第 n 个结点，并且返回链表的头结点。



 示例 1：




输入：head = [1,2,3,4,5], n = 2
输出：[1,2,3,5]


 示例 2：


输入：head = [1], n = 1
输出：[]


 示例 3：


输入：head = [1,2], n = 1
输出：[1]




 提示：


 链表中结点的数目为 sz
 1 <= sz <= 30
 0 <= Node.val <= 100
 1 <= n <= sz




 进阶：能尝试使用一趟扫描实现吗？




 注意：本题与主站 19 题相同： https://leetcode-cn.com/problems/remove-nth-node-from-end-of-
list/

 Related Topics 链表 双指针 👍 94 👎 0

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
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummyNode := &ListNode{Next: head}

	fast, slow := head, dummyNode
	cnt := 1
	for fast.Next != nil {
		if cnt < n {
			fast = fast.Next
			cnt++
		} else {
			fast = fast.Next
			slow = slow.Next
		}
	}

	slow.Next = slow.Next.Next
	return dummyNode.Next
}

// leetcode submit region end(Prohibit modification and deletion)

func TestSLwz0R(t *testing.T) {

}
