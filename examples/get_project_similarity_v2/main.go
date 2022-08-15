package main

import (
	"fmt"
	"github.com/zhangdapeng520/zdpgo_es"
	"github.com/zhangdapeng520/zdpgo_sim"
)

func main() {
	//projectName := "RuoYi"
	//openSourceAddress := "https://gitee.com/y_project/RuoYi.git"
	projectDir := "/data/tmp/RuoYi"
	//language := "java"
	suffix := ".java"
	ignoreDirs := []string{"out", ".mvn", ".git", ".idea"}

	// 连接池数量
	// 这里可以给每个文件分配一个Goroutine，还要根据内存来，1G内存可以8W个（需自行测试）
	poolSize := 10000

	// 获取token
	projectTokenMap, err := zdpgo_sim.GetProjectTokenMap(
		projectDir,
		poolSize,
		suffix,
		ignoreDirs)
	if err != nil {
		fmt.Println("获取项目token失败：", err)
		return
	}

	// 获取项目token
	token := zdpgo_sim.GetProjectToken(projectTokenMap)

	// 获取项目hash
	md5Hash := zdpgo_sim.GetMd5(token)

	// 方案1的相似度，查看ES中是否存在一样的hash
	fmt.Println(md5Hash)

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
	resp, err := e.SearchDocument(indexName, zdpgo_es.SearchRequest{
		Source: zdpgo_es.GetMap("excludes", []string{"token_content"}),
		Query: &zdpgo_es.Query{
			MatchPhrase: zdpgo_es.GetMap("clear_hash", "ef3a6c923294fb2f62a149eda2525dc5"),
		},
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
	for _, project := range resp.Hits.Hits {
		fmt.Println(project.Source)
	}
}
