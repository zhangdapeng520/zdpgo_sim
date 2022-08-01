package zdpgo_sim

// 参数内容
type option struct {
	ignore int                         // 忽略个数
	ascii  bool                        // 设置选用ascii还是utf8方式执行算法
	cmp    func(s1, s2 string) float64 // 计算相似度的方法
	base64 bool                        // 设置是否使用base64算法
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

// OptionFunc 参数方法类型
type OptionFunc func(*option)

// Apply 执行方法
func (o OptionFunc) Apply(opt *option) {
	o(opt)
}

// IgnoreCase 忽略大小写
func IgnoreCase() OptionFunc {
	return OptionFunc(func(o *option) {
		o.ignore |= ignoreCase
	})
}

// IgnoreSpace 忽略空白字符
func IgnoreSpace() OptionFunc {
	return OptionFunc(func(o *option) {
		o.ignore |= ignoreSpace
	})
}

// UseASCII 使用ascii编码
func UseASCII() OptionFunc {
	return OptionFunc(func(o *option) {
		o.ascii = true
	})
}

// UseBase64 使用base64编码
func UseBase64() OptionFunc {
	return OptionFunc(func(o *option) {
		o.base64 = true
	})
}
