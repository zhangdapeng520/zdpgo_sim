package main

import (
	"fmt"
	"github.com/zhangdapeng520/zdpgo_es"
	"github.com/zhangdapeng520/zdpgo_sim"
	"github.com/zhangdapeng520/zdpgo_type"
	"time"
)

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

	indexName := "token_project_java"

	// 连接池数量
	// 这里可以给每个文件分配一个Goroutine，还要根据内存来，1G内存可以8W个（需自行测试）
	poolSize := 10000

	// 获取token
	for _, p := range testData {
		fmt.Println("开始查询项目：", p.ProjectName)
		startTime := time.Now()

		projectTokenMap, err := zdpgo_sim.GetProjectTokenMap(
			p.ProjectDir,
			poolSize,
			p.Suffix,
			p.IgnoreDirs)
		if err != nil {
			fmt.Println("获取项目token失败：", err)
			return
		}

		// 获取项目token
		token := zdpgo_sim.GetProjectToken(projectTokenMap)

		// 获取项目hash
		md5Hash := zdpgo_sim.GetMd5(token)

		// 根据hash查询
		resp, err := e.SearchDocument(indexName, zdpgo_es.SearchRequest{
			Source: zdpgo_type.GetMap("excludes", []string{"token_content"}),
			Query: &zdpgo_es.Query{
				Match: zdpgo_type.GetMap("clear_hash", md5Hash),
			},
		})
		if err != nil {
			panic(err)
		}

		fmt.Println("项目查询完成：", p.ProjectName)
		fmt.Println("ES中相似的项目信息如下：\n", resp.Hits.Hits[0].Source)
		fmt.Println("消耗时间：", time.Since(startTime).Milliseconds(), "ms\n")
	}
}
