package main

import (
	"fmt"

	"github.com/zhangdapeng520/zdpgo_sim"
)

func main() {

	suffixs := []string{
		".py",
		// ".php",
		// ".java",
		// ".c",
		// ".cpp",
	}
	for _, suffix := range suffixs {
		fmt.Println("\n正在比较后缀为", suffix, "类型的源码文件\n")

		filePath1 := "examples/test_data/level3_1" + suffix
		filePath2 := "examples/test_data/level3_2" + suffix

		// 2、将代码按换行符拆分，并移除干扰数据。这里的干扰数据包括空字符串，`{`，`};`等这种没有语义的字符串，需要根据不同的语言进行收集和定义。
		// 3、将处理后代码数组进行token化。
		token1Arr, _ := zdpgo_sim.GetFileTokenArr(filePath1)
		token2Arr, _ := zdpgo_sim.GetFileTokenArr(filePath2)

		// 4、按指定的数量级进行比较，比如每次比较1行，每次比较3行等。将指定的行合并为一个token字符串，然后使用相似度算法比较相似度。
		lines := 2 // 每次比较2行
		token1SpreadArr := zdpgo_sim.GetSpreadTokenArr(token1Arr, lines)
		token2SpreadArr := zdpgo_sim.GetSpreadTokenArr(token2Arr, lines)

		// 比较2与1的相似度，以2为主
		for _, token2 := range token2SpreadArr {
			for _, token1 := range token1SpreadArr {
				fmt.Println("正在比较：", token2, token1)

				//	莱文斯坦-编辑距离(Levenshtein)
				r1 := zdpgo_sim.Compare(token1, token2)
				fmt.Println("莱文斯坦-编辑距离(Levenshtein)", r1)

				// 选择Dice's coefficient
				r1 = zdpgo_sim.Compare(token1, token2, zdpgo_sim.DiceCoefficient())
				fmt.Println("选择Dice's coefficient", r1)

				// 选择JaroWinkler
				r1 = zdpgo_sim.Compare(token1, token2, zdpgo_sim.JaroWinkler())
				fmt.Println("选择JaroWinkler", r1)

				// 选择Hamming
				zdpgo_sim.Compare(token1, token2, zdpgo_sim.Hamming())
				fmt.Println("选择Hamming", r1)

				// 选择Cosine
				r1 = zdpgo_sim.Compare(token1, token2, zdpgo_sim.Cosine())
				fmt.Println("选择Cosine", r1)
				fmt.Println("==============")
			}
			fmt.Println("xxxxxxxxxxxxxxxxxxxxx")
			fmt.Println()
			fmt.Println()
		}
	}

}
