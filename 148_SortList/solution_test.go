package leetcode

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"lc_p/structures"
	"strconv"
	"testing"
)

/*
*
给你链表的头结点 head ，请将其按 升序 排列并返回 排序后的链表 。

	示例 1：

输入：head = [4,2,1,3]
输出：[1,2,3,4]

	示例 2：

输入：head = [-1,5,3,4,0]
输出：[-1,0,3,4,5]

	示例 3：

输入：head = []
输出：[]

	提示：


	链表中节点的数目在范围 [0, 5 * 10⁴] 内
	-10⁵ <= Node.val <= 10⁵




	进阶：你可以在 O(n log n) 时间复杂度和常数级空间复杂度下，对链表进行排序吗？

	Related Topics 链表 双指针 分治 排序 归并排序 👍 2454 👎 0
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
func sortList(head *ListNode) *ListNode {

	// 获取链表长度
	length := getListLength(head)

	dummyNode := &ListNode{Next: head}
	// 每组处理的节点数
	for size := 1; size < length; size *= 2 {

		pre := dummyNode      // 记录合并结果
		cur := dummyNode.Next // 当前指针位置
		for cur != nil {
			head1 := cur // 初始化head1
			for i := 1; i < size && cur.Next != nil; i++ {
				cur = cur.Next
			}

			head2 := cur.Next // 初始化head2
			cur.Next = nil    // head1 tail node next 置为 nil
			cur = head2       // cur 移动到head2
			for i := 1; i < size && cur != nil && cur.Next != nil; i++ {
				cur = cur.Next
			}

			var next *ListNode // 记录剩余节点
			if cur != nil {
				next = cur.Next
				cur.Next = nil // head2 tail node next 置为 nil
			}

			// 合并结果首次挂在 dummy node, 之后挂在上一次的合并结果的尾节点后
			pre.Next = mergeSort(head1, head2)

			// pre 指向当前尾节点
			for pre.Next != nil {
				pre = pre.Next
			}

			// 重置指针位置
			cur = next
		}
	}

	return dummyNode.Next
}

// 21. 合并两个有序链表
func mergeSort(head1, head2 *ListNode) *ListNode {
	dummyNode := &ListNode{}
	p := dummyNode

	for head1 != nil && head2 != nil {
		if head1.Val < head2.Val {
			p.Next = head1
			head1 = head1.Next
		} else {
			p.Next = head2
			head2 = head2.Next
		}
		p = p.Next
	}

	if head1 != nil {
		p.Next = head1
	}

	if head2 != nil {
		p.Next = head2
	}

	return dummyNode.Next
}

func getListLength(head *ListNode) int {
	length := 0
	for head != nil {
		length++
		head = head.Next
	}
	return length
}

// leetcode submit region end(Prohibit modification and deletion)

type question148 struct {
	input  []int
	output []int
}

var qs = []question148{
	{
		[]int{5, 2, 3, 1},
		[]int{1, 2, 3, 5},
	},
	{
		[]int{5, 1, 1, 2, 0, 0},
		[]int{0, 0, 1, 1, 2, 5},
	},
}

func TestSortList(t *testing.T) {
	for i, q := range qs {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			fmt.Printf("param: %#v \n", q.input)

			result := sortList(structures.Ints2List(q.input))
			resultArray := structures.List2Ints(result)
			fmt.Printf("sortArray: %#v \n", resultArray)
			assert.Equal(t, q.output, resultArray, "The two result should be the same.")
		})
	}

}
