package zdpgo_sim

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/zhangdapeng520/zdpgo_pool_goroutine"
	"io/ioutil"
	"path/filepath"
	"sort"
	"strings"

	"github.com/zhangdapeng520/zdpgo_clearcode"
	"github.com/zhangdapeng520/zdpgo_file"
	"github.com/zhangdapeng520/zdpgo_lexers"
	"github.com/zhangdapeng520/zdpgo_password"
	"github.com/zhangdapeng520/zdpgo_pygments"
	"github.com/zhangdapeng520/zdpgo_type/maps/safemap"
)

var (
	projectTokenMap = new(safemap.SafeMap[string, FileInfo])
	projectDir      string
	lexerMap        = map[string]zdpgo_pygments.Lexer{
		".py":   zdpgo_lexers.Get("Python"),
		".java": zdpgo_lexers.Get("Java"),
		".php":  zdpgo_lexers.Get("PHP"),
		".c":    zdpgo_lexers.Get("C"),
		".cpp":  zdpgo_lexers.Get("C++"),
	}
)

type FileInfo struct {
	FilePath         string   `json:"file_path"`          // 文件路径
	FileSize         int64    `json:"file_size"`          // 文件大小
	OriginHash       string   `json:"origin_hash"`        // 原始文件hash
	ClearCode        string   `json:"clear_code"`         // 清洗后代码
	ClearHash        string   `json:"clear_hash"`         // 清洗后代码的hash
	TokenArr         []string `json:"token_arr"`          // 词法分析并按空格拆分后的token数组
	TokenHashArr     []string `json:"token_hash_arr"`     // 词法分析并按空格拆分后的token数组对应的hash数组
	TokenContent     string   `json:"token_content"`      // token数组用空格拼接
	TokenContentHash string   `json:"token_content_hash"` // token_content的hash
	HashContent      string   `json:"hash_content"`       // hash数组用空格拼接
}

func handleFileToken(filePath string) {
	// 1、将代码清洗，移除多余空行和字符串。
	content, err := zdpgo_clearcode.ClearCode(filePath)
	if err != nil {
		// TODO: 处理错误
		fmt.Println(err)
		return
	}

	// 文件信息
	originContent, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println(err)
		return
	}

	// 文件路径，去除项目所在位置的绝对路径
	relativeFilePath := strings.Replace(filePath, projectDir, "", 1)

	// 文件信息
	fileInfo := FileInfo{
		FilePath:   relativeFilePath,
		FileSize:   zdpgo_file.GetFileSize(filePath),
		OriginHash: zdpgo_password.GetMd5(string(originContent)),
		ClearCode:  content,
		ClearHash:  zdpgo_password.GetMd5(content),
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
			fmt.Println(err)
			return
		}
		fileInfo.TokenArr = tokenArr

		// token列表对应的hash列表
		var tokenHashArr []string
		for _, tk := range tokenArr {
			tokenHashArr = append(tokenHashArr, zdpgo_password.GetMd5(tk))
		}
		fileInfo.TokenHashArr = tokenHashArr

		// 按空格拼接token
		token := strings.Join(tokenArr, " ")

		// 添加token和hash
		fileInfo.TokenContent = token
		fileInfo.TokenContentHash = zdpgo_password.GetMd5(token)

		// 按空格拼接hash数组
		hashContent := strings.Join(tokenHashArr, " ")
		fileInfo.HashContent = hashContent

		// 将数据添加到map
		projectTokenMap.Set(relativeFilePath, fileInfo)
	}
}

// GetFileToken 获取文件token
func GetFileToken(filePath string) (string, error) {
	// 1、将代码清洗，移除多余空行和字符串。
	content, err := zdpgo_clearcode.ClearCode(filePath)
	if err != nil {
		return "", err
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
			return "", err
		}

		// 按空格拼接token
		token := strings.Join(tokenArr, " ")
		return token, nil
	}

	fmt.Println("不支持的文件类型：", filePath)
	return "", errors.New("不支持的文件类型")
}

//
//func GetFileTokenFromArr(filePath string) (string, error) {
//	// 获取token列表
//	tokens, err := GetFileTokenArr(filePath)
//	if err != nil {
//		return "", err
//	}
//
//	// 合并token列表
//	token := strings.Join(tokens, " ")
//
//	// 返回
//	return token, nil
//}
//
//// getFileTokenArr 获取文件的token数组
//func getFileTokenArr(filePath string) {
//	// 代码清洗
//	tokenArr, err := GetFileTokenArr(filePath)
//	if err != nil {
//		fmt.Println("获取文件token列表失败：", err)
//	}
//
//	// 处理文件路径
//	filePath = strings.Replace(filePath, projectDir, "", -1)
//
//	// 用空格合并token数组
//	token := strings.Join(tokenArr, " ")
//
//	projectTokenArrMap.Set(filePath, tokenArr)
//}

