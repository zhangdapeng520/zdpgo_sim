package edlib

import (
	"errors"

	"github.com/zhangdapeng520/zdpgo_sim/edlib/internal/orderedmap"
)

// Algorithm 使用Integer类型来标识编辑距离算法
type Algorithm uint8

// Algorithm 算法实现
const (
	Levenshtein Algorithm = iota
	DamerauLevenshtein
	OSADamerauLevenshtein
	Lcs
	Hamming
	Jaro
	JaroWinkler
	Cosine
	Jaccard
	SorensenDice
	Qgram
)

// StringsSimilarity 返回一个相似度索引[0..]1]根据参数中给定的编辑距离算法。
// 使用指定的 Algorithm 算法类型
// 通过这个函数，使用了cos和Jaccard算法，以及长度为2的Shingle split方法。
func StringsSimilarity(str1 string, str2 string, algo Algorithm) (float32, error) {
	switch algo {
	case Levenshtein:
		return matchingIndex(str1, str2, LevenshteinDistance(str1, str2)), nil
	case DamerauLevenshtein:
		return matchingIndex(str1, str2, DamerauLevenshteinDistance(str1, str2)), nil
	case OSADamerauLevenshtein:
		return matchingIndex(str1, str2, OSADamerauLevenshteinDistance(str1, str2)), nil
	case Lcs:
		return matchingIndex(str1, str2, LCSEditDistance(str1, str2)), nil
	case Hamming:
		distance, err := HammingDistance(str1, str2)
		if err == nil {
			return matchingIndex(str1, str2, distance), nil
		}
		return 0.0, err
	case Jaro:
		return JaroSimilarity(str1, str2), nil
	case JaroWinkler:
		return JaroWinklerSimilarity(str1, str2), nil
	case Cosine:
		return CosineSimilarity(str1, str2, 2), nil
	case Jaccard:
		return JaccardSimilarity(str1, str2, 2), nil
	case SorensenDice:
		return SorensenDiceCoefficient(str1, str2, 2), nil
	case Qgram:
		return QgramSimilarity(str1, str2, 2), nil
	default:
		return 0.0, errors.New("Illegal argument for algorithm method")
	}
}

// 返回匹配索引E [0..]1]从两个字符串和一个编辑距离
func matchingIndex(str1 string, str2 string, distance int) float32 {
	// 将字符串转换为符文片
	runeStr1 := []rune(str1)
	runeStr2 := []rune(str2)

	// 比较符文数组的长度，并使它们之间的百分比匹配
	if len(runeStr1) >= len(runeStr2) {
		return float32(len(runeStr1)-distance) / float32(len(runeStr1))
	}
	return float32(len(runeStr2)-distance) / float32(len(runeStr2))
}

// FuzzySearch 实现对字符串列表的近似搜索，并返回与输入字符串最接近的一个
func FuzzySearch(str string, strList []string, algo Algorithm) (string, error) {
	var higherMatchPercent float32
	var tmpStr string
	for _, strToCmp := range strList {
		sim, err := StringsSimilarity(str, strToCmp, algo)
		if err != nil {
			return "", err
		}

		if sim == 1.0 {
			return strToCmp, nil
		} else if sim > higherMatchPercent {
			higherMatchPercent = sim
			tmpStr = strToCmp
		}
	}

	return tmpStr, nil
}

// FuzzySearchThreshold 实现对字符串列表的近似搜索，并返回与输入字符串最接近的一个。在参数上取相似阈值。
func FuzzySearchThreshold(str string, strList []string, minSim float32, algo Algorithm) (string, error) {
	var higherMatchPercent float32
	var tmpStr string
	for _, strToCmp := range strList {
		sim, err := StringsSimilarity(str, strToCmp, algo)
		if err != nil {
			return "", err
		}

		if sim == 1.0 {
			return strToCmp, nil
		} else if sim > higherMatchPercent && sim >= minSim {
			higherMatchPercent = sim
			tmpStr = strToCmp
		}
	}
	return tmpStr, nil
}

// FuzzySearchSet 实现对字符串列表的近似搜索，返回一个由x个字符串组成的集合，与输入的字符串按与基本字符串的相似度排序。
// 接受a quantity参数来定义所需的输出字符串数量(例如谷歌键盘词建议的情况下为3)。
func FuzzySearchSet(str string, strList []string, quantity int, algo Algorithm) ([]string, error) {
	sortedMap := make(orderedmap.OrderedMap, quantity)
	for _, strToCmp := range strList {
		sim, err := StringsSimilarity(str, strToCmp, algo)
		if err != nil {
			return nil, err
		}

		if sim > sortedMap[sortedMap.Len()-1].Value {
			sortedMap[sortedMap.Len()-1].Key = strToCmp
			sortedMap[sortedMap.Len()-1].Value = sim
			sortedMap.SortByValues()
		}
	}

	return sortedMap.ToArray(), nil
}

// FuzzySearchSetThreshold 实现对字符串列表的近似搜索，返回一个由x个字符串组成的集合，与输入的字符串按与基本字符串的相似度排序。
// 在参数上取相似阈值。接受a quantity参数来定义所需的输出字符串数量(例如谷歌键盘词建议的情况下为3)。
// 对于与基字符串的相似度，也取一个阈值参数。
func FuzzySearchSetThreshold(str string, strList []string, quantity int, minSim float32, algo Algorithm) ([]string, error) {
	sortedMap := make(orderedmap.OrderedMap, quantity)
	for _, strToCmp := range strList {
		sim, err := StringsSimilarity(str, strToCmp, algo)
		if err != nil {
			return nil, err
		}

		if sim >= minSim && sim > sortedMap[sortedMap.Len()-1].Value {
			sortedMap[sortedMap.Len()-1].Key = strToCmp
			sortedMap[sortedMap.Len()-1].Value = sim
			sortedMap.SortByValues()
		}
	}

	return sortedMap.ToArray(), nil
}
