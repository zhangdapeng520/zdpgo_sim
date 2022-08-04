package main

import (
	"fmt"
	"strings"

	"github.com/zhangdapeng520/zdpgo_clearcode"
	"github.com/zhangdapeng520/zdpgo_file"
	"github.com/zhangdapeng520/zdpgo_pygments"
	"github.com/zhangdapeng520/zdpgo_pygments/lexers"
)

func main() {
	// 获取指定项目中每个文件的token
	handleFunc := func(filePath string) {
		if strings.HasSuffix(filePath, ".py") {
			// 代码清洗
			content, err := zdpgo_clearcode.ClearCode(filePath)
			if err != nil {
				fmt.Println("代码清洗失败：", err)
				return
			}

			// 词法分析获取token
			lexer := lexers.Match(filePath)
			token, err := zdpgo_pygments.GetToken(lexer, content)
			if err != nil {
				fmt.Println("词法分析获取token失败：", err)
				return
			}

			fmt.Println(filePath)
			fmt.Println(token)
			fmt.Println()
		}
	}

	// 被忽略的文件夹
	ignoreDirs := []string{"venv", ".git", ".idea"}

	// 执行处理方法
	err := zdpgo_file.HandleDirFileWithIgnoreDirList("D:\\zdppy\\django_for_api_4.0", handleFunc, ignoreDirs)
	if err != nil {
		fmt.Println("处理项目源码出错：", err)
	}
}
