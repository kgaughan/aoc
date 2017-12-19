package day2

func minInt(init int, rest ...int) int {
	min := init
	for _, n := range rest {
		if n < min {
			min = n
		}
	}
	return min
}

// Area gives the total wrapping paper area needed, with some slack based on
// the size of the smallest side.
func Area(l, w, h int) int {
	s1 := l * w
	s2 := w * h
	s3 := h * l
	return minInt(s1, s2, s3) + 2*(s1+s2+s3)
}

func Ribbon(l, w, h int) int {
	s1 := l + w
	s2 := w + h
	s3 := h + l
	perimeter := 2 * minInt(s1, s2, s3)
	// Size of the bow is based on the cubic area of the package.
	cubicArea := l * w * h
	return perimeter + cubicArea
}
