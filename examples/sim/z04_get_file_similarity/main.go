package main

import (
	"fmt"
	"github.com/zhangdapeng520/zdpgo_es"
	"github.com/zhangdapeng520/zdpgo_password"
	"github.com/zhangdapeng520/zdpgo_sim"
	"github.com/zhangdapeng520/zdpgo_type"
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

	fmt.Println("开始查询文件级相似度：", filePath)

	// 获取分割前token
	token, err := zdpgo_sim.GetFileToken(filePath)
	if err != nil {
		panic(err)
	}

	indexName := "token_java"

	// 直接查hash
	startTime := time.Now()
	md5Hash := zdpgo_password.GetMd5(token)
	resp, err := e.SearchDocument(indexName, zdpgo_es.SearchRequest{
		Source: zdpgo_type.GetMap("excludes", []string{"token_content", "tokens"}),
		Query: &zdpgo_es.Query{
			MatchPhrase: zdpgo_type.GetMap("token_content_hash", md5Hash),
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

	// 查询源码级token
	startTime = time.Now()
	resp, err = e.SearchDocument(indexName, zdpgo_es.SearchRequest{
		Source: zdpgo_type.GetMap("excludes", []string{"token_content", "tokens"}),
		Query: &zdpgo_es.Query{
			Match: zdpgo_type.GetMap("token_content", zdpgo_type.GetMap(
				"query", token,
				"minimum_should_match", "80%", // 可以通过这个参数调整相似度，这个值越低，匹配的数量越多
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
