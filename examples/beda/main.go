package main

import (
	"fmt"
	"github.com/zhangdapeng520/zdpgo_sim"
)

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

	fmt.Println(lDist)
	fmt.Println(tDiff)
	fmt.Println(jDiff)
	fmt.Println(jwDiff)
}
