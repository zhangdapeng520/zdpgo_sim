package zdpgo_sim

import (
	"fmt"
	"strings"

	"github.com/zhangdapeng520/zdpgo_clearcode"
	"github.com/zhangdapeng520/zdpgo_file"
	"github.com/zhangdapeng520/zdpgo_lexers"
	"github.com/zhangdapeng520/zdpgo_pool_goroutine"
	"github.com/zhangdapeng520/zdpgo_pygments"
	"github.com/zhangdapeng520/zdpgo_type/maps/safemap"
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
	lexer := zdpgo_lexers.Match(filePath)
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

// GetSourceCodeTokenArr 获取源代码的token列表
func GetSourceCodeTokenArr(lexer zdpgo_pygments.Lexer, codeStr string, splitStr string, removeArr []string) ([]string, error) {
	// 获取源码数组
	codeArr := zdpgo_clearcode.SplitCode(codeStr, splitStr, removeArr)

	// 获取token数组
	tokenArr, err := zdpgo_pygments.GetTokenArr(lexer, codeArr)
	if err != nil {
		return nil, err
	}

	// 返回
	return tokenArr, nil
}

// GetFileTokenArr 获取文件的token数组
func GetFileTokenArr(filePath string) ([]string, error) {
	// 1、将代码清洗，移除多余空行和字符串。
	content, err := zdpgo_clearcode.ClearCode(filePath)
	if err != nil {
		return nil, err
	}
	fmt.Println("源文件代码：", content)

	// 参数准备
	lexer := zdpgo_lexers.Match(filePath)
	var removeArr []string
	if strings.HasSuffix(filePath, ".py") {
		removeArr = PythonRemoveArr
		// 可选：如果是Python代码，清除main代码块
		content = zdpgo_clearcode.ClearPythonMain(content)
	}

	// 获取源码的token列表
	tokenArr, err := GetSourceCodeTokenArr(lexer, content, "\n", removeArr)
	if err != nil {
		return nil, err
	}

	// 返回
	return tokenArr, nil
}

// 将token数组按照指定的数量展开
func GetSpreadTokenArr(tokens []string, lines int) []string {
	var result []string

	// 特殊情况1：行数大于或等于token总数
	if lines >= len(tokens) {
		token := strings.Join(tokens, "")
		result = append(result, token)
		return result
	}

	// 特殊情况2：行数小于或等于0
	if lines <= 0 {
		return tokens
	}

	// 按照指定行数展开token，并合并
	for i := 0; i <= len(tokens)-lines; i++ {
		token := strings.Join(tokens[i:i+lines], "")
		result = append(result, token)
	}
	return result
}