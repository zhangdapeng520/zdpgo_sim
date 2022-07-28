package simhash

import (
	"math"
	"strconv"
	"strings"
)

/*
@Time : 2022/7/28 11:46
@Author : 张大鹏
@File : simhash.go
@Software: Goland2021.3.1
@Description: simhash算法实现
*/

type SimHash struct {
	IntSimHash int64
	HashBits   int
}

// HammingDistance 获取汉明距离
func (s *SimHash) HammingDistance(hash, other int64) int {
	x := (hash ^ other) & ((1 << uint64(s.HashBits)) - 1)
	tot := 0
	for x != 0 {
		tot += 1
		x &= x - 1
	}
	return tot
}

// Similarity 计算相似度
func (s *SimHash) Similarity(hash, other int64) float64 {
	a := float64(hash)
	b := float64(other)
	if a > b {
		return b / a
	}
	return a / b
}

// Simhash 获取文档签名
func (s *SimHash) Simhash(str string) int64 {
	m := strings.Split(str, " ")

	tokenInt := make([]int, s.HashBits)
	for i := 0; i < len(m); i++ {
		temp := m[i]
		t := s.Hash(temp)
		for j := 0; j < s.HashBits; j++ {
			bitMask := int64(1 << uint(j))
			if t&bitMask != 0 {
				tokenInt[j] += 1
			} else {
				tokenInt[j] -= 1
			}
		}

	}
	var fingerprint int64 = 0
	for i := 0; i < s.HashBits; i++ {
		if tokenInt[i] >= 0 {
			fingerprint += 1 << uint64(i)
		}
	}
	return fingerprint
}

// Params 初始化
func Params() (s *SimHash) {
	s = &SimHash{}
	s.HashBits = 64
	return s
}

// Hash 计算token的hash值
func (s *SimHash) Hash(token string) int64 {
	if token == "" {
		return 0
	} else {
		x := int64(int(token[0]) << 7)
		m := int64(1000003)
		mask := math.Pow(2, float64(s.HashBits-1))
		s := strconv.FormatFloat(mask, 'f', -1, 64)
		tsk, _ := strconv.ParseInt(s, 10, 64)
		for i := 0; i < len(token); i++ {
			tokens := int64(int(token[0]))
			x = ((x * m) ^ tokens) & tsk
		}
		x ^= int64(len(token))
		if x == -1 {
			x = -2
		}
		return int64(x)
	}
}
