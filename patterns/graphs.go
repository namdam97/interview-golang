package patterns

import (
	"fmt"
)

// Graph Pattern
// Bài toán: Number of Islands

func DemoGraphs() {
	fmt.Println("\n🏝️  8. GRAPH PATTERN")
	fmt.Println("Bài toán: Number of Islands")

	grid := [][]byte{
		{'1', '1', '1', '1', '0'},
		{'1', '1', '0', '1', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '0', '0', '0'},
	}

	fmt.Println("Grid:")
	printGrid(grid)

	// DFS approach
	gridCopy := copyGrid(grid)
	islands := numIslandsDFS(gridCopy)
	fmt.Printf("✅ Number of islands (DFS): %d\n", islands)

	// BFS approach
	gridCopy2 := copyGrid(grid)
	islands2 := numIslandsBFS(gridCopy2)
	fmt.Printf("✅ Number of islands (BFS): %d\n", islands2)

	// Demo thêm: Course Schedule (Topological Sort)
	numCourses := 4
	prerequisites := [][]int{{1, 0}, {2, 0}, {3, 1}, {3, 2}}
	canFinish := canFinishCourses(numCourses, prerequisites)
	fmt.Printf("\n📚 Can finish all courses: %t\n", canFinish)

	fmt.Println("\n🏢 Ứng dụng thực tế:")
	fmt.Println("- Social network analysis")
	fmt.Println("- Route optimization (GPS)")
	fmt.Println("- Dependency resolution")
	fmt.Println("- Network topology discovery")
}

// DFS approach
func numIslandsDFS(grid [][]byte) int {
	if len(grid) == 0 {
		return 0
	}

	islands := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == '1' {
				islands++
				dfsMark(grid, i, j)
			}
		}
	}
	return islands
}

func dfsMark(grid [][]byte, i, j int) {
	if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[0]) || grid[i][j] == '0' {
		return
	}

	grid[i][j] = '0' // Mark as visited

	// Explore 4 directions
	dfsMark(grid, i+1, j)
	dfsMark(grid, i-1, j)
	dfsMark(grid, i, j+1)
	dfsMark(grid, i, j-1)
}

// BFS approach
func numIslandsBFS(grid [][]byte) int {
	if len(grid) == 0 {
		return 0
	}

	islands := 0
	directions := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == '1' {
				islands++
				queue := [][]int{{i, j}}
				grid[i][j] = '0'

				for len(queue) > 0 {
					curr := queue[0]
					queue = queue[1:]

					for _, dir := range directions {
						ni, nj := curr[0]+dir[0], curr[1]+dir[1]
						if ni >= 0 && ni < len(grid) && nj >= 0 && nj < len(grid[0]) && grid[ni][nj] == '1' {
							queue = append(queue, []int{ni, nj})
							grid[ni][nj] = '0'
						}
					}
				}
			}
		}
	}
	return islands
}

// Ví dụ thực tế: Course Schedule (Topological Sort)
func canFinishCourses(numCourses int, prerequisites [][]int) bool {
	// Build adjacency list và indegree array
	graph := make([][]int, numCourses)
	indegree := make([]int, numCourses)

	for _, prereq := range prerequisites {
		course, pre := prereq[0], prereq[1]
		graph[pre] = append(graph[pre], course)
		indegree[course]++
	}

	// BFS với Kahn's algorithm
	queue := []int{}
	for i := 0; i < numCourses; i++ {
		if indegree[i] == 0 {
			queue = append(queue, i)
		}
	}

	completed := 0
	for len(queue) > 0 {
		course := queue[0]
		queue = queue[1:]
		completed++

		for _, nextCourse := range graph[course] {
			indegree[nextCourse]--
			if indegree[nextCourse] == 0 {
				queue = append(queue, nextCourse)
			}
		}
	}

	return completed == numCourses
}

func printGrid(grid [][]byte) {
	for _, row := range grid {
		fmt.Printf("%s\n", string(row))
	}
}

func copyGrid(grid [][]byte) [][]byte {
	copy := make([][]byte, len(grid))
	for i, row := range grid {
		copy[i] = make([]byte, len(row))
		for j, val := range row {
			copy[i][j] = val
		}
	}
	return copy
}
