package main

import (
	"fmt"

	"github.com/zhangdapeng520/zdpgo_sim/stringosim"
)

func main() {
	fmt.Println(stringosim.Levenshtein([]rune("stringosim"), []rune("stingobim")))

	fmt.Println(stringosim.Levenshtein([]rune("stringosim"), []rune("stingobim"),
		stringosim.LevenshteinSimilarityOptions{
			InsertCost:     3,
			DeleteCost:     5,
			SubstituteCost: 2,
		}))

	fmt.Println(stringosim.Levenshtein([]rune("stringosim"), []rune("STRINGOSIM"),
		stringosim.LevenshteinSimilarityOptions{
			InsertCost:      3,
			DeleteCost:      4,
			SubstituteCost:  5,
			CaseInsensitive: true,
		}))

	fmt.Println(stringosim.Jaccard([]rune("stringosim"), []rune("stingobim"), []int{2}))

	fmt.Println(stringosim.Jaccard([]rune("stringosim"), []rune("stingobim"), []int{3}))

	dis, _ := stringosim.Hamming([]rune("testing"), []rune("restink"))
	fmt.Println(dis)

	dis, _ = stringosim.Hamming([]rune("testing"), []rune("FESTING"), stringosim.HammingSimilarityOptions{
		CaseInsensitive: true,
	})
	fmt.Println(dis)

	_, err := stringosim.Hamming([]rune("testing"), []rune("testin"))
	fmt.Println(err)

	fmt.Println(stringosim.Jaro([]rune("abaccbabaacbcb"), []rune("bababbcabbaaca")))
	fmt.Println(stringosim.Jaro([]rune("abaccbabaacbcb"), []rune("ABABAbbCABbaACA"),
		stringosim.JaroSimilarityOptions{
			CaseInsensitive: true,
		}))

	fmt.Println(stringosim.JaroWinkler([]rune("abaccbabaacbcb"), []rune("bababbcabbaaca")))
	fmt.Println(stringosim.JaroWinkler([]rune("abaccbabaacbcb"), []rune("BABAbbCABbaACA"),
		stringosim.JaroSimilarityOptions{
			CaseInsensitive: true,
			Threshold:       0.7,
			PValue:          0.1,
			LValue:          4,
		}))
	fmt.Println(stringosim.QGram([]rune("abcde"), []rune("abdcde")))

	fmt.Println(stringosim.QGram([]rune("abcde"), []rune("ABDCDE"),
		stringosim.QGramSimilarityOptions{
			CaseInsensitive: true,
			NGramSizes:      []int{3},
		}))

	fmt.Println(stringosim.QGram([]rune("abcde"), []rune("abdcde")))

	fmt.Println(stringosim.QGram([]rune("abcde"), []rune("ABDCDE"),
		stringosim.QGramSimilarityOptions{
			CaseInsensitive: true,
			NGramSizes:      []int{3},
		}))

	fmt.Println(stringosim.LCS([]rune("abcde"), []rune("abdcde")))
}
