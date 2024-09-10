package sortpkg

import "fmt"

// Sort function classifies the package into STANDARD, SPECIAL, or REJECTED
func Sort(width, height, length, mass int) string {
	if width <= 0 || height <= 0 || length <= 0 {
		return "Invalid input: dimensions must be positive"
	}
	if mass < 0 {
		return "Invalid input: mass must be positive"
	}

	volume := width * height * length
	isBulky := volume > 1000000 || width >= 150 || height >= 150 || length >= 150
	isHeavy := mass >= 20

	if isBulky && isHeavy {
		return "REJECTED"
	} else if isBulky || isHeavy {
		return "SPECIAL"
	} else {
		return "STANDARD"
	}
}

// Main function testing all cases, including edge cases
func main() {
	// Standard packages (neither bulky nor heavy)
	fmt.Println("STANDARD", Sort(90, 90, 90, 10))  // STANDARD
	fmt.Println("STANDARD", Sort(50, 50, 50, 5))   // STANDARD
	fmt.Println("STANDARD", Sort(100, 100, 100, 15)) // STANDARD

	// Special packages (bulky but not heavy)
	fmt.Println("SPECIAL", Sort(200, 100, 100, 10)) // SPECIAL
	fmt.Println("SPECIAL", Sort(150, 90, 90, 10))   // SPECIAL
	fmt.Println("SPECIAL", Sort(90, 150, 90, 15))   // SPECIAL
	fmt.Println("SPECIAL", Sort(90, 90, 150, 15))   // SPECIAL
	fmt.Println("SPECIAL", Sort(100, 100, 1000, 15)) // SPECIAL (Volume = 1,000,000 cm³)

	// Special packages (heavy but not bulky)
	fmt.Println("SPECIAL", Sort(90, 90, 90, 20))  // SPECIAL (Mass is exactly 20 kg)
	fmt.Println("SPECIAL", Sort(95, 95, 95, 25))  // SPECIAL (Mass is greater than 20 kg, but not bulky)

	// Rejected packages (both bulky and heavy)
	fmt.Println("REJECTED", Sort(200, 200, 200, 25)) // REJECTED
	fmt.Println("REJECTED", Sort(150, 150, 150, 21)) // REJECTED (Dimensions exactly 150 cm, mass over 20 kg)

	// Edge cases
	fmt.Println(Sort(0, 90, 90, 10)) // Invalid input: dimensions must be positive
	fmt.Println(Sort(90, 90, 90, -5)) // Invalid input: mass must be positive
	fmt.Println(Sort(90, 0, 90, 10)) // Invalid input: dimensions must be positive
	fmt.Println(Sort(90, 90, 90, 0)) // STANDARD (Mass = 0 is technically not heavy or bulky)
}
