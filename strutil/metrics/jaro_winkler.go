package metrics

import (
	"strings"
	"unicode/utf8"

	"github.com/zhangdapeng520/zdpgo_sim/strutil/internal/stringutil"
)

// JaroWinkler 表示Jaro-Winkler度量，用于测量序列之间的相似性。
// 更多信息查看： https://en.wikipedia.org/wiki/Jaro-Winkler_distance.
type JaroWinkler struct {
	CaseSensitive bool // 是否忽略大小写

}

// NewJaroWinkler 获取Jaro-Winkler的对象
// 默认参数：
//   CaseSensitive: true
func NewJaroWinkler() *JaroWinkler {
	return &JaroWinkler{
		CaseSensitive: true,
	}
}

// Compare 返回a和b的Jaro-Winkler相似度。
// 返回的相似度是0到1之间的数字。相似度数字越大表示匹配越紧密。
func (m *JaroWinkler) Compare(a, b string) float64 {
	// 如果忽略大小写，则都转换为小写
	if !m.CaseSensitive {
		a = strings.ToLower(a)
		b = strings.ToLower(b)
	}

	// 计算公共前缀
	lenPrefix := utf8.RuneCountInString(stringutil.CommonPrefix(a, b))
	if lenPrefix > 4 {
		lenPrefix = 4
	}

	jaro := NewJaro()
	jaro.CaseSensitive = m.CaseSensitive

	// 返回相似度
	similarity := jaro.Compare(a, b)
	return similarity + (0.1 * float64(lenPrefix) * (1.0 - similarity))
}
