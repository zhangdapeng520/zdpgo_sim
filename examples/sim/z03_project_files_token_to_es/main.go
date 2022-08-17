package main

import (
	"fmt"
	"github.com/zhangdapeng520/zdpgo_es"
	"github.com/zhangdapeng520/zdpgo_sim"
	"github.com/zhangdapeng520/zdpgo_uuid"
	"time"
)

type projectToken struct {
	ProjectName       string `json:"project_name"`        // 项目名
	Language          string `json:"language"`            // 编程语言
	Suffix            string `json:"suffix"`              // 文件后缀
	OpenSourceAddress string `json:"open_source_address"` // 开源地址
	OriginHash        string `json:"origin_hash"`         // 原始文件hash
	ClearCode         string `json:"clear_code"`          // 清洗后代码
	ClearHash         string `json:"clear_hash"`          // 清洗后文件hash
	FilePath          string `json:"file_path"`           // 文件路径
	FileSize          int64  `json:"file_size"`           // 文件大小
	TokenContent      string `json:"token_content"`       // 按空格拼接文件的token数组
	TokenContentHash  string `json:"token_content_hash"`  // token_content的hash值
	HashContent       string `json:"hash_content"`        // 将文件的token数组全部hash，然后按空格拼接
}

func main() {
	testData := []struct {
		ProjectName       string
		OpenSourceAddress string
		ProjectDir        string
		Language          string
		Suffix            string
		IgnoreDirs        []string
	}{
		{
			ProjectName:       "RuoYi",
			OpenSourceAddress: "https://gitee.com/y_project/RuoYi.git",
			ProjectDir:        "/data/tmp/RuoYi",
			Language:          "java",
			Suffix:            ".java",
			IgnoreDirs:        []string{"out", ".mvn", ".git", ".idea"},
		},
		{
			ProjectName:       "hutool",
			OpenSourceAddress: "https://gitee.com/dromara/hutool.git",
			ProjectDir:        "/data/tmp/hutool",
			Language:          "java",
			Suffix:            ".java",
			IgnoreDirs:        []string{"out", ".mvn", ".git", ".idea"},
		},
		{
			ProjectName:       "jeesite4",
			OpenSourceAddress: "https://gitee.com/thinkgem/jeesite4.git",
			ProjectDir:        "/data/tmp/jeesite4",
			Language:          "java",
			Suffix:            ".java",
			IgnoreDirs:        []string{"out", ".mvn", ".git", ".idea"},
		},
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

	// 连接池数量
	// 这里可以给每个文件分配一个Goroutine，还要根据内存来，1G内存可以8W个（需自行测试）
	poolSize := 10000

	// 获取token
	for _, p := range testData {
		fmt.Println("开始计算项目：", p.ProjectName)
		startTime := time.Now()

		projectFileInfo, err := zdpgo_sim.GetProjectFileInfo(
			p.ProjectDir,
			poolSize,
			p.Suffix,
			p.IgnoreDirs)
		if err != nil {
			fmt.Println("获取项目token失败：", err)
			return
		}

		// 计算token和hash
		var (
			projectIdList []string
			projectList   []interface{}
		)

		// 遍历token，转换为hash，封装要存储到es的对象
		// 这里的key是文件名，value是该文件对应的token数组
		for _, k := range projectFileInfo.Keys() {
			// 源码文件信息
			fileInfo := projectFileInfo.Get(k)

			// 要保存到es的对象
			project := projectToken{
				ProjectName:       p.ProjectName,
				Language:          p.Language,
				Suffix:            p.Suffix,
				OpenSourceAddress: p.OpenSourceAddress,
				OriginHash:        fileInfo.OriginHash,
				FilePath:          k,
				FileSize:          fileInfo.FileSize,
				ClearCode:         fileInfo.ClearCode,
				ClearHash:         fileInfo.ClearHash,
				TokenContent:      fileInfo.TokenContent,
				TokenContentHash:  fileInfo.TokenContentHash,
				HashContent:       fileInfo.HashContent,
			}

			projectList = append(projectList, &project)
			projectIdList = append(projectIdList, zdpgo_uuid.StringNoLine())
		}

		// 批量添加文档
		response, err := e.AddManyDocument(indexName, projectIdList, projectList)
		if response.Status != 200 || err != nil {
			panic("添加文档失败")
		}

		fmt.Println("项目计算并存入ES8完成：", p.ProjectName)
		fmt.Println("消耗时间：", time.Since(startTime).Milliseconds(), "ms\n")
	}
}
