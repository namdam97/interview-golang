package patterns

import (
	"fmt"
	"math"
	"time"
)

// Binary Search Pattern
// Bài toán: Tìm vị trí insert để maintain sorted order

func DemoBinarySearch() {
	fmt.Println("\n🔍 5. BINARY SEARCH PATTERN")
	fmt.Println("Bài toán: Search Insert Position")

	nums := []int{1, 3, 5, 6}
	target := 5

	fmt.Printf("Input: %v, target: %d\n", nums, target)

	// Cách NAIVE - Linear Search O(n)
	start := time.Now()
	result1 := searchInsertLinear(nums, target)
	naive_time := time.Since(start)

	// Cách TỐI ƯU - Binary Search O(log n)
	start = time.Now()
	result2 := searchInsertBinary(nums, target)
	optimal_time := time.Since(start)

	fmt.Printf("❌ Linear Search O(n): %d (Time: %v)\n", result1, naive_time)
	fmt.Printf("✅ Binary Search O(log n): %d (Time: %v)\n", result2, optimal_time)

	// Demo thêm: Koko eating bananas
	piles := []int{3, 6, 7, 11}
	h := 8
	minSpeed := minEatingSpeed(piles, h)
	fmt.Printf("\n🍌 Koko min eating speed: %d bananas/hour\n", minSpeed)

	fmt.Println("\n🏢 Ứng dụng thực tế:")
	fmt.Println("- Database indexing & search")
	fmt.Println("- Load balancer capacity planning")
	fmt.Println("- Resource allocation optimization")
	fmt.Println("- Version control (git bisect)")
}

// Linear search O(n)
func searchInsertLinear(nums []int, target int) int {
	for i, num := range nums {
		if num >= target {
			return i
		}
	}
	return len(nums)
}

// Binary search O(log n)
func searchInsertBinary(nums []int, target int) int {
	left, right := 0, len(nums)

	for left < right {
		mid := left + (right-left)/2
		if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return left
}

// Ví dụ thực tế: Koko Eating Bananas
func minEatingSpeed(piles []int, h int) int {
	left, right := 1, maxInArray(piles)

	for left < right {
		mid := left + (right-left)/2
		if canFinish(piles, mid, h) {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}

func canFinish(piles []int, speed, h int) bool {
	time := 0
	for _, pile := range piles {
		time += int(math.Ceil(float64(pile) / float64(speed)))
	}
	return time <= h
}

func maxInArray(arr []int) int {
	max := arr[0]
	for _, val := range arr {
		if val > max {
			max = val
		}
	}
	return max
}
