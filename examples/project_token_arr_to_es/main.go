package main

import (
	"bytes"
	"crypto/md5"
	"fmt"
	"github.com/zhangdapeng520/zdpgo_es"
	"github.com/zhangdapeng520/zdpgo_sim"
	"github.com/zhangdapeng520/zdpgo_uuid"
)

type projectToken struct {
	ProjectName       string   `json:"project_name"`        // 项目名
	Language          string   `json:"language"`            // 编程语言
	Suffix            string   `json:"suffix"`              // 文件后缀
	OpenSourceAddress string   `json:"open_source_address"` // 开源地址
	FilePath          string   `json:"file_path"`           // 源码路径
	OriginHash        string   `json:"origin_hash"`         // 原始文件hash
	ClearHash         string   `json:"clear_hash"`          // 清洗后文件hash
	Tokens            []string `json:"tokens,omitempty"`    // token列表
	TokenContent      string   `json:"token_content"`       // 未hash之前的token按空格拼接
}

func main() {
	projectName := "RuoYi"
	openSourceAddress := "https://gitee.com/y_project/RuoYi.git"
	projectDir := "/data/tmp/RuoYi"
	language := "java"
	suffix := ".java"
	ignoreDirs := []string{"out", ".mvn", ".git", ".idea"}

	// 连接池数量
	// 这里可以给每个文件分配一个Goroutine，还要根据内存来，1G内存可以8W个（需自行测试）
	poolSize := 10000

	// 获取token
	projectTokenArrMap, err := zdpgo_sim.GetProjectTokenArr(
		projectDir,
		poolSize,
		suffix,
		ignoreDirs)
	if err != nil {
		fmt.Println("获取项目token失败：", err)
		return
	}

	var (
		projectIdList []string
		projectList   []interface{}
	)

	// 遍历token，转换为hash，封装要存储到es的对象
	// 这里的key是文件名，value是该文件对应的token数组
	for _, k := range projectTokenArrMap.Keys() {
		// 要保存到es的对象
		project := projectToken{
			ProjectName:       projectName,
			Language:          language,
			Suffix:            suffix,
			OpenSourceAddress: openSourceAddress,
			FilePath:          k,
			Tokens:            nil,
		}

		var (
			md5List []string
			buffer  bytes.Buffer // 将token列表按空格拼接
		)

		for _, token := range projectTokenArrMap.Get(k) {
			md5List = append(md5List, GetMd5(token))
			buffer.WriteString(token)
			buffer.WriteString(" ")
		}
		project.TokenContent = buffer.String()
		project.ClearHash = GetMd5(project.TokenContent)
		project.Tokens = md5List
		projectList = append(projectList, &project)
		projectIdList = append(projectIdList, zdpgo_uuid.StringNoLine())
	}

	// 保存到es
	// 创建ES对象
	e, err := zdpgo_es.NewWithConfig(&zdpgo_es.Config{
		Debug:     true,
		Addresses: []string{"https://localhost:9200"},
		Username:  "elastic",
		Password:  "123456",
		CertPath:  "/home/zhangdapeng/dev/es/ca.crt",
	})
	if err != nil {
		panic(err)
	}

	// 添加文档
	indexName := "token_java"

	// 测试阶段先清空索引，防止数据重复
	if _, err = e.DeleteIndex(indexName); err != nil {
		panic(err)
	}

	// 批量添加文档
	response, err := e.AddManyDocument(indexName, projectIdList, projectList)
	fmt.Println(response, err)
}

// GetMd5 获取一个文本的md5值
func GetMd5(text string) string {
	data := []byte(text)
	has := md5.Sum(data)
	md5str1 := fmt.Sprintf("%x", has) //将[]byte转成16进制
	return md5str1
}
