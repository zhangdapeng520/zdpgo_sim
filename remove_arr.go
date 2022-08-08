package zdpgo_sim

var (
	PythonRemoveArr = []string{
		"if __name__ == '__main__':",
		"^print", // ^表示以什么开头
	}
	JavaRemoveArr = []string{
		"{",
		"}",
	}
	PHPRemoveArr = []string{
		"<?php",
		"?>",
		"}",
	}
	CRemoveArr = []string{
		"{",
		"}",
	}
	CPPRemoveArr = []string{
		"{",
		"}",
		"^using namespace",
	}
)
