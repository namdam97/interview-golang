package main

import (
	"fmt"
	"interview-golang/patterns"
)

func main() {
	fmt.Println("🚀 Coding Patterns Demo - Golang Interview Preparation")
	fmt.Println("=" + string(make([]rune, 60)) + "=")

	// Demo các patterns
	patterns.DemoTwoPointers()
	patterns.DemoHashMap()
	patterns.DemoSlidingWindow()
	patterns.DemoFastSlowPointers()
	patterns.DemoBinarySearch()
	patterns.DemoDynamicProgramming()
	patterns.DemoBacktracking()
	patterns.DemoGraphs()
	patterns.DemoRealWorldExamples()

	fmt.Println("\n✅ Demo hoàn thành! Hãy xem code để hiểu rõ từng pattern.")
}
