package metrics

import (
	"strings"

	"github.com/zhangdapeng520/zdpgo_sim/strutil/internal/ngram"
)

// SorensenDice 表示用于测量序列之间相似性的Sorensen-Dice度量。
// 更多信息查看：https://en.wikipedia.org/wiki/Sorensen-Dice_coefficient.
type SorensenDice struct {
	CaseSensitive bool // 是否忽略大小写
	NgramSize     int  // 表示比较输入序列时生成的标记的大小(以字符为单位)。
}

// NewSorensenDice 返回一个新的Sorensen-Dice字符串度量值。
// 默认参数:
//   CaseSensitive: true
//   NGramSize: 2
func NewSorensenDice() *SorensenDice {
	return &SorensenDice{
		CaseSensitive: true,
		NgramSize:     2,
	}
}

// Compare 返回a和b的Sorensen-Dice相似系数。
// 返回的相似度是一个介于0和1之间的数字。相似度数字越大表示匹配越紧密。
// 如果提供的大小小于或等于0，则使用n-grams大小为2的值。
func (m *SorensenDice) Compare(a, b string) float64 {
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

	// Calculate n-gram intersection and union.
	_, common, totalA, totalB := ngram.Intersection(runesA, runesB, size)

	total := totalA + totalB
	if total == 0 {
		return 0
	}

	// Return similarity.
	return 2 * float64(common) / float64(total)
}
