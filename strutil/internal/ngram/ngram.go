package ngram

import "github.com/zhangdapeng520/zdpgo_sim/strutil/internal/mathutil"

// Count 返回提供的术语的指定大小的n元计数。
// 如果提供的大小小于或等于0，则使用n-grams大小为1的值。
func Count(runes []rune, size int) int {
	return mathutil.Max(len(runes)-(mathutil.Max(size, 1)-1), 0)
}

// Slice 返回所提供术语的指定大小的所有n元。
// 输出切片中的n-grams是按照它们在输入项中出现的顺序排列的。
// 如果提供的大小小于或等于0，则使用n-grams大小为1的值。
func Slice(runes []rune, size int) []string {
	// Use an n-gram size of 1 if the provided size is invalid.
	size = mathutil.Max(size, 1)

	// Check if term length is too small.
	lenRunes := len(runes)
	if lenRunes == 0 || lenRunes < size {
		return nil
	}

	// Generate n-gram slice.
	limit := lenRunes - (size - 1)
	ngrams := make([]string, limit)

	for i, j := 0, 0; i < limit; i++ {
		ngrams[j] = string(runes[i : i+size])
		j++
	}

	return ngrams
}

// Map 返回所提供术语的指定大小的所有n-gram的映射，以及它们的频率。
// 该函数还返回n-gram的总数，即输出映射中所有值的总和。
// 如果提供的大小小于或等于0，则使用n-gram大小为1的值。
func Map(runes []rune, size int) (map[string]int, int) {
	// Use an n-gram size of 1 if the provided size is invalid.
	size = mathutil.Max(size, 1)

	// Check if term length is too small.
	lenRunes := len(runes)
	if lenRunes == 0 || lenRunes < size {
		return map[string]int{}, 0
	}

	// Generate n-gram map.
	limit := lenRunes - (size - 1)
	ngrams := make(map[string]int, limit)

	var ngramCount int
	for i := 0; i < limit; i++ {
		ngram := string(runes[i : i+size])
		count, _ := ngrams[ngram]
		ngrams[ngram] = count + 1
		ngramCount++
	}

	return ngrams, ngramCount
}

// Intersection 返回在这两个术语中找到的指定大小的n-grams的映射，以及它们的频率。
// 该函数还返回公共n-gram的数量(输出映射中所有值的总和)、第一项中的n-gram总数和第二项中的n-gram总数。
// 如果提供的大小小于或等于0，则使用n-gram大小为1的值。
func Intersection(a, b []rune, size int) (map[string]int, int, int, int) {
	// Use an n-gram size of 1 if the provided size is invalid.
	size = mathutil.Max(size, 1)

	// Compute the n-grams of the first term.
	ngramsA, totalA := Map(a, size)

	// Calculate n-gram intersection with the second term.
	limit := len(b) - (size - 1)
	commonNgrams := make(map[string]int, mathutil.Max(limit, 0))

	var totalB, intersection int
	for i := 0; i < limit; i++ {
		ngram := string(b[i : i+size])
		totalB++

		if count, ok := ngramsA[ngram]; ok && count > 0 {
			// Decrease frequency of n-gram found in the first term each time
			// a successful match is found.
			intersection++
			ngramsA[ngram] = count - 1

			// Update common n-grams map with the matched n-gram and its
			// frequency.
			count, _ = commonNgrams[ngram]
			commonNgrams[ngram] = count + 1
		}
	}

	return commonNgrams, intersection, totalA, totalB
}
