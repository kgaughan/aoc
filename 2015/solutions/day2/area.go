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

type dimensions struct{ l, w, h int }

// Area gives the total wrapping paper area needed, with some slack based on
// the size of the smallest side.
func (d dimensions) area() int {
	s1 := d.l * d.w
	s2 := d.w * d.h
	s3 := d.h * d.l
	return minInt(s1, s2, s3) + 2*(s1+s2+s3)
}

func (d dimensions) ribbon() int {
	s1 := d.l + d.w
	s2 := d.w + d.h
	s3 := d.h + d.l
	perimeter := 2 * minInt(s1, s2, s3)
	// Size of the bow is based on the cubic area of the package.
	cubicArea := d.l * d.w * d.h
	return perimeter + cubicArea
}
