package main

import (
	"fmt"
	"github.com/zhangdapeng520/zdpgo_sim"
)

/*
@Time : 2022/7/27 16:16
@Author : 张大鹏
@File : main.go
@Software: Goland2021.3.1
@Description:
*/
var (
	s1 = `
def add(a, b):
    return a + b


if __name__ == '__main__':
    print(add(11, 22))
    print(add(111, 222))
    print(add(1111, 2222))
`
	s2 = `
def add(a, b):
    return a + b


if __name__ == '__main__':
    print(add(11, 22))
    print(add(111, 222))
    print(add(1111, 2222))


xxxxxxxxx
`
)

func main() {
	// 查询相似度最高的两个字符串
	r2 := zdpgo_sim.FindBestMatchOne("海刘", []string{"白日依山尽", "黄河入海流", "欲穷千里目", "更上一层楼"})
	fmt.Println(r2)

	//	莱文斯坦-编辑距离(Levenshtein)
	r1 := zdpgo_sim.Compare(s1, s2)
	fmt.Println("莱文斯坦-编辑距离(Levenshtein)", r1)

	// 选择Dice's coefficient
	r1 = zdpgo_sim.Compare(s1, s2, zdpgo_sim.DiceCoefficient())
	fmt.Println("选择Dice's coefficient", r1)

	// 选择jaro
	r1 = zdpgo_sim.Compare(s1, s2, zdpgo_sim.Jaro())
	fmt.Println("选择jaro", r1)

	// 选择JaroWinkler
	r1 = zdpgo_sim.Compare(s1, s2, zdpgo_sim.JaroWinkler())
	fmt.Println("选择JaroWinkler", r1)

	// 选择Hamming
	zdpgo_sim.Compare(s1, s2, zdpgo_sim.Hamming())
	fmt.Println("选择Hamming", r1)

	// 选择Cosine
	r1 = zdpgo_sim.Compare(s1, s2, zdpgo_sim.Cosine())
	fmt.Println("选择Cosine", r1)

	// 选择SimHash
	r1 = zdpgo_sim.Compare(s1, s2, zdpgo_sim.SimHash())
	fmt.Println("选择SimHash", r1)

	var result float64
	fmt.Println(zdpgo_sim.SimilarText(s1, s2, &result), result)
}
