package main

import (
	"fmt"
	"github.com/zhangdapeng520/zdpgo_es"
	"github.com/zhangdapeng520/zdpgo_password"
	"github.com/zhangdapeng520/zdpgo_sim"
	"github.com/zhangdapeng520/zdpgo_uuid"
	"time"
)

type projectToken struct {
	ProjectName       string `json:"project_name"`        // 项目名
	Language          string `json:"language"`            // 编程语言
	Suffix            string `json:"suffix"`              // 文件后缀
	OpenSourceAddress string `json:"open_source_address"` // 开源地址
	ClearHash         string `json:"clear_hash"`          // 清洗后文件hash
	TokenContent      string `json:"token_content"`       // 未hash之前的token按空格拼接
	HashContent       string `json:"hash_content"`        // 将每个token进行hash然后拼接成一个字符串
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
	indexName := "token_project_java"

	// 测试阶段先清空索引，防止数据重复
	if _, err = e.DeleteIndex(indexName); err != nil {
		panic(err)
	}

	// 连接池数量
	// 这里可以给每个文件分配一个Goroutine，还要根据内存来，1G内存可以8W个（需自行测试）
	poolSize := 10000

	// 获取token
	for _, p := range testData {
		fmt.Println("开始计算项目级指纹：", p.ProjectName)
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

		// 获取项目token
		token := zdpgo_sim.GetProjectToken(projectFileInfo)

		// 获取项目hash
		md5Hash := zdpgo_password.GetMd5(token)

		// 获取项目hash内容
		hashContent := zdpgo_sim.GetProjectHash(projectFileInfo)

		// 要保存到es的对象
		projectId := zdpgo_uuid.StringNoLine()
		project := projectToken{
			ProjectName:       p.ProjectName,
			Language:          p.Language,
			Suffix:            p.Suffix,
			OpenSourceAddress: p.OpenSourceAddress,
			ClearHash:         md5Hash,
			TokenContent:      token,
			HashContent:       hashContent,
		}

		// 批量添加文档
		response, err := e.AddDocument(indexName, projectId, &project)
		if response.Status != 201 || err != nil {
			panic("添加文档失败")
		}

		fmt.Println("项目级指纹计算并存入ES8完成：", p.ProjectName)
		fmt.Println("消耗时间：", time.Since(startTime).Milliseconds(), "ms\n")
	}
}
