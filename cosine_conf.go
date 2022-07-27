package zdpgo_sim

import "github.com/zhangdapeng520/zdpgo_sim/similarity"

// Cosine CosineConf是余弦相似度的配置结构。
func Cosine() OptionFunc {
	return OptionFunc(func(o *option) {
		if o.cmp == nil {
			l := similarity.Cosine{}
			o.base64 = true
			o.cmp = l.CompareUtf8
			if o.ascii {
				o.cmp = l.CompareAscii
			}
		}
	})

}
