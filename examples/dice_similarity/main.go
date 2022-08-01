package main

import (
	"encoding/base64"
	"fmt"
	"github.com/zhangdapeng520/zdpgo_clearcode"
	"reflect"
	"strings"
	"unicode/utf8"
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

// DiceCoefficientFunc ngram 是筛子系数需要用的一个值
func DiceCoefficientFunc(ngram ...int) OptionFunc {
	f := func(o *option) {
		ngram2 := 2
		if len(ngram) > 0 {
			ngram2 = ngram[0]
		}

		// 创建算法对象
		d := &DiceCoefficient{Ngram: ngram2}

		// 赋值比较方法
		o.cmp = d.CompareUtf8
	}

	// 将func类型的f转换为OptionFunc类型
	return OptionFunc(f)
}

type DiceCoefficient struct {
	Ngram int // k-gram 分段数量
	l1    int
	l2    int
	mixed int
	key   []string
	test  bool
}

type value struct {
	s1Count int
	s2Count int
}

func (d *DiceCoefficient) CompareAscii(s1, s2 string) float64 {
	return d.CompareUtf8(s1, s2)
}

func (d *DiceCoefficient) setOrGet(set map[string]value, s string, add bool) (mixed, l int) {
	var key strings.Builder
	ngram := d.Ngram
	if ngram == 0 {
		ngram = 2
	}

	for i := 0; i < len(s); {
		firstSize := 0
		for j, total := 0, 0; j < ngram && i+total < len(s); j++ {
			r, size := utf8.DecodeRuneInString(s[i+total:])
			key.WriteRune(r)
			total += size
			if j == 0 {
				firstSize = size
			}

		}
		if utf8.RuneCountInString(key.String()) != ngram {
			break
		}
		val, ok := set[key.String()]
		if add {
			if !ok {
				val = value{}
			}
			val.s1Count++
		} else {

			if !ok {
				goto next
			}

			val.s2Count++
			if val.s1Count >= val.s2Count {
				mixed++
			}
		}

		set[key.String()] = val

	next:
		if d.test {
			d.key = append(d.key, key.String())
		}

		key.Reset()
		l++
		i += firstSize
	}

	return mixed, l
}

// CompareUtf8 比较UTF8编码的字符串
func (d *DiceCoefficient) CompareUtf8(s1, s2 string) float64 {

	set := make(map[string]value, len(s1)/3)

	mixed, l1 := d.setOrGet(set, s1, true)
	mixed, l2 := d.setOrGet(set, s2, false)

	d.l1 = l1
	d.l2 = l2
	d.mixed = mixed
	return 2.0 * float64(mixed) / float64(l1+l2)
}

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

// 调用Option接口设置option
func (o *option) fillOption(opts ...Option) {
	for _, opt := range opts {
		opt.Apply(o)
	}
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

		// Dice相似系数
		r1 := Compare(content1, content2, DiceCoefficientFunc())
		fmt.Println(r1)
	}

}
