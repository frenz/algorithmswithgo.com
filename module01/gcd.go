package module01

func RemoveIndex(s []int, index int) []int {
	ret := make([]int, 0)
	ret = append(ret, s[:index]...)
	return append(ret, s[index+1:]...)
}

func GCD(a, b int) int {
	tenPrimes := []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31}
	result := 1
	var commonFactors []int
	factorsA := Factor(tenPrimes, a)
	factorsB := Factor(tenPrimes, b)
	for _, itemA := range factorsA {
		for i, itemB := range factorsB {
			if itemA == itemB {
				commonFactors = append(commonFactors, itemA)
				factorsB = RemoveIndex(factorsB, i)
				break
			}
		}
	}
	if len(commonFactors) >= 1 {
		for _, val := range commonFactors {
			result *= val
		}
	}
	return result
}
