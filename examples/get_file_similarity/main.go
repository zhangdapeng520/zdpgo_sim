package main

import (
	"fmt"
	"github.com/zhangdapeng520/zdpgo_es"
	"github.com/zhangdapeng520/zdpgo_sim"
	"time"
)

func main() {
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

	// 源码文件路径
	filePath := "/data/tmp/RuoYi/ruoyi-admin/src/main/java/com/ruoyi/web/controller/common/CommonController.java"

	startTime := time.Now()
	fmt.Println("开始查询文件级相似度：", filePath)

	// 获取分割前token
	token, err := zdpgo_sim.GetFileToken(filePath)
	if err != nil {
		panic(err)
	}

	// 从项目级token里面匹配
	// 获取按换行符分割后的token
	token, err = zdpgo_sim.GetFileTokenFromArr(filePath)
	if err != nil {
		panic(err)
	}

	// 查询源码级token
	indexName := "token_java"
	resp, err := e.SearchDocument(indexName, zdpgo_es.SearchRequest{
		Source: zdpgo_es.GetMap("excludes", []string{"token_content", "tokens"}),
		Query: &zdpgo_es.Query{
			MatchPhrase: nil,
			Match: zdpgo_es.GetMap("token_content", zdpgo_es.GetMap(
				"query", token,
				"fuzziness", "auto",
				"minimum_should_match", "90%",
			)),
		},
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("查询文件级相似度结束，消耗时间：", time.Since(startTime).Milliseconds(), "ms")
	fmt.Println("查询到以下数据可能相似：")
	for _, source := range resp.Hits.Hits {
		fmt.Println(source.Source["project_name"], source.Source["file_path"])
	}
}
