package zdpgo_sim

import (
	"github.com/zhangdapeng520/zdpgo_sim/similarity"
)

// DiceCoefficient ngram 是筛子系数需要用的一个值
func DiceCoefficient(ngram ...int) OptionFunc {
	return OptionFunc(func(o *option) {
		ngram2 := 2
		if len(ngram) > 0 {
			ngram2 = ngram[0]
		}

		// 创建算法对象
		d := &similarity.DiceCoefficient{Ngram: ngram2}

		// 赋值比较方法
		o.cmp = d.CompareUtf8
	})
}
