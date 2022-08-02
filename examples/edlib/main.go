package main

import (
	"fmt"
	"github.com/zhangdapeng520/zdpgo_clearcode"
	"github.com/zhangdapeng520/zdpgo_sim/edlib"
)

/*
@Time : 2022/7/27 16:59
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
		r1, _ := edlib.StringsSimilarity(content1, content2, edlib.Levenshtein)
		fmt.Println("莱文斯坦-编辑距离(Levenshtein)", r1)

		// 选择SorensenDice
		r1, _ = edlib.StringsSimilarity(content1, content2, edlib.SorensenDice)
		fmt.Println("选择SorensenDice", r1)

		// 选择jaro
		r1, _ = edlib.StringsSimilarity(content1, content2, edlib.Jaro)
		fmt.Println("选择jaro", r1)

		// 选择JaroWinkler
		r1, _ = edlib.StringsSimilarity(content1, content2, edlib.JaroWinkler)
		fmt.Println("选择JaroWinkler", r1)

		// 选择Hamming
		r1, _ = edlib.StringsSimilarity(content1, content2, edlib.Hamming)
		fmt.Println("选择Hamming", r1)

		// 选择Cosine
		r1, _ = edlib.StringsSimilarity(content1, content2, edlib.Cosine)
		fmt.Println("选择Cosine", r1)

		// 选择Lcs
		r1, _ = edlib.StringsSimilarity(content1, content2, edlib.Lcs)
		fmt.Println("选择Lcs", r1)

		// 选择Jaccard
		r1, _ = edlib.StringsSimilarity(content1, content2, edlib.Jaccard)
		fmt.Println("选择Jaccard", r1)

		// 选择Qgram
		r1, _ = edlib.StringsSimilarity(content1, content2, edlib.Qgram)
		fmt.Println("选择Qgram", r1)

		// 选择DamerauLevenshtein
		r1, _ = edlib.StringsSimilarity(content1, content2, edlib.DamerauLevenshtein)
		fmt.Println("选择DamerauLevenshtein", r1)

		// 选择OSADamerauLevenshtein
		r1, _ = edlib.StringsSimilarity(content1, content2, edlib.OSADamerauLevenshtein)
		fmt.Println("选择OSADamerauLevenshtein", r1)

		fmt.Println("==============")
	}

	// 查看相似度
	res, err := edlib.StringsSimilarity("string1111", "string2", edlib.Levenshtein)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Similarity: %f", res)
	}
}
