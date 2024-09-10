package sortpkg

// Sort determines whether a package is STANDARD, SPECIAL, or REJECTED
func Sort(width, height, length, mass int) string {
	volume := width * height * length
	isBulky := volume >= 1000000 || width >= 150 || height >= 150 || length >= 150
	isHeavy := mass >= 20

	if isBulky && isHeavy {
		return "REJECTED"
	} else if isBulky || isHeavy {
		return "SPECIAL"
	} else {
		return "STANDARD"
	}
}
