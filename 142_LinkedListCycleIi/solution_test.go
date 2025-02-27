package leetcode

import (
	"lc_p/structures"
	"testing"
)

/*
*
给定一个链表的头节点 head ，返回链表开始入环的第一个节点。 如果链表无环，则返回 null。

	如果链表中有某个节点，可以通过连续跟踪 next 指针再次到达，则链表中存在环。 为了表示给定链表中的环，评测系统内部使用整数 pos 来表示链表尾连接到链表

中的位置（索引从 0 开始）。如果 pos 是 -1，则在该链表中没有环。注意：pos 不作为参数进行传递，仅仅是为了标识链表的实际情况。

	不允许修改 链表。






	示例 1：

输入：head = [3,2,0,-4], pos = 1
输出：返回索引为 1 的链表节点
解释：链表中有一个环，其尾部连接到第二个节点。

	示例 2：

输入：head = [1,2], pos = 0
输出：返回索引为 0 的链表节点
解释：链表中有一个环，其尾部连接到第一个节点。

	示例 3：

输入：head = [1], pos = -1
输出：返回 null
解释：链表中没有环。

	提示：


	链表中节点的数目范围在范围 [0, 10⁴] 内
	-10⁵ <= Node.val <= 10⁵
	pos 的值为 -1 或者链表中的一个有效索引




	进阶：你是否可以使用 O(1) 空间解决此题？

	Related Topics 哈希表 链表 双指针 👍 2690 👎 0
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
// detectCycle 快慢指针
//
// 假设进环前长度 = a, 入环距离相遇长度 = b, 环剩余长度 = c,
// fast 是 slow 的两倍，可以得出如下公式
//
// 2(a+b) = a + b + n(b+c) ---> a = (n-1)(b+c) + c
//
// 所以在fast & slow 相遇后, slow 从相遇点出发, head 从头节点出发, 最终会在入环点相遇
//
// tips:
// 1. 为何慢指针第一圈走不完一定会和快指针相遇?
//
// 可以认为快指针和慢指针是相对运动的，假设慢指针的速度是 1节点/秒，快指针的速度是 2节点/秒，当以慢指针为参考系的话（即慢指针静止），快指针的移动速度就是 1节点/秒，所以肯定会相遇。
//
// 2. 为什么在第一圈就会相遇呢？
//
// 设环的长度为 L，当慢指针刚进入环时，慢指针需要走 L 步(即 L 秒)才能走完一圈，此时快指针距离慢指针的最大距离为 L-1，我们再次以慢指针为参考系，如上所说，快指针在按照1节点/秒的速度在追赶慢指针，所以肯定能在 L 秒内追赶到慢指针。
func detectCycle(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	fast, slow := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			p := head
			for p != slow {
				p = p.Next
				slow = slow.Next
			}
			return p
		}
	}

	return nil
}

// leetcode submit region end(Prohibit modification and deletion)

// detectCycle1  哈希表
func detectCycle1(head *ListNode) *ListNode {
	set := make(map[*ListNode]struct{})
	for head != nil {
		if _, ok := set[head]; ok {
			return head
		}
		set[head] = struct{}{}
		head = head.Next
	}

	return nil
}
func TestLinkedListCycleIi(t *testing.T) {

}
