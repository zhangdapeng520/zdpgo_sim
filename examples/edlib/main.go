package main

import (
	"fmt"
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
	// 查看相似度
	res, err := edlib.StringsSimilarity("string1111", "string2", edlib.Levenshtein)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Similarity: %f", res)
	}
}
