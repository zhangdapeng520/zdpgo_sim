package edlib

// SorensenDiceCoefficient 以两个字符串为参数，split length定义k-gram瓦的长度
func SorensenDiceCoefficient(str1, str2 string, splitLength int) float32 {
	if str1 == "" && str2 == "" {
		return 0
	}
	shingle1 := Shingle(str1, splitLength)
	shingle2 := Shingle(str2, splitLength)

	intersection := float32(0)
	for i := range shingle1 {
		if _, ok := shingle2[i]; ok {
			intersection++
		}
	}
	return 2.0 * intersection / float32(len(shingle1)+len(shingle2))
}
