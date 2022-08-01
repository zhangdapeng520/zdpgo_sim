package levenshtein

// 原始代码地址：https://github.com/agnivade/levenshtein

import "unicode/utf8"

// minLengthThreshold 是字符串的长度,在这个字符串中将会有一个分配。小于这个的弦是零alloc。
const minLengthThreshold = 32

// ComputeDistance 计算两个字符串之间的levenshtein距离作为一个参数。返回值是levenshtein距离
func ComputeDistance(a, b string) int {
	if len(a) == 0 {
		return utf8.RuneCountInString(b)
	}

	if len(b) == 0 {
		return utf8.RuneCountInString(a)
	}

	if a == b {
		return 0
	}

	// 如果字符串不是ascii编码类型，我们需要转换为[]rune类型
	s1 := []rune(a)
	s2 := []rune(b)

	// 交换值，使得s1长度小的那个字符串
	if len(s1) > len(s2) {
		s1, s2 = s2, s1
	}
	lenS1 := len(s1)
	lenS2 := len(s2)

	// 初始化一行
	var x []uint16
	if lenS1+1 > minLengthThreshold {
		x = make([]uint16, lenS1+1)
	} else {
		// 我们在这里做一个小的优化。
		// 因为一个常数长度的片实际上是一个数组,它不分配。所以我们可以把它重新切片到正确的长度,只要它低于一个期望的阈值。
		x = make([]uint16, minLengthThreshold)
		x = x[:lenS1+1]
	}

	// we start from 1 because index 0 is already 0.
	for i := 1; i < len(x); i++ {
		x[i] = uint16(i)
	}

	// make a dummy bounds check to prevent the 2 bounds check down below.
	// The one inside the loop is particularly costly.
	_ = x[lenS1]

	// 循环比较编辑距离
	for i := 1; i <= lenS2; i++ {
		prev := uint16(i)
		for j := 1; j <= lenS1; j++ {
			current := x[j-1] // match
			if s2[i-1] != s1[j-1] {
				current = min(min(x[j-1]+1, prev+1), x[j]+1)
			}
			x[j-1] = prev
			prev = current
		}
		x[lenS1] = prev
	}

	// 返回右下角那个值，就是最短编辑距离
	return int(x[lenS1])
}

// 求两个uint16类型数字的最小值
func min(a, b uint16) uint16 {
	if a < b {
		return a
	}
	return b
}
