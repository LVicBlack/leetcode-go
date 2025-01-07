package leetcode

import (
	"encoding/json"
	"fmt"
	"lc_p/structures"
	"testing"
)

/**
给你单链表的头节点 head ，请你反转链表，并返回反转后的链表。







 示例 1：


输入：head = [1,2,3,4,5]
输出：[5,4,3,2,1]


 示例 2：


输入：head = [1,2]
输出：[2,1]


 示例 3：


输入：head = []
输出：[]




 提示：


 链表中节点的数目范围是 [0, 5000]
 -5000 <= Node.val <= 5000




 进阶：链表可以选用迭代或递归方式完成反转。你能否用两种方法解决这道题？

 Related Topics 递归 链表 👍 3757 👎 0

*/

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.

 */
/**
 * submit前注释
 */
type ListNode = structures.ListNode

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	cur := head
	for cur != nil {
		next := cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}
	return prev
}

//leetcode submit region end(Prohibit modification and deletion)

func TestReverseLinkedList(t *testing.T) {
	link := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}
	data, _ := json.Marshal(link)
	fmt.Println(string(data))
	reverse, _ := json.Marshal(reverseList(link))
	fmt.Println(string(reverse))
}
