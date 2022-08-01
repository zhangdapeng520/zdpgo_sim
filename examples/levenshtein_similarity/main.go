package main

import (
	"encoding/base64"
	"fmt"
	"github.com/zhangdapeng520/zdpgo_clearcode"
	"reflect"
	"strings"
	"unsafe"
)

const (
	ignoreCase = 1 << iota
	ignoreSpace
)

var (
	// 需要替换的字符串，是各种空格
	replace = strings.NewReplacer("\r", "", "\n", "", "\t", "", "\f", "", " ", "")
	// base64的编码表
	base64Table = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"
)

// 参数内容
type option struct {
	ignore int                         // 忽略个数
	ascii  bool                        // 设置选用ascii还是utf8方式执行算法
	cmp    func(s1, s2 string) float64 // 计算相似度的方法
	base64 bool                        // 设置是否使用base64算法
}

// OptionFunc 参数方法类型
type OptionFunc func(*option)

// Apply 执行方法
func (o OptionFunc) Apply(opt *option) {
	o(opt)
}

// EditDistance 编辑距离，Levenshtein算法对象
type EditDistance struct {
	mixed int
}

// CompareAscii 比较ascii编码，适合纯英文文本
func (e *EditDistance) CompareAscii(s1, s2 string) float64 {
	cacheX := make([]int, len(s2))

	diagonal := 0
	for y, yLen := 0, len(s1); y < yLen; y++ {
		for x, xLen := 0, len(cacheX); x < xLen; x++ {
			on := x + 1
			left := y + 1
			if x == 0 {
				diagonal = y
			} else if y == 0 {
				diagonal = x
			}
			if y > 0 {
				on = cacheX[x]
			}
			if x-1 >= 0 {
				left = cacheX[x-1]
			}

			same := 0
			if s1[y] != s2[x] {
				same = 1
			}

			oldDiagonal := cacheX[x]
			cacheX[x] = min(min(on+1, left+1), same+diagonal)
			diagonal = oldDiagonal
		}
	}

	e.mixed = cacheX[len(cacheX)-1]
	return 1.0 - float64(cacheX[len(cacheX)-1])/float64(max(len(s1), len(s2)))
}

// 计算两个int类型整数的最小值
func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

// 计算两个int类型整数的最大值
func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

// CompareUtf8 比较UTF8编码，适合非纯英文文本，比如带中文的文本
func (e *EditDistance) CompareUtf8(utf8Str1, utf8Str2 string) float64 {
	// 转换为rune数组
	r1 := []rune(utf8Str1)
	r2 := []rune(utf8Str2)

	// 缓存
	cacheX := make([]int, len(r2))

	diagonal := 0

	// 遍历字符串，进行比较
	for y, yLen := 0, len(r1); y < yLen; y++ {
		for x, xLen := 0, len(cacheX); x < xLen; x++ {
			on := x + 1
			left := y + 1
			if x == 0 {
				diagonal = y
			} else if y == 0 {
				diagonal = x
			}
			if y > 0 {
				on = cacheX[x]
			}
			if x-1 >= 0 {
				left = cacheX[x-1]
			}

			same := 0
			if r1[y] != r2[x] {
				same = 1
			}

			oldDiagonal := cacheX[x]
			cacheX[x] = min(min(on+1, left+1), same+diagonal)
			diagonal = oldDiagonal

		}
	}

	e.mixed = cacheX[len(cacheX)-1]
	return 1.0 - float64(cacheX[len(cacheX)-1])/float64(max(len(r1), len(r2)))
}

// Default 默认参数
func Default() OptionFunc {
	return OptionFunc(func(o *option) {
		if o.cmp == nil {
			l := EditDistance{}   // 编辑距离
			o.cmp = l.CompareUtf8 // 比较utf8编码
			if o.ascii {
				o.cmp = l.CompareAscii // 比较ascii编码
			}
		}
	})
}

// 调用Option接口设置option
func (o *option) fillOption(opts ...Option) {
	for _, opt := range opts {
		opt.Apply(o)
	}

	// 创建默认参数
	opt := Default()

	// 默认参数调用Apply方法
	opt.Apply(o)
}

// Option 参数接口
type Option interface {
	Apply(*option) // 参数是一个对象，这个对象要有Apply方法
}

// 修改字符串的map方法
var modiyTab = map[int]func(s *string){
	// 修改大小写
	ignoreCase: func(s *string) {
		*s = strings.ToLower(*s)
	},

	// 修改空格
	ignoreSpace: func(s *string) {
		*s = replace.Replace(*s)
	},
}

// 修改字符串
func modifyString(o *option, s *string) {
	for i := 1; i <= ignoreSpace; i <<= 1 {
		if i&o.ignore > 0 {
			modiyTab[i](s)
		}
	}
}

// StringToBytes 将字符串类型转换为[]byte类型
func StringToBytes(s string) (b []byte) {
	// 获取反射类型
	bh := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	sh := *(*reflect.StringHeader)(unsafe.Pointer(&s))

	// 修改数据，长度，容量
	bh.Data = sh.Data
	bh.Len = sh.Len
	bh.Cap = sh.Len

	// 返回
	return b
}

// Base64Encode 将byte切片转换为base64字符串
func Base64Encode(s string) string {
	base := base64.NewEncoding(base64Table)
	bytes := StringToBytes(s)
	return base.EncodeToString(bytes)
}

// 将字符串修改为base64字符串
func modifyStrToBase64Str(o *option, s *string) {
	if o.base64 {
		// 将字符串转换为base64编码
		*s = Base64Encode(*s)
	}

}

// 检查边界情况
func check(s1, s2 string) (score float64, exit bool) {
	// 两个字符串相等
	if s1 == s2 {
		return 1.0, true
	}

	// 字符串1为空
	if len(s1) == 0 {
		return 0.0, true
	}

	// 字符串2为空
	if len(s2) == 0 {
		return 0.0, true
	}

	// 不存在边界情况
	return 0, false
}

// 前处理主要涉及，修改字符串，和边界判断
func modifyStrAndCheck(o *option, s1, s2 *string) (score float64, exit bool) {
	modifyString(o, s1)         // 修改字符串1
	modifyString(o, s2)         // 修改字符串2
	modifyStrToBase64Str(o, s1) // 将字符串1转换为base64字符串
	modifyStrToBase64Str(o, s2) // 将字符串2转换为base64字符串
	return check(*s1, *s2)      // 检查边界情况
}

// 比较两个字符串内部函数
func compare(s1, s2 string, o *option) float64 {
	// 检查边界情况
	if s, e := modifyStrAndCheck(o, &s1, &s2); e {
		return s
	}

	// 进行比较
	return o.cmp(s1, s2)
}

// Compare 比较两个字符串相似度
func Compare(s1, s2 string, opts ...Option) float64 {
	var o option
	o.fillOption(opts...)
	return compare(s1, s2, &o)
}

func main() {
	suffixs := []string{
		".py",
		".php",
		".java",
		".c",
		".cpp",
	}
	for _, suffix := range suffixs {
		filePath1 := "examples/test_data/level1_1" + suffix
		filePath2 := "examples/test_data/level1_2" + suffix

		// 代码清洗
		content1, _ := zdpgo_clearcode.ClearCode(filePath1)
		content2, _ := zdpgo_clearcode.ClearCode(filePath2)

		//	莱文斯坦-编辑距离(Levenshtein)
		r1 := Compare(content1, content2)
		fmt.Println(r1)
	}

}
