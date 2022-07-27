package main

import (
	"fmt"
	"github.com/zhangdapeng520/zdpgo_sim"
)

/*
@Time : 2022/7/27 16:46
@Author : 张大鹏
@File : main.go
@Software: Goland2021.3.1
@Description:
*/

func main() {
	// 参考项目：https://github.com/hyperjumptech/beda
	sd := zdpgo_sim.NewStringDiff("The First String", "The Second String")

	// wiki地址：https://en.wikipedia.org/wiki/Levenshtein_distance
	lDist := sd.LevenshteinDistance()

	// wiki地址；https://en.wikipedia.org/wiki/N-gram
	tDiff := sd.TrigramCompare()

	// wiki地址：https://en.wikipedia.org/wiki/Jaro%E2%80%93Winkler_distance
	jDiff := sd.JaroDistance()
	jwDiff := sd.JaroWinklerDistance(0.1)

	fmt.Println(tDiff)
	fmt.Println(lDist)
	fmt.Println(jDiff)
	fmt.Println(jwDiff)
}