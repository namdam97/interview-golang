package patterns

import (
	"fmt"
	"time"
)

// Two Pointers Pattern
// Bài toán: Tìm hai số trong mảng đã sắp xếp có tổng bằng target

func DemoTwoPointers() {
	fmt.Println("\n🎯 1. TWO POINTERS PATTERN")
	fmt.Println("Bài toán: Tìm hai số có tổng = target trong mảng đã sắp xếp")

	nums := []int{2, 7, 11, 15, 20, 25, 30}
	target := 22

	fmt.Printf("Input: %v, target: %d\n", nums, target)

	// Cách NAIVE - O(n²)
	start := time.Now()
	result1 := twoSumNaive(nums, target)
	naive_time := time.Since(start)

	// Cách TỐI ƯU - Two Pointers O(n)
	start = time.Now()
	result2 := twoSumTwoPointers(nums, target)
	optimal_time := time.Since(start)

	fmt.Printf("❌ Naive O(n²): %v (Time: %v)\n", result1, naive_time)
	fmt.Printf("✅ Two Pointers O(n): %v (Time: %v)\n", result2, optimal_time)

	// Ví dụ thực tế
	fmt.Println("\n🏢 Ứng dụng thực tế:")
	fmt.Println("- Tìm cặp sản phẩm có tổng giá = budget")
	fmt.Println("- Merge hai sorted arrays")
	fmt.Println("- Palindrome checking")
	fmt.Println("- Container with most water")
}

// Giải pháp naive - Brute force O(n²)
func twoSumNaive(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

// Giải pháp tối ưu - Two Pointers O(n)
func twoSumTwoPointers(nums []int, target int) []int {
	left, right := 0, len(nums)-1

	for left < right {
		sum := nums[left] + nums[right]
		if sum == target {
			return []int{left, right}
		} else if sum < target {
			left++
		} else {
			right--
		}
	}
	return nil
}

// Ví dụ thực tế: Kiểm tra palindrome
func IsPalindrome(s string) bool {
	left, right := 0, len(s)-1

	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}
