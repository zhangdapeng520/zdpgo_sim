package similarity

import (
	"fmt"
	"testing"
)

/*
@Time : 2022/7/28 11:26
@Author : 张大鹏
@File : simhash_test.go
@Software: Goland2021.3.1
@Description:
*/

func TestSimHash_GetToken(t *testing.T) {
	s := NewSimHash()
	token := s.GetToken("abcd")
	fmt.Println(token)
}

func TestSimHash_GetHanMingDistance(t *testing.T) {
	s := NewSimHash()
	distance := s.GetHanMingDistance("abcd", "abcd")
	fmt.Println(distance)
}
