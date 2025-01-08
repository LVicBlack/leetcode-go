package leetcode

import (
	"lc_p/structures"
	"testing"
)

/*
*
给定一个已排序的链表的头

	head ， 删除所有重复的元素，使每个元素只出现一次 。返回 已排序的链表 。



	示例 1：

输入：head = [1,1,2]
输出：[1,2]

	示例 2：

输入：head = [1,1,2,3,3]
输出：[1,2,3]

	提示：


	链表中节点数目在范围 [0, 300] 内
	-100 <= Node.val <= 100
	题目数据保证链表已经按升序 排列


	Related Topics 链表 👍 1177 👎 0
*/
type ListNode = structures.ListNode

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	cur := head
	for cur.Next != nil {
		if cur.Val == cur.Next.Val {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}

	return head
}

//leetcode submit region end(Prohibit modification and deletion)

func deleteDuplicates1(head *ListNode) *ListNode {
	dummyNode := &ListNode{}
	pre := dummyNode
	cur := head
	preValue := -101
	for cur != nil {
		// 跳过重复节点
		if cur.Val == preValue {
			cur = cur.Next
			pre.Next = nil
			continue
		}
		preValue = cur.Val
		pre.Next = cur
		pre = cur
	}

	return dummyNode.Next
}
func TestRemoveDuplicatesFromSortedList(t *testing.T) {

}
