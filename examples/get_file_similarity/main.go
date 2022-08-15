package main

import (
	"fmt"
	"github.com/zhangdapeng520/zdpgo_es"
	"github.com/zhangdapeng520/zdpgo_sim"
	"strings"
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
	projectTokenMap, err := zdpgo_sim.GetProjectTokenArr(
		projectDir,
		poolSize,
		suffix,
		ignoreDirs)
	if err != nil {
		fmt.Println("获取项目token失败：", err)
		return
	}

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

	indexName := "token_java"

	// 查询文件token
	fileToken := strings.Join(projectTokenMap.Values()[0], " ")
	resp, err := e.SearchDocument(indexName, zdpgo_es.SearchRequest{
		//Source: zdpgo_es.GetMap("excludes", []string{"token_content"}),
		Query: &zdpgo_es.Query{
			Match: zdpgo_es.GetMap("token_content", zdpgo_es.GetMap(
				"query", fileToken,
				"fuzziness", "auto",
				"minimum_should_match", "90%")),
		},
	})
	fmt.Println(fileToken)
	fmt.Println("xxxxxxxxxxxxxxxxxxxx")
	if err != nil {
		panic(err)
	}
	for _, project := range resp.Hits.Hits {
		fmt.Println(project.Source["file_path"])
		pToken := project.Source["token_content"]
		fmt.Println(pToken)
		fmt.Println("================")
	}
}
