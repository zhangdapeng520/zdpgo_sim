package metrics

import (
	"strings"
)

// Hamming 表示度量序列之间相似性的汉明度量。
// 更多信息查看： https://en.wikipedia.org/wiki/Hamming_distance.
type Hamming struct {
	// CaseSensitive specifies if the string comparison is case sensitive.
	CaseSensitive bool
}

// NewHamming 返回一个新的汉明字符串度量值。
// 默认参数：
//   CaseSensitive: true
func NewHamming() *Hamming {
	return &Hamming{
		CaseSensitive: true,
	}
}

// Compare 返回a和b的汉明相似度。返回的相似度是0到1之间的数字。
// 相似度数字越大表示匹配越紧密。
func (m *Hamming) Compare(a, b string) float64 {
	distance, maxLen := m.distance(a, b)
	return 1 - float64(distance)/float64(maxLen)
}

// Distance 返回a和b之间的汉明距离。较低的距离表示较近的匹配。
// 距离为0意味着字符串是相同的。
func (m *Hamming) Distance(a, b string) int {
	distance, _ := m.distance(a, b)
	return distance
}

// 返回距离和最大长度
func (m *Hamming) distance(a, b string) (int, int) {
	// 如果指定了不区分大小写的比较，则使用较低的词汇。
	if !m.CaseSensitive {
		a = strings.ToLower(a)
		b = strings.ToLower(b)
	}
	runesA, runesB := []rune(a), []rune(b)

	// 检查两个项是否都为空。
	lenA, lenB := len(runesA), len(runesB)
	if lenA == 0 && lenB == 0 {
		return 0, 0
	}

	// 如果序列的长度不相等，则将距离初始化为它们的绝对差值。否则，它被设置为0。
	if lenA > lenB {
		lenA, lenB = lenB, lenA
	}
	distance := lenB - lenA

	// 汉明距离计算。
	for i := 0; i < lenA; i++ {
		if runesA[i] != runesB[i] {
			distance++
		}
	}

	return distance, lenB
}
