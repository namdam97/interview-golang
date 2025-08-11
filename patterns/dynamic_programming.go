package patterns

import (
	"fmt"
	"math"
	"time"
)

// Dynamic Programming Pattern
// Bài toán: Fibonacci sequence

func DemoDynamicProgramming() {
	fmt.Println("\n🧠 6. DYNAMIC PROGRAMMING PATTERN")
	fmt.Println("Bài toán: Fibonacci Sequence")

	n := 40
	fmt.Printf("Input: n = %d\n", n)

	// Cách NAIVE - Recursion O(2^n)
	start := time.Now()
	result1 := fibRecursive(n)
	naive_time := time.Since(start)

	// Cách TỐI ƯU - DP O(n)
	start = time.Now()
	result2 := fibDP(n)
	optimal_time := time.Since(start)

	fmt.Printf("❌ Recursive O(2^n): %d (Time: %v)\n", result1, naive_time)
	fmt.Printf("✅ DP O(n): %d (Time: %v)\n", result2, optimal_time)

	// Demo thêm: House Robber
	houses := []int{2, 7, 9, 3, 1}
	maxMoney := rob(houses)
	fmt.Printf("\n🏠 House Robber max money: %d\n", maxMoney)

	fmt.Println("\n🏢 Ứng dụng thực tế:")
	fmt.Println("- Cache/Memoization systems")
	fmt.Println("- Resource optimization")
	fmt.Println("- Game AI decision making")
	fmt.Println("- Financial modeling")
}

// Naive recursive approach O(2^n)
func fibRecursive(n int) int {
	if n <= 1 {
		return n
	}
	return fibRecursive(n-1) + fibRecursive(n-2)
}

// DP approach O(n)
func fibDP(n int) int {
	if n <= 1 {
		return n
	}

	dp := make([]int, n+1)
	dp[0], dp[1] = 0, 1

	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}

// Ví dụ thực tế: House Robber
func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}

	prev2, prev1 := 0, nums[0]

	for i := 1; i < len(nums); i++ {
		current := int(math.Max(float64(prev1), float64(prev2+nums[i])))
		prev2, prev1 = prev1, current
	}
	return prev1
}
