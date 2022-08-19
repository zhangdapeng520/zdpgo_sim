package main

import (
	"encoding/json"
	"fmt"
	"github.com/zhangdapeng520/zdpgo_es"
	"github.com/zhangdapeng520/zdpgo_sim"
	"github.com/zhangdapeng520/zdpgo_type"
	"strings"
	"time"
)

func main() {
	startTime := time.Now()

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

	fmt.Println("开始查询代码片段级相似度：", filePath)

	// 获取分割前token
	token, err := zdpgo_sim.GetFileToken(filePath)
	if err != nil {
		panic(err)
	}
	indexName := "token_java"
	lines := 10

	// 将token数组按照指定精度拆分
	tokenArr := strings.Split(token, " ")
	spreadTokenArr := zdpgo_sim.GetSpreadTokenArr(tokenArr, lines)

	// 构造查询条件
	var intervals []map[string]interface{}
	for _, spreadToken := range spreadTokenArr {
		intervals = append(intervals, zdpgo_type.GetMap("match", zdpgo_type.GetMap(
			"query", spreadToken,
			"max_gaps", 3,
			"ordered", true,
		)))
	}
	request := zdpgo_es.SearchRequest{
		Query: &zdpgo_es.Query{
			Intervals: zdpgo_type.GetMap("token_content",
				zdpgo_type.GetMap("any_of",
					zdpgo_type.GetMap("intervals", intervals))),
		},
	}
	data, _ := json.Marshal(request)
	fmt.Println(string(data))

	// 查询代码片段级token
	resp, err := e.SearchDocument(indexName, request)
	if err != nil {
		panic(err)
	}
	fmt.Println("查询代码片段级相似度结束，消耗时间：", time.Since(startTime).Milliseconds(), "ms")
	fmt.Println("查询到以下文件可能相似：")
	for _, source := range resp.Hits.Hits {
		fmt.Println(source.Source["project_name"], source.Source["file_path"])
	}
}