// GetProjectFileInfo 获取项目的源码文件信息
func GetProjectFileInfo(
	projectPath string,
	poolSize int,
	codeSuffix string,
	ignoreDirs []string) (*safemap.SafeMap[string, FileInfo], error) {
	// 项目路径
	projectDir = projectPath

	// 清空全局变量
	defer func() {
		projectTokenMap = new(safemap.SafeMap[string, FileInfo])
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
		ignoreDirs = []string{"venv", ".git", ".idea", ".vscode"}
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
	zdpgo_pool_goroutine.RunBatchArgTask[string](poolSize, handleFileToken, filePathList)

	// 返回
	return projectTokenMap, nil
}

//func GetProjectTokenArr(
//	projectPath string,
//	poolSize int,
//	codeSuffix string,
//	ignoreDirs []string) (*safemap.SafeMap[string, []string], error) {
//
//	// 清空全局变量
//	defer func() {
//		projectTokenArrMap = new(safemap.SafeMap[string, []string])
//		projectDir = ""
//	}()
//
//	// 项目路径
//	projectDir = projectPath
//
//	// 清空全局变量
//	defer func() {
//		projectTokenMap = new(safemap.SafeMap[string, string])
//		projectDir = ""
//	}()
//
//	// 默认Goroutine池大小
//	if poolSize <= 0 {
//		poolSize = 333333
//	}
//
//	// 默认源码后缀
//	if codeSuffix == "" {
//		codeSuffix = ".py"
//	}
//
//	// 默认被忽略的文件夹
//	if ignoreDirs == nil {
//		ignoreDirs = []string{"venv", ".git", ".idea"}
//	}
//
//	// 所有需要生成token的文件列表
//	var filePathList []string
//
//	// 获取需要生成token的文件
//	handleFunc := func(filePath string) {
//		if strings.HasSuffix(filePath, codeSuffix) {
//			filePathList = append(filePathList, filePath)
//		}
//	}
//
//	// 执行处理方法
//	err := zdpgo_file.HandleDirFileWithIgnoreDirList(projectPath, handleFunc, ignoreDirs)
//	if err != nil {
//		return nil, err
//	}
//
//	// 使用Goroutine协程池，并发的生成token
//	zdpgo_pool_goroutine.RunBatchArgTask[string](poolSize, getFileTokenArr, filePathList)
//
//	// 返回
//	return projectTokenArrMap, nil
//}

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

//// GetFileTokenArr 获取文件的token数组
//func GetFileTokenArr(filePath string) ([]string, error) {
//	// 1、将代码清洗，移除多余空行和字符串。
//	content, err := zdpgo_clearcode.ClearCode(filePath)
//	if err != nil {
//		return nil, err
//	}
//
//	// 参数准备
//	if lexer, ok := lexerMap[filepath.Ext(filePath)]; ok {
//		var removeArr []string
//
//		// 处理不同编程语言的特殊内容
//		if strings.HasSuffix(filePath, ".py") {
//			removeArr = PythonRemoveArr
//			// 可选：如果是Python代码，清除main代码块
//			content = zdpgo_clearcode.ClearPythonMain(content)
//		} else if strings.HasSuffix(filePath, ".java") {
//			removeArr = JavaRemoveArr
//		} else if strings.HasSuffix(filePath, ".php") {
//			removeArr = PHPRemoveArr
//		} else if strings.HasSuffix(filePath, ".c") {
//			removeArr = CRemoveArr
//		} else if strings.HasSuffix(filePath, ".cpp") {
//			removeArr = CPPRemoveArr
//		}
//
//		// 获取源码的token列表
//		tokenArr, err := GetSourceCodeTokenArr(lexer, content, "\n", removeArr)
//		if err != nil {
//			return nil, err
//		}
//
//		// 返回
//		return tokenArr, nil
//	} else {
//		fmt.Println("不支持的文件类型：", filePath)
//		return nil, errors.New("不支持的文件类型")
//	}
//}
//
//// GetSpreadTokenArr 将token数组按照指定的数量展开
//func GetSpreadTokenArr(tokens []string, lines int) []string {
//	var result []string
//
//	// 特殊情况1：行数大于或等于token总数
//	if lines >= len(tokens) {
//		token := strings.Join(tokens, "")
//		result = append(result, token)
//		return result
//	}
//
//	// 特殊情况2：行数小于或等于0
//	if lines <= 0 {
//		return tokens
//	}
//
//	// 按照指定行数展开token，并合并
//	for i := 0; i <= len(tokens)-lines; i++ {
//		token := strings.Join(tokens[i:i+lines], "")
//		result = append(result, token)
//	}
//	return result
//}

// GetProjectToken 获取项目的token
func GetProjectToken(tokenMap *safemap.SafeMap[string, FileInfo]) string {
	// 按照文件名排序
	keys := tokenMap.Keys()
	sort.Strings(keys)

	// 遍历并生成大的token
	var tokenBuffer bytes.Buffer
	for _, k := range keys {
		fileInfo := tokenMap.Get(k)
		token := strings.Replace(fileInfo.TokenContent, " ", "", -1)
		tokenBuffer.WriteString(token)
		tokenBuffer.WriteString(" ")
	}

	// 返回
	return tokenBuffer.String()
}

// GetProjectHash 获取项目的每个文件的hash，按空格拼接为一个字符串
func GetProjectHash(tokenMap *safemap.SafeMap[string, FileInfo]) string {
	// 按照文件名排序
	keys := tokenMap.Keys()
	sort.Strings(keys)

	// 遍历并生成大的token
	var tokenBuffer bytes.Buffer
	for _, k := range keys {
		fileInfo := tokenMap.Get(k)
		tokenBuffer.WriteString(fileInfo.TokenContentHash)
		tokenBuffer.WriteString(" ")
	}

	// 返回
	return tokenBuffer.String()
}

//
//// GetProjectTokenSplitArr 将token字典按照指定的精准度，转换为token数组
//func GetProjectTokenSplitArr(tokenMap *safemap.SafeMap[string, string], splitNum int) []string {
//	// 按照文件名排序
//	keys := tokenMap.Keys()
//	sort.Strings(keys)
//
//	// [1,2,3,4] 4
//	// [1,2,3],[2,3,4]
//	var tokens []string
//	for i := 0; i <= len(keys)-splitNum; i++ {
//		var tokenBuffer bytes.Buffer
//		for j := i; j < i+splitNum; j++ {
//			tokenBuffer.WriteString(tokenMap.Get(keys[i]))
//			tokenBuffer.WriteString(" ")
//		}
//		tokens = append(tokens, tokenBuffer.String())
//	}
//
//	// 返回
//	return tokens
//}
