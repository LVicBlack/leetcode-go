package leetcode

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"lc_p/structures"
	"testing"
)

/*
*
给你两个 非空 的链表，表示两个非负的整数。它们每位数字都是按照 逆序 的方式存储的，并且每个节点只能存储 一位 数字。

	请你将两个数相加，并以相同形式返回一个表示和的链表。

	你可以假设除了数字 0 之外，这两个数都不会以 0 开头。



	示例 1：

输入：l1 = [2,4,3], l2 = [5,6,4]
输出：[7,0,8]
解释：342 + 465 = 807.

	示例 2：

输入：l1 = [0], l2 = [0]
输出：[0]

	示例 3：

输入：l1 = [9,9,9,9,9,9,9], l2 = [9,9,9,9]
输出：[8,9,9,9,0,0,0,1]

	提示：


	每个链表中的节点数在范围 [1, 100] 内
	0 <= Node.val <= 9
	题目数据保证列表表示的数字不含前导零


	Related Topics 递归 链表 数学 👍 11086 👎 0
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
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var p *ListNode
	var head *ListNode
	// 进位值
	carry := 0
	for l1 != nil || l2 != nil {
		sum := carry
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		newVal := sum % 10
		carry = sum / 10
		if p == nil {
			head = &ListNode{Val: newVal}
			p = head
		} else {
			p.Next = &ListNode{Val: newVal}
			p = p.Next
		}
	}

	if carry != 0 {
		p.Next = &ListNode{Val: carry}
	}

	return head
}

// leetcode submit region end(Prohibit modification and deletion)
func addTwoNumbers1(l1 *ListNode, l2 *ListNode) *ListNode {
	p := &ListNode{}
	dummyNode := &ListNode{Next: p}
	// 进位值
	carry := 0
	for l1 != nil || l2 != nil {
		sum := carry
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		newVal := sum % 10
		carry = sum / 10

		p.Next = &ListNode{Val: newVal}

		p = p.Next
	}

	if carry != 0 {
		p.Next = &ListNode{Val: carry}
	}

	return dummyNode.Next.Next
}

type question2 struct {
	para21
	result []int
}

type para21 struct {
	l1 []int
	l2 []int
}

var qs = []question2{
	{
		para21{
			[]int{2, 4, 3},
			[]int{5, 6, 4},
		},
		[]int{7, 0, 8},
	},
	{
		para21{
			[]int{0},
			[]int{0},
		},
		[]int{0},
	},
	{
		para21{
			[]int{9, 9, 9, 9, 9, 9, 9},
			[]int{9, 9, 9, 9},
		},
		[]int{8, 9, 9, 9, 0, 0, 0, 1},
	},
}

func TestAddTwoNumbers(t *testing.T) {
	for i, q := range qs {
		testName := fmt.Sprintf("%s%v", "test", i)
		t.Run(testName, func(t *testing.T) {
			l1 := structures.Ints2List(q.l1)
			fmt.Printf("l1: %v \n", structures.List2Ints(l1))
			l2 := structures.Ints2List(q.l2)
			fmt.Printf("l2: %v \n", structures.List2Ints(l2))

			result := addTwoNumbers(l1, l2)
			resultArray := structures.List2Ints(result)
			fmt.Printf("addTwoNumbers result: %v \n", resultArray)
			assert.Equal(t, q.result, resultArray, "The two result should be the same.")
		})
	}
}
