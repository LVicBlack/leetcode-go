package leetcode

import (
	"encoding/json"
	"fmt"
	"lc_p/structures"
	"testing"
)

/**
给你单链表的头指针 head 和两个整数 left 和 right ，其中 left <= right 。请你反转从位置 left 到位置 right 的链表节
点，返回 反转后的链表 。



 示例 1：


输入：head = [1,2,3,4,5], left = 2, right = 4
输出：[1,4,3,2,5]


 示例 2：


输入：head = [5], left = 1, right = 1
输出：[5]




 提示：


 链表中节点数目为 n
 1 <= n <= 500
 -500 <= Node.val <= 500
 1 <= left <= right <= n




 进阶： 你可以使用一趟扫描完成反转吗？

 Related Topics 链表 👍 1902 👎 0

*/

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 */

/**
 * submit前注释
 */
type ListNode = structures.ListNode

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	var dummyNode = &ListNode{}
	dummyNode.Next = head
	// 找到reverse的前一个节点
	pre := dummyNode
	for i := 1; i < left; i++ {
		pre = pre.Next
	}

	// 反转的区间里，每遍历到一个节点，让这个新节点来到反转部分的起始位置
	cur := pre.Next
	for i := 0; i < right-left; i++ {
		next := cur.Next
		cur.Next = next.Next
		next.Next = pre.Next
		pre.Next = next
	}

	return dummyNode.Next
}

//leetcode submit region end(Prohibit modification and deletion)

func TestReverseLinkedListIi(t *testing.T) {
	link := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}
	data, _ := json.Marshal(link)
	fmt.Println(string(data))
	reverse, _ := json.Marshal(reverseBetween(link, 2, 4))
	fmt.Println(string(reverse))
}
