package zdpgo_sim

import (
	"github.com/zhangdapeng520/zdpgo_sim/similarity"
)

// Hamming 汉明距离
func Hamming() OptionFunc {
	return OptionFunc(func(o *option) {

		h := &similarity.Hamming{}
		o.cmp = h.CompareUtf8
		if o.ascii {
			o.cmp = h.CompareAscii
		}
	})
}
