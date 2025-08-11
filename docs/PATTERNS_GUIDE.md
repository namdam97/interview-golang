# 🎯 Coding Patterns Guide - Golang Interview Preparation

## Tổng quan

Đây là bộ sưu tập 8 coding patterns quan trọng nhất cho phỏng vấn kỹ thuật, được implement bằng Golang với:

- ✅ So sánh giải pháp naive vs tối ưu
- ✅ Phân tích độ phức tạp thời gian/không gian
- ✅ Ví dụ thực tế trong production
- ✅ Benchmark performance

## Các Patterns

### 1. 🎯 Two Pointers
**Vấn đề giải quyết:** Tối ưu từ O(n²) xuống O(n) khi làm việc với mảng/chuỗi đã sắp xếp

**Khi nào sử dụng:**
- Tìm cặp số có tổng = target
- Palindrome checking
- Merge sorted arrays
- Container with most water

**Ví dụ thực tế:**
- Tìm cặp sản phẩm có tổng giá = budget
- Load balancer distribution

### 2. 🗂️ HashMap
**Vấn đề giải quyết:** Truy cập O(1) thay vì O(n) tìm kiếm

**Khi nào sử dụng:**
- Two Sum variants
- Group Anagrams
- Frequency counting
- Cache/Memoization

**Ví dụ thực tế:**
- Session management
- Database indexing
- Analytics systems

### 3. 🪟 Sliding Window
**Vấn đề giải quyết:** Tối ưu từ O(n³) xuống O(n) cho subarray/substring problems

**Khi nào sử dụng:**
- Longest substring without repeating chars
- Max sum subarray of size k
- Minimum window substring

**Ví dụ thực tế:**
- Rate limiting
- Moving averages
- Network monitoring

### 4. 🐰🐢 Fast & Slow Pointers
**Vấn đề giải quyết:** Detect cycles với O(1) space thay vì O(n)

**Khi nào sử dụng:**
- Cycle detection
- Finding middle element
- Linked list problems

**Ví dụ thực tế:**
- Memory leak detection
- Infinite loop detection

### 5. 🔍 Binary Search
**Vấn đề giải quyết:** Tìm kiếm O(log n) thay vì O(n)

**Khi nào sử dụng:**
- Search in sorted array
- Find insertion position
- Search in rotated array
- Optimization problems

**Ví dụ thực tế:**
- Database queries
- Resource allocation
- Load testing

### 6. 🧠 Dynamic Programming
**Vấn đề giải quyết:** Tránh tính toán lặp lại, từ O(2^n) xuống O(n)

**Khi nào sử dụng:**
- Fibonacci, climbing stairs
- Longest increasing subsequence
- Coin change, knapsack
- Path counting

**Ví dụ thực tế:**
- Cache systems
- Resource optimization
- Game AI

### 7. 🔄 Backtracking
**Vấn đề giải quyết:** Explore solution space hiệu quả

**Khi nào sử dụng:**
- Generate permutations/combinations
- N-Queens, Sudoku
- Path finding
- Constraint satisfaction

**Ví dụ thực tế:**
- Route planning
- Configuration management
- Game AI

### 8. 🏝️ Graph Algorithms
**Vấn đề giải quyết:** Navigate complex relationships

**Khi nào sử dụng:**
- Connected components
- Shortest path
- Topological sort
- Cycle detection

**Ví dụ thực tế:**
- Social networks
- Dependency resolution
- Network routing

## Performance Comparison

| Pattern | Naive | Optimized | Improvement |
|---------|--------|-----------|-------------|
| Two Pointers | O(n²) | O(n) | n times faster |
| HashMap | O(n²) | O(n) | n times faster |
| Sliding Window | O(n³) | O(n) | n² times faster |
| Binary Search | O(n) | O(log n) | n/log n times faster |
| DP | O(2^n) | O(n) | Exponential improvement |

## Cách sử dụng

```bash
# Chạy demo
go run main.go

# Chạy benchmark
go run cmd/benchmark/main.go

# Hoặc dùng Python script
python3 scripts/run_demo.py
```

## Tips cho phỏng vấn

1. **Nhận diện pattern:** Học cách nhận ra khi nào áp dụng pattern nào
2. **Start simple:** Bắt đầu với giải pháp naive, sau đó tối ưu
3. **Explain tradeoffs:** Luôn giải thích time/space complexity
4. **Practice coding:** Không chỉ hiểu lý thuyết mà phải code thành thạo
5. **Real-world context:** Liên kết với ví dụ thực tế để show understanding

## Kết luận

Các patterns này là "xương sống" của coding interviews. Thay vì học thuộc lời giải, hãy:

- 🎯 Hiểu bản chất của từng pattern
- 🔍 Nhận diện khi nào áp dụng
- ⚡ Linh hoạt kết hợp nhiều patterns
- 🏢 Liên kết với ứng dụng thực tế

Happy coding! 🚀 