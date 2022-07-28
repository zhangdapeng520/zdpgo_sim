package similarity

import (
	"hash/crc32"
	"strconv"
	"unicode/utf8"
)

// SimHash simhash算法实现
// 注意：simhash算法适用于500字以上的长文本，短文本会有偏差
type SimHash struct {
}

// NewSimHash 新建SimHash对象
func NewSimHash() *SimHash {
	return &SimHash{}
}

func (s SimHash) CompareAscii(s1, s2 string) float64 {
	return s.CompareUtf8(s1, s2)

}

func (s SimHash) CompareUtf8(utf8Str1, utf8Str2 string) float64 {
	// 字符串长度
	l1 := utf8.RuneCountInString(utf8Str1)
	l2 := utf8.RuneCountInString(utf8Str2)

	// 将字符串转换为字符数组
	s1s := StrToStrs4(utf8Str1, l1)
	s2s := StrToStrs4(utf8Str2, l2)

	// 计算每个字符在字符数组中出现的次数
	counts1 := make(map[string]int)
	counts2 := make(map[string]int)
	for _, s := range s1s {
		// 如果字符在字符数组中出现过，则计数加1
		if _, ok := counts1[s]; ok {
			counts1[s]++
		} else {
			// 如果字符在字符数组中没出现过，则计数设为1
			counts1[s] = 1
		}
	}
	for _, s := range s2s {
		if _, ok := counts2[s]; ok {
			counts2[s]++
		} else {
			counts2[s] = 1
		}
	}
	h1 := IntsToStr(Dimensionality(merge(hashcodeAndAdd(counts1))))
	h2 := IntsToStr(Dimensionality(merge(hashcodeAndAdd(counts2))))

	// 计算h1, h2的汉明距离
	hm := Hamming{}
	return hm.CompareUtf8(h1, h2)

}

// GetHanMingDistance 获取两个字符串的汉明距离
func (s *SimHash) GetHanMingDistance(utf8Str1, utf8Str2 string) float64 {
	// 获取两个字符的SimHash token
	h1 := s.GetToken(utf8Str1)
	h2 := s.GetToken(utf8Str2)

	// 计算h1, h2的汉明距离
	hm := Hamming{}
	result := hm.CompareUtf8(h1, h2)

	// 返回结果
	return result
}

// GetToken 获取文本的simhash值
func (s *SimHash) GetToken(utf8Str string) string {
	// 字符串长度
	l1 := utf8.RuneCountInString(utf8Str)

	// 将字符串转换为字符数组
	s1s := StrToStrs4(utf8Str, l1)

	// 计算每个字符在字符数组中出现的次数
	counts := make(map[string]int)
	for _, s := range s1s {
		// 如果字符在字符数组中出现过，则计数加1
		if _, ok := counts[s]; ok {
			counts[s]++
		} else {
			// 如果字符在字符数组中没出现过，则计数设为1
			counts[s] = 1
		}
	}

	token := IntsToStr(Dimensionality(merge(hashcodeAndAdd(counts))))
	return token
}

// Dimensionality 降维度
func Dimensionality(ins []int) []int {
	for i := 0; i < len(ins); i++ {
		if ins[i] > 0 {
			ins[i] = 1
		} else {
			ins[i] = 0
		}
	}
	return ins
}

// merge 合并
func merge(ins [][]int) []int {
	res := make([]int, len(ins[0]))
	lens := len(ins)
	for i := 0; i < lens; i++ {
		for j := 0; j < len(ins[i]); j++ {
			res[j] += ins[i][j]
		}
	}
	return res
}

// hashcodeAndAdd 计算hashcode并加权
func hashcodeAndAdd(counts map[string]int) [][]int {
	// hashmap
	lens := len(counts)
	h1 := make([][]int, lens)

	// 计算counts1,counts2 中每个字符的hash值, 并且将出现的次数分为5个等级, 将每个字符的hash值与出现的次数等级相乘
	c1 := (lens - 1) * 4.0
	j := 0
	for k, v := range counts {
		// 计算每一个字符串的hash
		// 出现的次数除以5
		c := strconv.FormatUint(uint64(crc32.ChecksumIEEE([]byte(k))), 2)

		// 将字符串转换为数字数组
		cs := Int32StrToInts(c)
		if v <= c1/5.0 {
			// 加权
			h1[j] = Add(cs, 1)
		} else if v <= c1/5.0*2 {
			// 加权
			h1[j] = Add(cs, 2)
		} else if v <= c1/5.0*3 {
			// 加权
			h1[j] = Add(cs, 3)
		} else if v <= c1/5.0*4 {
			// 加权
			h1[j] = Add(cs, 4)
		} else {
			// 加权
			h1[j] = Add(cs, 5)
		}
		j++
	}

	return h1
}
