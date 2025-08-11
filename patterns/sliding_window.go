package patterns

import (
	"fmt"
	"math"
	"time"
)

// Sliding Window Pattern
// Bài toán: Tìm substring dài nhất không có ký tự lặp

func DemoSlidingWindow() {
	fmt.Println("\n🪟 3. SLIDING WINDOW PATTERN")
	fmt.Println("Bài toán: Longest Substring Without Repeating Characters")

	s := "abcabcbb"
	fmt.Printf("Input: \"%s\"\n", s)

	// Cách NAIVE - O(n³)
	start := time.Now()
	result1 := lengthOfLongestSubstringNaive(s)
	naive_time := time.Since(start)

	// Cách TỐI ƯU - Sliding Window O(n)
	start = time.Now()
	result2 := lengthOfLongestSubstringSlidingWindow(s)
	optimal_time := time.Since(start)

	fmt.Printf("❌ Naive O(n³): %d (Time: %v)\n", result1, naive_time)
	fmt.Printf("✅ Sliding Window O(n): %d (Time: %v)\n", result2, optimal_time)

	// Demo thêm: Max sum subarray of size k
	nums := []int{2, 1, 5, 1, 3, 2}
	k := 3
	maxSum := maxSumSubarray(nums, k)
	fmt.Printf("\n📊 Max sum subarray size %d: %d\n", k, maxSum)

	fmt.Println("\n🏢 Ứng dụng thực tế:")
	fmt.Println("- Rate limiting (sliding window counter)")
	fmt.Println("- Moving averages trong trading")
	fmt.Println("- Network packet analysis")
	fmt.Println("- Real-time metrics monitoring")
}

// Naive approach O(n³)
func lengthOfLongestSubstringNaive(s string) int {
	n := len(s)
	maxLen := 0

	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			if hasUniqueChars(s, i, j) {
				maxLen = int(math.Max(float64(maxLen), float64(j-i+1)))
			}
		}
	}
	return maxLen
}

func hasUniqueChars(s string, start, end int) bool {
	chars := make(map[byte]bool)
	for i := start; i <= end; i++ {
		if chars[s[i]] {
			return false
		}
		chars[s[i]] = true
	}
	return true
}

// Sliding Window approach O(n)
func lengthOfLongestSubstringSlidingWindow(s string) int {
	charMap := make(map[byte]int)
	left, maxLen := 0, 0

	for right := 0; right < len(s); right++ {
		if lastIndex, exists := charMap[s[right]]; exists && lastIndex >= left {
			left = lastIndex + 1
		}
		charMap[s[right]] = right
		maxLen = int(math.Max(float64(maxLen), float64(right-left+1)))
	}
	return maxLen
}

// Ví dụ thực tế: Max sum subarray of size k
func maxSumSubarray(nums []int, k int) int {
	if len(nums) < k {
		return 0
	}

	// Tính tổng window đầu tiên
	windowSum := 0
	for i := 0; i < k; i++ {
		windowSum += nums[i]
	}

	maxSum := windowSum

	// Slide window
	for i := k; i < len(nums); i++ {
		windowSum = windowSum - nums[i-k] + nums[i]
		maxSum = int(math.Max(float64(maxSum), float64(windowSum)))
	}

	return maxSum
}
