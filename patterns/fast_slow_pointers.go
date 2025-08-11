package patterns

import (
	"fmt"
)

// Fast and Slow Pointers Pattern
// Bài toán: Phát hiện cycle trong linked list

type ListNode struct {
	Val  int
	Next *ListNode
}

func DemoFastSlowPointers() {
	fmt.Println("\n🐰🐢 4. FAST & SLOW POINTERS PATTERN")
	fmt.Println("Bài toán: Detect Cycle in Linked List")

	// Tạo linked list với cycle
	head := &ListNode{Val: 3}
	head.Next = &ListNode{Val: 2}
	head.Next.Next = &ListNode{Val: 0}
	head.Next.Next.Next = &ListNode{Val: -4}
	head.Next.Next.Next.Next = head.Next // Tạo cycle

	// Cách NAIVE - Sử dụng HashSet O(n) space
	result1 := hasCycleHashSet(head)

	// Cách TỐI ƯU - Fast & Slow Pointers O(1) space
	result2 := hasCycleFastSlow(head)

	fmt.Printf("❌ HashSet O(n) space: %t\n", result1)
	fmt.Printf("✅ Fast&Slow O(1) space: %t\n", result2)

	// Demo thêm: Tìm middle node
	head2 := createLinkedList([]int{1, 2, 3, 4, 5})
	middle := findMiddle(head2)
	fmt.Printf("\n🎯 Middle node value: %d\n", middle.Val)

	fmt.Println("\n🏢 Ứng dụng thực tế:")
	fmt.Println("- Detect infinite loops in systems")
	fmt.Println("- Find middle element for binary search")
	fmt.Println("- Memory leak detection")
	fmt.Println("- Pagination systems")
}

// Naive approach using HashSet
func hasCycleHashSet(head *ListNode) bool {
	visited := make(map[*ListNode]bool)

	current := head
	for current != nil {
		if visited[current] {
			return true
		}
		visited[current] = true
		current = current.Next
	}
	return false
}

// Optimal approach using Fast & Slow pointers
func hasCycleFastSlow(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}

	slow, fast := head, head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next

		if slow == fast {
			return true
		}
	}
	return false
}

// Ví dụ thực tế: Tìm middle node
func findMiddle(head *ListNode) *ListNode {
	slow, fast := head, head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

func createLinkedList(vals []int) *ListNode {
	if len(vals) == 0 {
		return nil
	}

	head := &ListNode{Val: vals[0]}
	current := head

	for i := 1; i < len(vals); i++ {
		current.Next = &ListNode{Val: vals[i]}
		current = current.Next
	}
	return head
}
