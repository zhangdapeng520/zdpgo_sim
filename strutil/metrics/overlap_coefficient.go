package metrics

import (
	"strings"

	"github.com/zhangdapeng520/zdpgo_sim/strutil/internal/mathutil"
	"github.com/zhangdapeng520/zdpgo_sim/strutil/internal/ngram"
)

// OverlapCoefficient 表示序列间相似性度量的重叠系数。这个度规也被称为辛基维茨-辛普森系数。
// 更多信息查看： https://en.wikipedia.org/wiki/Overlap_coefficient.
type OverlapCoefficient struct {
	CaseSensitive bool // 是否忽略大小写
	NgramSize     int  // 表示比较输入序列时生成的标记的大小(以字符为单位)。

}

// NewOverlapCoefficient 返回一个新的重叠系数字符串度量。
// 默认参数:
//   CaseSensitive: true
//   NGramSize: 2
func NewOverlapCoefficient() *OverlapCoefficient {
	return &OverlapCoefficient{
		CaseSensitive: true,
		NgramSize:     2,
	}
}

// Compare 返回a和b的重叠系数相似系数。
// 返回的相似系数是0到1之间的一个数字。相似度数字越大表示匹配越紧密。
// 如果提供的大小小于或等于0，则使用n-grams大小为2的值。
func (m *OverlapCoefficient) Compare(a, b string) float64 {
	// Lower terms if case insensitive comparison is specified.
	if !m.CaseSensitive {
		a = strings.ToLower(a)
		b = strings.ToLower(b)
	}

	// Check if both terms are empty.
	runesA, runesB := []rune(a), []rune(b)
	if len(runesA) == 0 && len(runesB) == 0 {
		return 1
	}

	size := m.NgramSize
	if size <= 0 {
		size = 2
	}

	// Calculate n-gram intersection and minimum subset.
	_, common, totalA, totalB := ngram.Intersection(runesA, runesB, size)

	min := mathutil.Min(totalA, totalB)
	if min == 0 {
		return 0
	}

	// Return similarity.
	return float64(common) / float64(min)
}
