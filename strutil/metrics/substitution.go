package metrics

// Substitution 表示一个替换函数，该函数用于计算字符替换的得分。
type Substitution interface {
	// Compare 返回字符a[idxA]和b[idxB]的替换得分。
	Compare(a []rune, idxA int, b []rune, idxB int) float64

	// Max 返回字符替换操作的最大分数。
	Max() float64

	// Min 返回字符替换操作的最小分数。
	Min() float64
}
