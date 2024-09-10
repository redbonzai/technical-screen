package main

import (
	"fmt"

	"github.com/redbonzai/sortpkg/sortpkg" // Import the sorting package
)

func main() {
	fmt.Printf("STANDARD TYPE: %v\n", sortpkg.Sort(90, 90, 90, 15))    // STANDARD
	fmt.Printf("SPECIAL TYPE: %v\n", sortpkg.Sort(95, 95, 95, 20))     // SPECIAL
	fmt.Printf("REJECTED TYPE: %v\n", sortpkg.Sort(150, 150, 150, 21)) // REJECTED
}
