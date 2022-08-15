package zdpgo_sim

import (
	"bytes"
	"crypto/md5"
	"errors"
	"fmt"
	"path/filepath"
	"sort"
	"strings"

	"github.com/zhangdapeng520/zdpgo_clearcode"
	"github.com/zhangdapeng520/zdpgo_file"
	"github.com/zhangdapeng520/zdpgo_lexers"
	"github.com/zhangdapeng520/zdpgo_pool_goroutine"
	"github.com/zhangdapeng520/zdpgo_pygments"
	"github.com/zhangdapeng520/zdpgo_type/maps/safemap"
)

var (
	projectTokenMap    = new(safemap.SafeMap[string, string])
	projectTokenArrMap = new(safemap.SafeMap[string, []string])
	projectDir         string
	lexerMap           = map[string]zdpgo_pygments.Lexer{
		".py":   zdpgo_lexers.Get("Python"),
		".java": zdpgo_lexers.Get("Java"),
		".php":  zdpgo_lexers.Get("PHP"),
		".c":    zdpgo_lexers.Get("C"),
		".cpp":  zdpgo_lexers.Get("C++"),
	}
)

func getFileToken(filePath string) {
	// 代码清洗
	content, err := zdpgo_clearcode.ClearCode(filePath)
	if err != nil {
		fmt.Println("代码清洗失败：", err)
		return
	}

	// 词法分析获取token
	//lexer := zdpgo_lexers.Match(filePath)
	if lexer, ok := lexerMap[filepath.Ext(filePath)]; ok {
		token, err := zdpgo_pygments.GetToken(lexer, content)
		if err != nil {
			fmt.Println("词法分析获取token失败：", err)
			return
		}

		// 添加token
		filePath = strings.Replace(filePath, projectDir, "", -1)
		projectTokenMap.Set(filePath, token)
	} else {
		fmt.Println("不支持的文件类型：", filePath)
		return
	}
}

// getFileTokenArr 获取文件的token数组
func getFileTokenArr(filePath string) {
	// 代码清洗
	tokenArr, err := GetFileTokenArr(filePath)
	if err != nil {
		fmt.Println("获取文件token列表失败：", err)
	}

	// 添加
	filePath = strings.Replace(filePath, projectDir, "", -1)
	projectTokenArrMap.Set(filePath, tokenArr)
}

// GetProjectTokenMap 获取项目的token字典
func GetProjectTokenMap(
	projectPath string,
	poolSize int,
	codeSuffix string,
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
	zdpgo_pool_goroutine.RunBatchArgTask[string](poolSize, getFileToken, filePathList)

	// 返回
	return projectTokenMap, nil
}

func GetProjectTokenArr(
	projectPath string,
	poolSize int,
	codeSuffix string,
	ignoreDirs []string) (*safemap.SafeMap[string, []string], error) {

	// 清空全局变量
	defer func() {
		projectTokenArrMap = new(safemap.SafeMap[string, []string])
		projectDir = ""
	}()

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
	zdpgo_pool_goroutine.RunBatchArgTask[string](poolSize, getFileTokenArr, filePathList)

	// 返回
	return projectTokenArrMap, nil
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

	// 参数准备
	if lexer, ok := lexerMap[filepath.Ext(filePath)]; ok {
		var removeArr []string

		// 处理不同编程语言的特殊内容
		if strings.HasSuffix(filePath, ".py") {
			removeArr = PythonRemoveArr
			// 可选：如果是Python代码，清除main代码块
			content = zdpgo_clearcode.ClearPythonMain(content)
		} else if strings.HasSuffix(filePath, ".java") {
			removeArr = JavaRemoveArr
		} else if strings.HasSuffix(filePath, ".php") {
			removeArr = PHPRemoveArr
		} else if strings.HasSuffix(filePath, ".c") {
			removeArr = CRemoveArr
		} else if strings.HasSuffix(filePath, ".cpp") {
			removeArr = CPPRemoveArr
		}

		// 获取源码的token列表
		tokenArr, err := GetSourceCodeTokenArr(lexer, content, "\n", removeArr)
		if err != nil {
			return nil, err
		}

		// 返回
		return tokenArr, nil
	} else {
		fmt.Println("不支持的文件类型：", filePath)
		return nil, errors.New("不支持的文件类型")
	}
}

// GetSpreadTokenArr 将token数组按照指定的数量展开
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

// GetProjectToken 获取项目的token
func GetProjectToken(tokenMap *safemap.SafeMap[string, string]) string {
	// 按照文件名排序
	keys := tokenMap.Keys()
	sort.Strings(keys)

	// 遍历并生成大的token
	var tokenBuffer bytes.Buffer
	for _, k := range keys {
		tokenBuffer.WriteString(tokenMap.Get(k))
		tokenBuffer.WriteString("\n")
	}

	// 返回
	return tokenBuffer.String()
}

// GetProjectTokenSplitArr 将token字典按照指定的精准度，转换为token数组
func GetProjectTokenSplitArr(tokenMap *safemap.SafeMap[string, string], splitNum int) []string {
	// 按照文件名排序
	keys := tokenMap.Keys()
	sort.Strings(keys)

	// [1,2,3,4] 4
	// [1,2,3],[2,3,4]
	var tokens []string
	for i := 0; i <= len(keys)-splitNum; i++ {
		var tokenBuffer bytes.Buffer
		for j := i; j < i+splitNum; j++ {
			tokenBuffer.WriteString(tokenMap.Get(keys[i]))
			tokenBuffer.WriteString(" ")
		}
		tokens = append(tokens, tokenBuffer.String())
	}

	// 返回
	return tokens
}

// GetMd5 获取一个文本的md5值
func GetMd5(text string) string {
	data := []byte(text)
	has := md5.Sum(data)
	md5str1 := fmt.Sprintf("%x", has) //将[]byte转成16进制
	return md5str1
}
