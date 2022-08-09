package zdpgo_sim

var (
	PythonRemoveArr = []string{
		"if __name__ == '__main__':",
		"^print", // ^表示以什么开头
		"]",
		"{",
		"},",
		"}",
		")",
		"],",
		"),",
	}
	JavaRemoveArr = []string{
		"{",
		"}",
	}
	PHPRemoveArr = []string{
		"<?php",
		"?>",
		"{",
		"}",
		"],",
		"];",
		"],",
		"});",
		"{}",
		"%}else{", // % 表示清空空格后相等
		"^use",    // ^ 表示以use开头
		"^namespace",
		"%return[",
	}
	CRemoveArr = []string{
		"{",
		"}",
		"%}else{",
		"%else",
		"return1;",
		"};",
	}
	CPPRemoveArr = []string{
		"{",
		"}",
		"^using namespace",
		"%}else{",
		"%else",
		"};",
	}
)
