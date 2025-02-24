package leetcode

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"lc_p/structures"
	"testing"
)

/**
给你一个单链表的头节点 head ，请你判断该链表是否为回文链表。如果是，返回 true ；否则，返回 false 。



 示例 1：


输入：head = [1,2,2,1]
输出：true


 示例 2：


输入：head = [1,2]
输出：false




 提示：


 链表中节点数目在范围[1, 10⁵] 内
 0 <= Node.val <= 9




 进阶：你能否用 O(n) 时间复杂度和 O(1) 空间复杂度解决此题？

 Related Topics 栈 递归 链表 双指针 👍 2011 👎 0

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

// isPalindrome
// 方案2：
// 1. 找到前半部分链表的尾节点。
// 2. 反转后半部分链表。
// 3. 判断是否回文。
// 4. 恢复链表。
// 5. 返回结果。
// O(n) 时间复杂度 & O(1) 空间复杂度
func isPalindrome(head *ListNode) bool {
	if head == nil {
		return false
	}

	// 找到前半部分链表的尾节点（一定得是 first half tail node，而不能是second half head node，否则无法还原）
	firstHalfEnd := endOfFirstHalf(head)
	// 反转后半部分链表 (操作的是first half tail node.next, 不会对first half链表有更新)
	secondHalfStart := reverse(firstHalfEnd.Next)

	// 判断回文
	p1 := head
	p2 := secondHalfStart
	for p2 != nil {
		if p1.Val != p2.Val {
			firstHalfEnd.Next = reverse(secondHalfStart)
			return false
		}
		p1 = p1.Next
		p2 = p2.Next
	}

	// 还原链表并返回结果
	firstHalfEnd.Next = reverse(secondHalfStart)
	return true
}

func endOfFirstHalf(head *ListNode) *ListNode {
	fast, slow := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	return slow
}

func reverse(head *ListNode) *ListNode {
	var prev *ListNode
	p := head
	for p != nil {
		next := p.Next
		p.Next = prev
		prev = p
		p = next
	}
	return prev
}

// leetcode submit region end(Prohibit modification and deletion)

// isPalindrome1
// 方案一：转数组, 双指针比较
// O(n) 时间复杂度 & O(n) 空间复杂度
func isPalindrome1(head *ListNode) bool {
	// 空返回false
	if head == nil {
		return false
	}

	// 链表转数组
	var list []int
	for head != nil {
		list = append(list, head.Val)
		head = head.Next
	}

	// head & tail 双指针
	p1 := 0
	p2 := len(list) - 1
	// head point < tail point
	for p1 < p2 {
		if list[p1] != list[p2] {
			return false
		}
		p1 += 1
		p2 -= 1
	}
	return true
}

type question234 struct {
	input  []int
	result bool
}

var qs = []question234{
	{
		[]int{1, 2, 2, 1}, true,
	},
	// {
	// 	[]int{1, 2}, false,
	// },
	// {
	// 	[]int{}, false,
	// },
	// {
	// 	nil, false,
	// },
}

func TestPalindromeLinkedList(t *testing.T) {

	for i, q := range qs {
		testName := fmt.Sprintf("%s%v", "test", i)
		t.Run(testName, func(t *testing.T) {
			input := structures.Ints2List(q.input)
			fmt.Printf("list1: %v \n", structures.List2Ints(input))

			result := isPalindrome(input)
			fmt.Printf("isPalindrome: %v \n", result)
			assert.Equal(t, result, q.result, "The two result should be the same.")
		})
	}
}
