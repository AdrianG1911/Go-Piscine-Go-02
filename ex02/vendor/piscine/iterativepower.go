package piscine

func IterativePower(nb int, power int) int {
	ans := 1

	if power < 0 || nb == 0 {
		return (0)
	}
	if power == 0 {
		return (1)
	}
	for power > 0 {
		ans = ans * nb
		power--
	}
	return ans
}
