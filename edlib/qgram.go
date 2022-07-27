package edlib

import (
	"math"
)

// QgramDistance 以两个字符串为参数，一个split length定义k-gram shinglength
func QgramDistance(str1, str2 string, splitLength int) int {
	splittedStr1 := Shingle(str1, splitLength)
	splittedStr2 := Shingle(str2, splitLength)

	union := make(map[string]int)
	for i := range splittedStr1 {
		union[i] = 0
	}
	for i := range splittedStr2 {
		union[i] = 0
	}

	res := 0

	for i := range union {
		res += int(math.Abs(float64(splittedStr1[i] - splittedStr2[i])))
	}

	return res
}

// QgramDistanceCustomNgram 计算两个自定义个体集合之间的q-gram相似度以两个n-gram map为参数
func QgramDistanceCustomNgram(splittedStr1, splittedStr2 map[string]int) int {
	union := make(map[string]int)
	for i := range splittedStr1 {
		union[i] = 0
	}
	for i := range splittedStr2 {
		union[i] = 0
	}

	res := 0
	for i := range union {
		res += int(math.Abs(float64(splittedStr1[i] - splittedStr2[i])))
	}

	return res
}

// QgramSimilarity 从Qgram距离计算两个字符串之间的相似度指数(在0和1之间)以两个字符串为参数，分割长度定义k-gram瓦的长度
func QgramSimilarity(str1, str2 string, splitLength int) float32 {
	splittedStr1 := Shingle(str1, splitLength)
	splittedStr2 := Shingle(str2, splitLength)
	res := float32(QgramDistanceCustomNgram(splittedStr1, splittedStr2))
	return 1 - (res / float32(len(splittedStr1)+len(splittedStr2)))
}
