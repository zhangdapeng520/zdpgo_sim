package stringutil

import "unicode/utf8"

// CommonPrefix 返回指定字符串的公共前缀。如果参数没有共同的前缀，则返回空字符串。
func CommonPrefix(first, second string) string {
	if utf8.RuneCountInString(first) > utf8.RuneCountInString(second) {
		first, second = second, first
	}

	var commonLen int
	sRunes := []rune(second)
	for i, r := range first {
		if r != sRunes[i] {
			break
		}

		commonLen++
	}

	return string(sRunes[0:commonLen])
}

// UniqueSlice 返回一个包含指定字符串片中唯一项的片。
// 输出切片中的项目按照它们在输入切片中出现的顺序排列。
func UniqueSlice(items []string) []string {
	var uniq []string
	registry := map[string]struct{}{}

	for _, item := range items {
		if _, ok := registry[item]; ok {
			continue
		}
		registry[item] = struct{}{}
		uniq = append(uniq, item)
	}

	return uniq
}

// SliceContains 如果terms包含q，则返回true，否则返回false。
func SliceContains(terms []string, q string) bool {
	for _, term := range terms {
		if q == term {
			return true
		}
	}

	return false
}
