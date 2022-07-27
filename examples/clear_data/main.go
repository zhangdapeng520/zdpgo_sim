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
		"examples/clear_data/test_data/demo.py",
		"examples/clear_data/test_data/demo.java",
		"examples/clear_data/test_data/demo.php",
		"examples/clear_data/test_data/demo.c",
		"examples/clear_data/test_data/demo.cpp",
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
