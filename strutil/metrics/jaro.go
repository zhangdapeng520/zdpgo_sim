package metrics

import (
	"strings"
	"unicode/utf8"

	"github.com/zhangdapeng520/zdpgo_sim/strutil/internal/mathutil"
)

// Jaro 表示Jaro度量，用于测量序列之间的相似性。
// 更多信息查看： https://en.wikipedia.org/wiki/Jaro-Winkler_distance.
type Jaro struct {
	// CaseSensitive 指定字符串比较是否区分大小写。
	CaseSensitive bool
}

// NewJaro 返回一个新的Jaro字符串指标。
// 默认参数：
//   CaseSensitive: true
func NewJaro() *Jaro {
	return &Jaro{
		CaseSensitive: true,
	}
}

// Compare 返回a和b的Jaro相似度。
// 返回的相似度是0到1之间的数字。相似度数字越大表示匹配越紧密。
func (m *Jaro) Compare(a, b string) float64 {
	// 检查是否都为空
	lenA, lenB := utf8.RuneCountInString(a), utf8.RuneCountInString(b)
	if lenA == 0 && lenB == 0 {
		return 1
	}

	// 检查其中一个是否为空
	if lenA == 0 || lenB == 0 {
		return 0
	}

	// 如果忽略大小写，则都转换为小写
	if !m.CaseSensitive {
		a = strings.ToLower(a)
		b = strings.ToLower(b)
	}

	// 获取匹配数组
	halfLen := mathutil.Max(0, mathutil.Max(lenA, lenB)/2)
	mrA := matchingRunes(a, b, halfLen)
	mrB := matchingRunes(b, a, halfLen)

	fmLen, smLen := len(mrA), len(mrB)
	if fmLen == 0 || smLen == 0 {
		return 0.0
	}

	// 返回相似度
	return (float64(fmLen)/float64(lenA) +
		float64(smLen)/float64(lenB) +
		float64(fmLen-transpositions(mrA, mrB)/2)/float64(fmLen)) / 3.0
}

// matchingRunes 获取匹配数组
func matchingRunes(a, b string, limit int) []rune {
	common := []rune{}
	runesB := []rune(b)
	lenB := len(runesB)

	for i, r := range a {
		end := mathutil.Min(i+limit+1, lenB)
		for j := mathutil.Max(0, i-limit); j < end; j++ {
			if r == runesB[j] && runesB[j] != -1 {
				common = append(common, runesB[j])
				runesB[j] = -1
				break
			}
		}
	}

	return common
}

func transpositions(a, b []rune) int {
	var count int

	minLen := mathutil.Min(len(a), len(b))
	for i := 0; i < minLen; i++ {
		if a[i] != b[i] {
			count++
		}
	}

	return count
}
