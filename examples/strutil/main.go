package main

/*
@Time : 2022/7/28 15:33
@Author : 张大鹏
@File : main.go
@Software: Goland2021.3.1
@Description:
*/

import (
	"fmt"
	"github.com/zhangdapeng520/zdpgo_sim/strutil"
	"github.com/zhangdapeng520/zdpgo_sim/strutil/metrics"
)

func main() {
	// 计算相似度
	similarity := strutil.Similarity("text", "test", metrics.NewHamming())
	fmt.Printf("%.2f\n", similarity) // Output: 0.75

	// 计算距离
	ham := metrics.NewHamming()
	fmt.Printf("%d\n", ham.Distance("one", "once")) // Output: 2

	// 使用默认参数计算相似度
	similarity = strutil.Similarity("graph", "giraffe", metrics.NewLevenshtein())
	fmt.Printf("%.2f\n", similarity) // Output: 0.43

	// 配置计算相似度的参数
	lev := metrics.NewLevenshtein()
	lev.CaseSensitive = false
	lev.InsertCost = 1
	lev.ReplaceCost = 2
	lev.DeleteCost = 1
	similarity = strutil.Similarity("make", "Cake", lev)
	fmt.Printf("%.2f\n", similarity) // Output: 0.50

	// Jaro
	similarity = strutil.Similarity("think", "tank", metrics.NewJaro())
	fmt.Printf("%.2f\n", similarity) // Output: 0.78

	// Jaro-Winkler
	similarity = strutil.Similarity("think", "tank", metrics.NewJaroWinkler())
	fmt.Printf("%.2f\n", similarity) // Output: 0.80

	// Smith-Waterman-Gotoh
	swg := metrics.NewSmithWatermanGotoh()
	similarity = strutil.Similarity("times roman", "times new roman", swg)
	fmt.Printf("%.2f\n", similarity) // Output: 0.82

	//Customize gap penalty and substitution function.
	swg = metrics.NewSmithWatermanGotoh()
	swg.CaseSensitive = false
	swg.GapPenalty = -0.1
	swg.Substitution = metrics.MatchMismatch{
		Match:    1,
		Mismatch: -0.5,
	}
	similarity = strutil.Similarity("Times Roman", "times new roman", swg)
	fmt.Printf("%.2f\n", similarity) // Output: 0.96

	// Sorensen-Dice
	sd := metrics.NewSorensenDice()
	similarity = strutil.Similarity("time to make haste", "no time to waste", sd)
	fmt.Printf("%.2f\n", similarity) // Output: 0.62

	//Customize n-gram size.
	sd = metrics.NewSorensenDice()
	sd.CaseSensitive = false
	sd.NgramSize = 3
	similarity = strutil.Similarity("Time to make haste", "no time to waste", sd)
	fmt.Printf("%.2f\n", similarity) // Output: 0.53

	// Jaccard
	j := metrics.NewJaccard()
	similarity = strutil.Similarity("time to make haste", "no time to waste", j)
	fmt.Printf("%.2f\n", similarity) // Output: 0.45

	// Customize n-gram size.
	j = metrics.NewJaccard()
	j.CaseSensitive = false
	j.NgramSize = 3
	similarity = strutil.Similarity("Time to make haste", "no time to waste", j)
	fmt.Printf("%.2f\n", similarity) // Output: 0.36

	// Overlap Coefficient
	oc := metrics.NewOverlapCoefficient()
	similarity = strutil.Similarity("time to make haste", "no time to waste", oc)
	fmt.Printf("%.2f\n", similarity) // Output: 0.67

	//Customize n-gram size.
	oc = metrics.NewOverlapCoefficient()
	oc.CaseSensitive = false
	oc.NgramSize = 3
	similarity = strutil.Similarity("Time to make haste", "no time to waste", oc)
	fmt.Printf("%.2f\n", similarity) // Output: 0.57
}
