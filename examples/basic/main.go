package main

import (
	"fmt"
	"github.com/zhangdapeng520/zdpgo_clearcode"
	"github.com/zhangdapeng520/zdpgo_sim"
)

/*
@Time : 2022/7/27 16:16
@Author : 张大鹏
@File : main.go
@Software: Goland2021.3.1
@Description:
*/

func main() {
	suffixs := []string{
		".py",
		".php",
		".java",
		".c",
		".cpp",
	}
	for _, suffix := range suffixs {
		filePath1 := "examples/test_data/level1_1" + suffix
		filePath2 := "examples/test_data/level1_2" + suffix
		content1, _ := zdpgo_clearcode.ClearCode(filePath1)
		content2, _ := zdpgo_clearcode.ClearCode(filePath2)
		fmt.Println(filePath1)
		fmt.Println(filePath2)

		//	莱文斯坦-编辑距离(Levenshtein)
		r1 := zdpgo_sim.Compare(string(content1), string(content2))
		fmt.Println("莱文斯坦-编辑距离(Levenshtein)", r1)

		// 选择Dice's coefficient
		r1 = zdpgo_sim.Compare(string(content1), string(content2), zdpgo_sim.DiceCoefficient())
		fmt.Println("选择Dice's coefficient", r1)

		// 选择jaro
		r1 = zdpgo_sim.Compare(string(content1), string(content2), zdpgo_sim.Jaro())
		fmt.Println("选择jaro", r1)

		// 选择JaroWinkler
		r1 = zdpgo_sim.Compare(string(content1), string(content2), zdpgo_sim.JaroWinkler())
		fmt.Println("选择JaroWinkler", r1)

		// 选择Hamming
		zdpgo_sim.Compare(string(content1), string(content2), zdpgo_sim.Hamming())
		fmt.Println("选择Hamming", r1)

		// 选择Cosine
		r1 = zdpgo_sim.Compare(string(content1), string(content2), zdpgo_sim.Cosine())
		fmt.Println("选择Cosine", r1)

		// 选择SimHash
		r1 = zdpgo_sim.Compare(string(content1), string(content2), zdpgo_sim.SimHash())
		fmt.Println("选择SimHash", r1)
		fmt.Println("==============")
	}

}
