package metrics

import (
	"strings"

	"github.com/zhangdapeng520/zdpgo_sim/strutil/internal/ngram"
)

// Jaccard 表示Jaccard指数，用于衡量序列之间的相似性。
// 更多信息查看： https://en.wikipedia.org/wiki/Jaccard_index.
type Jaccard struct {
	// CaseSensitive 指定字符串比较是否区分大小写。
	CaseSensitive bool

	// NgramSize 表示比较输入序列时生成的标记的大小(以字符为单位)。
	NgramSize int
}

// NewJaccard 返回一个新的Jaccard字符串指标。
// 默认参数：
//   CaseSensitive: true
//   NGramSize: 2
func NewJaccard() *Jaccard {
	return &Jaccard{
		CaseSensitive: true,
		NgramSize:     2,
	}
}

// Compare 返回a和b的Jaccard相似系数。
// 返回的相似度是一个介于0和1之间的数字。相似度数字越大表示匹配越紧密。
// 如果提供的大小小于或等于0，则使用n-grams大小为2的值。
func (m *Jaccard) Compare(a, b string) float64 {
	// 如果指定了不区分大小写的比较，则使用较低的词汇。
	if !m.CaseSensitive {
		a = strings.ToLower(a)
		b = strings.ToLower(b)
	}

	// 检查两个项是否都为空。
	runesA, runesB := []rune(a), []rune(b)
	if len(runesA) == 0 && len(runesB) == 0 {
		return 1
	}

	size := m.NgramSize
	if size <= 0 {
		size = 2
	}

	// 计算n-grams的交集和并集。
	_, common, totalA, totalB := ngram.Intersection(runesA, runesB, size)

	total := totalA + totalB
	if total == 0 {
		return 0
	}

	// 返回相似度
	return float64(common) / float64(total-common)
}
