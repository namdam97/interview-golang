package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("🔥 PERFORMANCE BENCHMARK")
	fmt.Println("=" + string(make([]rune, 50)) + "=")

	// Benchmark Two Pointers vs Brute Force
	benchmarkTwoPointers()

	// Benchmark HashMap vs Brute Force
	benchmarkHashMap()

	// Benchmark Sliding Window
	benchmarkSlidingWindow()

	// Benchmark Binary Search vs Linear Search
	benchmarkBinarySearch()
}

func benchmarkTwoPointers() {
	fmt.Println("\n🎯 Two Pointers Benchmark")

	// Tạo large sorted array
	size := 100000
	nums := make([]int, size)
	for i := 0; i < size; i++ {
		nums[i] = i * 2
	}
	target := nums[size/2] + nums[size/2+1000]

	// Benchmark brute force
	start := time.Now()
	_ = twoSumBruteForce(nums, target)
	bruteTime := time.Since(start)

	// Benchmark two pointers
	start = time.Now()
	_ = twoSumTwoPointers(nums, target)
	optimizedTime := time.Since(start)

	fmt.Printf("Array size: %d\n", size)
	fmt.Printf("Brute Force: %v\n", bruteTime)
	fmt.Printf("Two Pointers: %v\n", optimizedTime)
	fmt.Printf("Speedup: %.2fx\n", float64(bruteTime)/float64(optimizedTime))
}

func benchmarkHashMap() {
	fmt.Println("\n🗂️  HashMap Benchmark")

	size := 50000
	nums := make([]int, size)
	for i := 0; i < size; i++ {
		nums[i] = rand.Intn(size * 2)
	}
	target := nums[100] + nums[200]

	start := time.Now()
	_ = twoSumBruteForceUnsorted(nums, target)
	bruteTime := time.Since(start)

	start = time.Now()
	_ = twoSumHashMap(nums, target)
	hashMapTime := time.Since(start)

	fmt.Printf("Array size: %d\n", size)
	fmt.Printf("Brute Force: %v\n", bruteTime)
	fmt.Printf("HashMap: %v\n", hashMapTime)
	fmt.Printf("Speedup: %.2fx\n", float64(bruteTime)/float64(hashMapTime))
}

func benchmarkSlidingWindow() {
	fmt.Println("\n🪟 Sliding Window Benchmark")

	// Create long string
	s := generateRandomString(10000)

	start := time.Now()
	_ = lengthOfLongestSubstringNaive(s)
	naiveTime := time.Since(start)

	start = time.Now()
	_ = lengthOfLongestSubstringSlidingWindow(s)
	slidingTime := time.Since(start)

	fmt.Printf("String length: %d\n", len(s))
	fmt.Printf("Naive: %v\n", naiveTime)
	fmt.Printf("Sliding Window: %v\n", slidingTime)
	fmt.Printf("Speedup: %.2fx\n", float64(naiveTime)/float64(slidingTime))
}

func benchmarkBinarySearch() {
	fmt.Println("\n🔍 Binary Search Benchmark")

	size := 1000000
	nums := make([]int, size)
	for i := 0; i < size; i++ {
		nums[i] = i * 3
	}
	target := nums[size/2] + 1

	start := time.Now()
	_ = searchInsertLinear(nums, target)
	linearTime := time.Since(start)

	start = time.Now()
	_ = searchInsertBinary(nums, target)
	binaryTime := time.Since(start)

	fmt.Printf("Array size: %d\n", size)
	fmt.Printf("Linear Search: %v\n", linearTime)
	fmt.Printf("Binary Search: %v\n", binaryTime)
	fmt.Printf("Speedup: %.2fx\n", float64(linearTime)/float64(binaryTime))
}

// Helper functions
func twoSumBruteForce(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

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

func twoSumBruteForceUnsorted(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

func twoSumHashMap(nums []int, target int) []int {
	numMap := make(map[int]int)
	for i, num := range nums {
		complement := target - num
		if idx, exists := numMap[complement]; exists {
			return []int{idx, i}
		}
		numMap[num] = i
	}
	return nil
}

func lengthOfLongestSubstringNaive(s string) int {
	maxLen := 0
	for i := 0; i < len(s); i++ {
		for j := i; j < len(s); j++ {
			if hasUniqueChars(s, i, j) {
				if j-i+1 > maxLen {
					maxLen = j - i + 1
				}
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

func lengthOfLongestSubstringSlidingWindow(s string) int {
	charMap := make(map[byte]int)
	left, maxLen := 0, 0

	for right := 0; right < len(s); right++ {
		if lastIndex, exists := charMap[s[right]]; exists && lastIndex >= left {
			left = lastIndex + 1
		}
		charMap[s[right]] = right
		if right-left+1 > maxLen {
			maxLen = right - left + 1
		}
	}
	return maxLen
}

func searchInsertLinear(nums []int, target int) int {
	for i, num := range nums {
		if num >= target {
			return i
		}
	}
	return len(nums)
}

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

func generateRandomString(length int) string {
	chars := "abcdefghijklmnopqrstuvwxyz"
	result := make([]byte, length)
	for i := range result {
		result[i] = chars[rand.Intn(len(chars))]
	}
	return string(result)
}
