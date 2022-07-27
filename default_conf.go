package zdpgo_sim

import (
	"github.com/zhangdapeng520/zdpgo_sim/similarity"
)

func Default() OptionFunc {
	return OptionFunc(func(o *option) {
		if o.cmp == nil {
			l := similarity.EditDistance{}
			o.cmp = l.CompareUtf8
			if o.ascii {
				o.cmp = l.CompareAscii
			}
		}
	})
}
