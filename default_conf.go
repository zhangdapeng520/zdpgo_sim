package zdpgo_sim

import (
	"github.com/zhangdapeng520/zdpgo_sim/similarity"
)

// Default 默认参数
func Default() OptionFunc {
	return OptionFunc(func(o *option) {
		if o.cmp == nil {
			l := similarity.EditDistance{} // 编辑距离
			o.cmp = l.CompareUtf8          // 比较utf8编码
			if o.ascii {
				o.cmp = l.CompareAscii // 比较ascii编码
			}
		}
	})
}
