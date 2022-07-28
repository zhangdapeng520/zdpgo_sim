package main

import (
	"fmt"
	"github.com/zhangdapeng520/zdpgo_sim"
)

/*
@Time : 2022/7/27 20:45
@Author : 张大鹏
@File : main.go
@Software: Goland2021.3.1
@Description:
*/
func main() {
	filePathList := []string{
		"examples/test_data/level1_1.py",
		"examples/test_data/level1_1.java",
		"examples/test_data/level1_1.php",
		"examples/test_data/level1_1.c",
		"examples/test_data/level1_1.cpp",
		"examples/test_data/level1_2.py",
		"examples/test_data/level1_2.java",
		"examples/test_data/level1_2.php",
		"examples/test_data/level1_2.c",
		"examples/test_data/level1_2.cpp",
	}

	// 清除代码
	for _, filePath := range filePathList {
		fmt.Println(filePath)
		result, err := zdpgo_sim.ClearCode(filePath)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(result)
		fmt.Println("========================")
	}
}
