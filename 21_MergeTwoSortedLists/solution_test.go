package leetcode

import (
	"fmt"
	"lc_p/structures"
	"testing"
)

/*
*
将两个升序链表合并为一个新的 升序 链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。

	示例 1：

输入：l1 = [1,2,4], l2 = [1,3,4]
输出：[1,1,2,3,4,4]

	示例 2：

输入：l1 = [], l2 = []
输出：[]

	示例 3：

输入：l1 = [], l2 = [0]
输出：[0]

	提示：


	两个链表的节点数目范围是 [0, 50]
	-100 <= Node.val <= 100
	l1 和 l2 均按 非递减顺序 排列


	Related Topics 递归 链表 👍 3662 👎 0
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

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	var dummyNode = &ListNode{}
	var tailNode = dummyNode
	cur1 := list1
	cur2 := list2
	for cur1 != nil && cur2 != nil {
		if cur1.Val < cur2.Val {
			nextNode := cur1
			//nextNode := &ListNode{cur1.Val, nil}
			tailNode.Next = nextNode
			tailNode = nextNode
			cur1 = cur1.Next
		} else {
			nextNode := cur2
			//nextNode := &ListNode{cur2.Val, nil}
			tailNode.Next = nextNode
			tailNode = nextNode
			cur2 = cur2.Next
		}
	}
	if cur1 != nil {
		tailNode.Next = cur1
	}
	if cur2 != nil {
		tailNode.Next = cur2
	}
	return dummyNode.Next
}

//leetcode submit region end(Prohibit modification and deletion)

func TestMergeTwoSortedLists(t *testing.T) {
	list1 := structures.Ints2List([]int{1, 2, 4})
	//list1 := &ListNode{}
	fmt.Printf("list1: %v \n", structures.List2Ints(list1))
	list2 := structures.Ints2List([]int{1, 3, 4})
	//list2 := &ListNode{}
	fmt.Printf("list2: %v \n", structures.List2Ints(list2))
	//list3 := mergeTwoLists(nil, nil)
	list3 := mergeTwoLists(list1, list2)
	fmt.Printf("list3 merge: %v \n", structures.List2Ints(list3))
	fmt.Printf("list1: %v \n", structures.List2Ints(list1))
	fmt.Printf("list2: %v \n", structures.List2Ints(list2))
}
