package metrics

import (
	"strings"

	"github.com/zhangdapeng520/zdpgo_sim/strutil/internal/mathutil"
)

// SmithWatermanGotoh 表示用于测量序列之间相似性的Smith-Waterman-Gotoh度量。
// 更多信息查看： https://en.wikipedia.org/wiki/Smith-Waterman_algorithm.
type SmithWatermanGotoh struct {
	CaseSensitive bool         // 是否忽略大小写
	GapPenalty    float64      // 定义字符插入或删除的评分惩罚。对于相关的结果，差距惩罚应该是一个非正数。
	Substitution  Substitution // 表示一个替换函数，该函数用于计算字符替换的得分。
}

// NewSmithWatermanGotoh 返回一个新的Smith-Waterman-Gotoh字符串度量。
// 默认参数:
//   CaseSensitive: true
//   GapPenalty: -0.5
//   Substitution: MatchMismatch{
//   	Match:    1,
//   	Mismatch: -2,
//   },
func NewSmithWatermanGotoh() *SmithWatermanGotoh {
	return &SmithWatermanGotoh{
		CaseSensitive: true,
		GapPenalty:    -0.5,
		Substitution: MatchMismatch{
			Match:    1,
			Mismatch: -2,
		},
	}
}

// Compare 返回a和b的Smith-Waterman-Gotoh相似度。
// 返回的相似度是0到1之间的数字。相似度数字越大表示匹配越紧密。
func (m *SmithWatermanGotoh) Compare(a, b string) float64 {
	gap := m.GapPenalty

	// Lower terms if case insensitive comparison is specified.
	if !m.CaseSensitive {
		a = strings.ToLower(a)
		b = strings.ToLower(b)
	}
	runesA, runesB := []rune(a), []rune(b)

	// Check if both terms are empty.
	lenA, lenB := len(runesA), len(runesB)
	if lenA == 0 && lenB == 0 {
		return 1
	}

	// Check if one of the terms is empty.
	if lenA == 0 || lenB == 0 {
		return 0
	}

	// Use default substitution, if none is specified.
	subst := m.Substitution
	if subst == nil {
		subst = MatchMismatch{
			Match:    1,
			Mismatch: -2,
		}
	}

	// Calculate max distance.
	maxDistance := mathutil.Minf(float64(lenA), float64(lenB)) * mathutil.Maxf(subst.Max(), gap)

	// Calculate distance.
	v0 := make([]float64, lenB)
	v1 := make([]float64, lenB)

	distance := mathutil.Maxf(0, gap, subst.Compare(runesA, 0, runesB, 0))
	v0[0] = distance

	for i := 1; i < lenB; i++ {
		v0[i] = mathutil.Maxf(0, v0[i-1]+gap, subst.Compare(runesA, 0, runesB, i))
		distance = mathutil.Maxf(distance, v0[i])
	}

	for i := 1; i < lenA; i++ {
		v1[0] = mathutil.Maxf(0, v0[0]+gap, subst.Compare(runesA, i, runesB, 0))
		distance = mathutil.Maxf(distance, v1[0])

		for j := 1; j < lenB; j++ {
			v1[j] = mathutil.Maxf(0, v0[j]+gap, v1[j-1]+gap, v0[j-1]+subst.Compare(runesA, i, runesB, j))
			distance = mathutil.Maxf(distance, v1[j])
		}

		for j := 0; j < lenB; j++ {
			v0[j] = v1[j]
		}
	}

	// Return similarity.
	return distance / maxDistance
}
