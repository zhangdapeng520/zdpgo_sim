package metrics

// MatchMismatch 表示一个替换函数，该函数根据比较字符的相等程度返回匹配或不匹配值。
// 匹配值必须大于不匹配值。
type MatchMismatch struct {
	Match    float64 // 表示相等字符替换的得分。
	Mismatch float64 // 表示不相等字符替换的得分。
}

// Compare 如果a[idxA]等于b[idxB]，返回匹配值，否则返回不匹配的值。
func (m MatchMismatch) Compare(a []rune, idxA int, b []rune, idxB int) float64 {
	if a[idxA] == b[idxB] {
		return m.Match
	}

	return m.Mismatch
}

// Max 获取匹配的值
func (m MatchMismatch) Max() float64 {
	return m.Match
}

// Min 获取不匹配的值
func (m MatchMismatch) Min() float64 {
	return m.Mismatch
}
