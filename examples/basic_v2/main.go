package main

import (
	"fmt"

	"github.com/zhangdapeng520/zdpgo_clearcode"
	"github.com/zhangdapeng520/zdpgo_pygments"
	"github.com/zhangdapeng520/zdpgo_pygments/lexers"
	"github.com/zhangdapeng520/zdpgo_sim"
)

func main() {
	suffixs := []string{
		".py",
		".php",
		".java",
		".c",
		".cpp",
	}
	for _, suffix := range suffixs {
		fmt.Println("\n正在比较后缀为", suffix, "类型的源码文件\n")

		filePath1 := "examples/test_data/level2_1" + suffix
		filePath2 := "examples/test_data/level2_2" + suffix

		// 代码清洗
		content1, _ := zdpgo_clearcode.ClearCode(filePath1)
		content2, _ := zdpgo_clearcode.ClearCode(filePath2)

		// 词法分析获取token
		lexer1 := lexers.Match(filePath1)
		token1, _ := zdpgo_pygments.GetToken(lexer1, content1)
		lexer2 := lexers.Match(filePath2)
		token2, _ := zdpgo_pygments.GetToken(lexer2, content2)

		//	莱文斯坦-编辑距离(Levenshtein)
		r1 := zdpgo_sim.Compare(token1, token2)
		fmt.Println("莱文斯坦-编辑距离(Levenshtein)", r1)

		// 选择Dice's coefficient
		r1 = zdpgo_sim.Compare(token1, token2, zdpgo_sim.DiceCoefficient())
		fmt.Println("选择Dice's coefficient", r1)

		// 选择jaro
		r1 = zdpgo_sim.Compare(token1, token2, zdpgo_sim.Jaro())
		fmt.Println("选择jaro", r1)

		// 选择JaroWinkler
		r1 = zdpgo_sim.Compare(token1, token2, zdpgo_sim.JaroWinkler())
		fmt.Println("选择JaroWinkler", r1)

		// 选择Hamming
		zdpgo_sim.Compare(token1, token2, zdpgo_sim.Hamming())
		fmt.Println("选择Hamming", r1)

		// 选择Cosine
		r1 = zdpgo_sim.Compare(token1, token2, zdpgo_sim.Cosine())
		fmt.Println("选择Cosine", r1)

		// 选择SimHash
		r1 = zdpgo_sim.Compare(token1, token2, zdpgo_sim.SimHash())
		fmt.Println("选择SimHash", r1)
		fmt.Println("==============")
	}

}
