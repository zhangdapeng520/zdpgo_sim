package edlib

// Shingle 接受一个字符串和一个整数作为参数，并返回一个映射。
// 如果字符串为空或k为0，返回一个空映射
func Shingle(s string, k int) map[string]int {
	m := make(map[string]int)
	if s != "" && k != 0 {
		runeS := []rune(s)

		for i := 0; i < len(runeS)-k+1; i++ {
			m[string(runeS[i:i+k])]++
		}
	}
	return m
}

// ShingleSlice 为给定的k找到一个字符串的k-gram接受一个字符串和一个整数作为参数并返回一个切片。
// 如果字符串为空或k为0，返回一个空切片
func ShingleSlice(s string, k int) []string {
	var out []string
	m := make(map[string]int)
	if s != "" && k != 0 {
		runeS := []rune(s)
		for i := 0; i < len(runeS)-k+1; i++ {
			m[string(runeS[i:i+k])]++
		}
		for k := range m {
			out = append(out, k)
		}
	}
	return out
}
