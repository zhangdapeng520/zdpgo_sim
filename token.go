package zdpgo_sim

import (
	"fmt"
	"github.com/zhangdapeng520/zdpgo_clearcode"
	"github.com/zhangdapeng520/zdpgo_file"
	"github.com/zhangdapeng520/zdpgo_pool_goroutine"
	"github.com/zhangdapeng520/zdpgo_pygments"
	"github.com/zhangdapeng520/zdpgo_pygments/lexers"
	"github.com/zhangdapeng520/zdpgo_type/maps/safemap"
	"strings"
)

var (
	projectTokenMap = new(safemap.SafeMap[string, string])
	projectDir      string
)

func handleFile(filePath string) {
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

	// 添加token
	filePath = strings.Replace(filePath, projectDir, "", -1)
	projectTokenMap.Set(filePath, token)
}

func GetProjectToken(projectPath string, poolSize int, codeSuffix string,
	ignoreDirs []string) (*safemap.SafeMap[string, string], error) {
	// 项目路径
	projectDir = projectPath

	// 清空全局变量
	defer func() {
		projectTokenMap = new(safemap.SafeMap[string, string])
		projectDir = ""
	}()

	// 默认Goroutine池大小
	if poolSize <= 0 {
		poolSize = 333333
	}

	// 默认源码后缀
	if codeSuffix == "" {
		codeSuffix = ".py"
	}

	// 默认被忽略的文件夹
	if ignoreDirs == nil {
		ignoreDirs = []string{"venv", ".git", ".idea"}
	}

	// 所有需要生成token的文件列表
	var filePathList []string

	// 获取需要生成token的文件
	handleFunc := func(filePath string) {
		if strings.HasSuffix(filePath, codeSuffix) {
			filePathList = append(filePathList, filePath)
		}
	}

	// 执行处理方法
	err := zdpgo_file.HandleDirFileWithIgnoreDirList(projectPath, handleFunc, ignoreDirs)
	if err != nil {
		return nil, err
	}

	// 使用Goroutine协程池，并发的生成token
	zdpgo_pool_goroutine.RunBatchArgTask[string](poolSize, handleFile, filePathList)

	// 返回
	return projectTokenMap, nil
}
