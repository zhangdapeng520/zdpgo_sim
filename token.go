package zdpgo_sim

import (
	"bytes"
	"fmt"
	"github.com/zhangdapeng520/zdpgo_pool_goroutine"
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
	removeOperator  = map[string]bool{";": true, ",": true, ".": true}
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

// 处理文件token
func handleFileToken(filePath string) {
	_, _ = GetFileToken(filePath)
}

// GetToken 使用指定的词法分析器和代码内容获取token
func GetToken(lexer zdpgo_pygments.Lexer, content string) (string, error) {
	// 词法分析
	tokenise, err := lexer.Tokenise(nil, content)
	if err != nil {
		return "", err
	}

	// 处理token
	var buf bytes.Buffer

	for _, token := range tokenise.Tokens() {
		// 获取类型字符串
		typeStr := token.Type.String()
		valueStr := strings.TrimSpace(token.Value)

		// 忽略空格和注释
		if typeStr == "CommentSingle" || // 忽略单行注释
			typeStr == "CommentMultiline" || // 忽略多行注释
			typeStr == "Text" || // 忽略无法解析的文本
			typeStr == "NameNamespace" || // 忽略名称空间
			typeStr == "KeywordNamespace" || // Java名称空间关键字
			typeStr == "CommentPreproc" || // PHP前缀标识符，<?php
			typeStr == "CommentPreprocFile" || // C的<stdio.h>等标识符
			valueStr == "" { // 忽略空字符串
			continue
		} else if typeStr == "NameVariableMagic" { // 魔法方法，比如Python的__file__
			buf.WriteString("m")
		} else if typeStr == "NameDecorator" { // 注解
			//buf.WriteString("@Z")
			buf.WriteString("z")
		} else if typeStr == "KeywordConstant" { // 关键字常量，比如True
			buf.WriteString("b")
		} else if typeStr == "Keyword" ||
			typeStr == "KeywordDeclaration" { // 关键字
			buf.WriteString("k")
		} else if typeStr == "KeywordType" { // Java类型关键字
			buf.WriteString("t")
		} else if typeStr == "Operator" { // 运算符
			if _, ok := removeOperator[valueStr]; !ok {
				buf.WriteString("e")
			}
		} else if typeStr == "NameClass" { // Java：所有的类名改为C
			buf.WriteString("c")
		} else if typeStr == "NameOther" { // 所有其他变量名称为“O”
			buf.WriteString("o")
		} else if typeStr == "Name" ||
			typeStr == "NameVariable" { // 所有变量名称为“N”
			buf.WriteString("n")
		} else if typeStr == "NameAttribute" { // Java中的对象属性，全部转换为A
			buf.WriteString("a")
		} else if typeStr == "LiteralString" || // 所有双引号都变成D
			typeStr == "LiteralStringDouble" {
			buf.WriteString("d")
		} else if typeStr == "LiteralStringSingle" { // 所有单引号变成S
			buf.WriteString("s")
		} else if typeStr == "NameFunction" { // 用户定义的函数名为“F”
			buf.WriteString("f")
		} else if typeStr == "Punctuation" { // PHP的,;等运算符，C和CPP的{}等运算符
			buf.WriteString("p")
		} else if typeStr == "LiteralNumberInteger" { // 将所有的数字都替换为1，需要确认是否影响准确度
			buf.WriteString("1")
		} else if typeStr == "LiteralNumberFloat" { // 将所有的浮点数都替换为2，需要确认是否影响准确度
			buf.WriteString("2")
		} else {
			buf.WriteString(valueStr)
		}
	}
	return buf.String(), nil
}

// GetTokenArr 获取token数组
func GetTokenArr(lexer zdpgo_pygments.Lexer, contents []string) ([]string, error) {
	var results []string
	for _, content := range contents {
		// 处理PHP，解决无法获取token的问题
		isHandlePHP := lexer.Config().Name == "PHTML" && !strings.Contains(content, "<?php")
		if isHandlePHP {
			content = "<?php " + content + " ?>"
		}
		token, err := GetToken(lexer, content)
		if err != nil {
			return nil, err
		}

		// 处理token
		if isHandlePHP {
			token = strings.Replace(token, "<?php", "", 1)
			token = strings.Replace(token, "?>", "", 1)
		}

		// token不为空，且长度大于或等于3，才追加
		if token != "" && len(token) >= 3 {
			results = append(results, token)
		}
	}
	return results, nil
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
		token := strings.Join(tokens[i:i+lines], " ")
		result = append(result, token)
	}
	return result
}

func GetFileToken(filePath string) (string, error) {
	// 获取源码和后缀
	code, suffix, err := zdpgo_clearcode.GetFileContentAndSuffix(filePath)
	if err != nil {
		fmt.Println(err)
	}

	// 代码格式化
	code = zdpgo_clearcode.Format(code)

	// 代码清洗，移除多余空行和字符串。
	content, err := zdpgo_clearcode.ClearCode(code, suffix)
	if err != nil {
		// TODO: 处理错误
		fmt.Println(err)
		return "", err
	}

	// 文件路径，去除项目所在位置的绝对路径
	relativeFilePath := strings.Replace(filePath, projectDir, "", 1)

	// 文件信息
	fileInfo := FileInfo{
		FilePath:   relativeFilePath,
		FileSize:   zdpgo_file.GetFileSize(filePath),
		OriginHash: zdpgo_password.GetMd5(code),
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
			return "", err
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

		// 返回token
		return token, nil
	}

	// 返回空
	return "", nil
}

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

// GetSourceCodeTokenArr 获取源代码的token列表
func GetSourceCodeTokenArr(lexer zdpgo_pygments.Lexer, codeStr string, splitStr string, removeArr []string) ([]string, error) {
	// 获取源码数组
	codeArr := zdpgo_clearcode.SplitCode(codeStr, splitStr, removeArr)

	// 获取token数组
	tokenArr, err := GetTokenArr(lexer, codeArr)
	if err != nil {
		return nil, err
	}

	// 返回
	return tokenArr, nil
}

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
