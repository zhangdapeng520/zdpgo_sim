package metrics

import (
	"strings"
	"unicode/utf8"

	"github.com/zhangdapeng520/zdpgo_sim/strutil/internal/mathutil"
)

// Levenshtein 表示度量序列之间相似性的Levenshtein度量。
// 更多信息查看： https://en.wikipedia.org/wiki/Levenshtein_distance.
type Levenshtein struct {
	CaseSensitive bool // 是否忽略大小写
	InsertCost    int  // 表示字符插入的Levenshtein代价。
	DeleteCost    int  // 表示字符删除的Levenshtein代价。
	ReplaceCost   int  // 表示字符替换的Levenshtein代价。
}

// NewLevenshtein 返回一个新的Levenshtein字符串度量值。
// 默认参数:
//   CaseSensitive: true
//   InsertCost: 1
//   DeleteCost: 1
//   ReplaceCost: 1
func NewLevenshtein() *Levenshtein {
	return &Levenshtein{
		CaseSensitive: true,
		InsertCost:    1,
		DeleteCost:    1,
		ReplaceCost:   1,
	}
}

// Compare 返回a和b的Levenshtein相似度。
// 返回的相似度是0到1之间的数字。相似度数字越大表示匹配越紧密。
func (m *Levenshtein) Compare(a, b string) float64 {
	distance, maxLen := m.distance(a, b)
	return 1 - float64(distance)/float64(maxLen)
}

// Distance 返回a和b之间的Levenshtein距离。
// 较低的距离表示较近的匹配。距离为0意味着字符串是相同的。
func (m *Levenshtein) Distance(a, b string) int {
	distance, _ := m.distance(a, b)
	return distance
}

func (m *Levenshtein) distance(a, b string) (int, int) {
	// Check if both terms are empty.
	lenA, lenB := utf8.RuneCountInString(a), utf8.RuneCountInString(b)
	if lenA == 0 && lenB == 0 {
		return 0, 0
	}

	// Check if one of the terms is empty.
	maxLen := mathutil.Max(lenA, lenB)
	if lenA == 0 {
		return m.InsertCost * lenB, maxLen
	}
	if lenB == 0 {
		return m.DeleteCost * lenA, maxLen
	}

	// Lower terms if case insensitive comparison is specified.
	if !m.CaseSensitive {
		a = strings.ToLower(a)
		b = strings.ToLower(b)
	}

	// Initialize cost slice.
	prevCol := make([]int, lenB+1)
	for i := 0; i <= lenB; i++ {
		prevCol[i] = i
	}

	// Calculate distance.
	col := make([]int, lenB+1)
	for i := 0; i < lenA; i++ {
		col[0] = i + 1
		for j := 0; j < lenB; j++ {
			delCost := prevCol[j+1] + m.DeleteCost
			insCost := col[j] + m.InsertCost

			subCost := prevCol[j]
			if a[i] != b[j] {
				subCost += m.ReplaceCost
			}

			col[j+1] = mathutil.Min(delCost, insCost, subCost)
		}

		col, prevCol = prevCol, col
	}

	return prevCol[lenB], maxLen
}
