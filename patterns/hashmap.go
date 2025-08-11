package patterns

import (
	"fmt"
	"time"
)

// HashMap Pattern
// Bài toán: Two Sum - tìm hai số có tổng bằng target

func DemoHashMap() {
	fmt.Println("\n🗂️  2. HASHMAP PATTERN")
	fmt.Println("Bài toán: Two Sum - tìm hai số có tổng = target")

	nums := []int{2, 7, 11, 15}
	target := 9

	fmt.Printf("Input: %v, target: %d\n", nums, target)

	// Cách NAIVE - O(n²)
	start := time.Now()
	result1 := twoSumBruteForce(nums, target)
	naive_time := time.Since(start)

	// Cách TỐI ƯU - HashMap O(n)
	start = time.Now()
	result2 := twoSumHashMap(nums, target)
	optimal_time := time.Since(start)

	fmt.Printf("❌ Brute Force O(n²): %v (Time: %v)\n", result1, naive_time)
	fmt.Printf("✅ HashMap O(n): %v (Time: %v)\n", result2, optimal_time)

	// Demo thêm: Group Anagrams
	words := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	groups := groupAnagrams(words)
	fmt.Printf("\n🔤 Group Anagrams: %v\n", groups)

	fmt.Println("\n🏢 Ứng dụng thực tế:")
	fmt.Println("- Cache/Memoization trong hệ thống")
	fmt.Println("- Đếm tần suất (analytics)")
	fmt.Println("- Session management")
	fmt.Println("- Database indexing")
}

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

// Ví dụ thực tế: Group Anagrams
func groupAnagrams(strs []string) [][]string {
	groups := make(map[string][]string)

	for _, str := range strs {
		// Tạo key bằng cách sort các ký tự
		key := sortString(str)
		groups[key] = append(groups[key], str)
	}

	result := make([][]string, 0, len(groups))
	for _, group := range groups {
		result = append(result, group)
	}
	return result
}

func sortString(s string) string {
	chars := []rune(s)
	// Simple bubble sort for demo
	for i := 0; i < len(chars); i++ {
		for j := i + 1; j < len(chars); j++ {
			if chars[i] > chars[j] {
				chars[i], chars[j] = chars[j], chars[i]
			}
		}
	}
	return string(chars)
}
