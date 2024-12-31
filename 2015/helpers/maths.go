package helpers

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func lcm(ns []int) int {
	lcm := ns[0]
	for i := 1; i < len(ns); i++ {
		lcm = lcm * ns[i] / gcd(lcm, ns[i])
	}
	return lcm
}
