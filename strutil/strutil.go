package strutil

import (
	"github.com/zhangdapeng520/zdpgo_sim/strutil/internal/ngram"
	"github.com/zhangdapeng520/zdpgo_sim/strutil/internal/stringutil"
)

// StringMetric 表示用于测量字符串之间相似性的度量。度量包实现了以下字符串度量:
//  - Hamming
//  - Jaro
//  - Jaro-Winkler
//  - Levenshtein
//  - Smith-Waterman-Gotoh
//  - Sorensen-Dice
//  - Jaccard
//  - Overlap coefficient
type StringMetric interface {
	Compare(a, b string) float64
}

// Similarity 返回a和b的相似度，使用指定的字符串度量计算。
// 返回的相似性是一个介于0和1之间的数字。
// 相似度数字越大表示匹配越紧密。
func Similarity(a, b string, metric StringMetric) float64 {
	return metric.Compare(a, b)
}

// CommonPrefix 返回指定字符串的公共前缀。如果参数没有共同的前缀，则返回空字符串。
func CommonPrefix(a, b string) string {
	return stringutil.CommonPrefix(a, b)
}

// UniqueSlice 返回一个包含指定字符串片中唯一项的片。
// 输出切片中的项目按照它们在输入切片中出现的顺序排列。
func UniqueSlice(items []string) []string {
	return stringutil.UniqueSlice(items)
}

// SliceContains 如果terms包含q，则返回true，否则返回false。
func SliceContains(terms []string, q string) bool {
	return stringutil.SliceContains(terms, q)
}

// NgramCount 返回提供的术语的指定大小的n元计数。
// 如果提供的大小小于或等于0，则使用n克大小为1的值。
func NgramCount(term string, size int) int {
	return ngram.Count([]rune(term), size)
}

// Ngrams 返回所提供术语的指定大小的所有n元。
// 输出切片中的n-grams是按照它们在输入项中出现的顺序排列的。
// 如果提供的大小小于或等于0，则使用n克大小为1的值。
func Ngrams(term string, size int) []string {
	return ngram.Slice([]rune(term), size)
}

// NgramMap 返回所提供术语的指定大小的所有n-gram的映射，以及它们的频率。
// 该函数还返回n克的总数，即输出映射中所有值的总和。
// 如果提供的大小小于或等于0，则使用n克大小为1的值。
func NgramMap(term string, size int) (map[string]int, int) {
	return ngram.Map([]rune(term), size)
}

// NgramIntersection 返回在这两个术语中找到的指定大小的n-grams的映射，以及它们的频率。
// 该函数还返回公共n克的数量(输出映射中所有值的总和)、第一项中的n克总数和第二项中的n克总数。
// 如果提供的大小小于或等于0，则使用n克大小为1的值。
func NgramIntersection(a, b string, size int) (map[string]int, int, int, int) {
	return ngram.Intersection([]rune(a), []rune(b), size)
}
