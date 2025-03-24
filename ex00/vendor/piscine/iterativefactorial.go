package piscine

func IterativeFactorial(nb int) int {
	ans := 1
	if (nb < 0) {
		return 0
	}
	if nb > 20 {
		return 0
	}
	for nb > 0 {
		ans = ans * nb
		nb--
	}
	return ans
}