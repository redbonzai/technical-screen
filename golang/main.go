package main

import (
	"fmt"
	"github.com/redbonzai/sortpkg/sortpkg" // Adjust the import path based on your Go module path
)

func main() {
	// Standard packages (neither bulky nor heavy)
	fmt.Println("STANDARD", sortpkg.Sort(90, 90, 90, 10))  // STANDARD
	fmt.Println("STANDARD", sortpkg.Sort(50, 50, 50, 5))   // STANDARD
	fmt.Println("STANDARD", sortpkg.Sort(100, 100, 100, 15)) // STANDARD

	// Special packages (bulky but not heavy)
	fmt.Println("SPECIAL", sortpkg.Sort(200, 100, 100, 10)) // SPECIAL
	fmt.Println("SPECIAL", sortpkg.Sort(150, 90, 90, 10))   // SPECIAL
	fmt.Println("SPECIAL", sortpkg.Sort(90, 150, 90, 15))   // SPECIAL
	fmt.Println("SPECIAL", sortpkg.Sort(90, 90, 150, 15))   // SPECIAL
	fmt.Println("SPECIAL", sortpkg.Sort(100, 100, 1000, 15)) // SPECIAL (Volume = 1,000,000 cm³)

	// Special packages (heavy but not bulky)
	fmt.Println("SPECIAL", sortpkg.Sort(90, 90, 90, 20))  // SPECIAL (Mass is exactly 20 kg)
	fmt.Println("SPECIAL", sortpkg.Sort(95, 95, 95, 25))  // SPECIAL (Mass is greater than 20 kg, but not bulky)

	// Rejected packages (both bulky and heavy)
	fmt.Println("REJECTED", sortpkg.Sort(200, 200, 200, 25)) // REJECTED
	fmt.Println("REJECTED", sortpkg.Sort(150, 150, 150, 21)) // REJECTED (Dimensions exactly 150 cm, mass over 20 kg)

	// Edge cases
	fmt.Println(sortpkg.Sort(0, 90, 90, 10)) // Invalid input: dimensions must be positive
	fmt.Println(sortpkg.Sort(90, 90, 90, -5)) // Invalid input: mass must be positive
	fmt.Println(sortpkg.Sort(90, 0, 90, 10)) // Invalid input: dimensions must be positive
	fmt.Println(sortpkg.Sort(90, 90, 90, 0)) // STANDARD (Mass = 0 is technically not heavy or bulky)
}
