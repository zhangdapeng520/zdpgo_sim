package main

import (
	"bytes"
	"fmt"
	"github.com/zhangdapeng520/zdpgo_sim"
	"sort"
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

	// 按照文件名排序
	keys := projectTokenMap.Keys()
	sort.Strings(keys)

	// 遍历并生成大的token
	var tokenBuffer bytes.Buffer
	for _, k := range keys {
		tokenBuffer.WriteString(projectTokenMap.Get(k))
		tokenBuffer.WriteString("\n")
	}

	// 查看
	fmt.Println(tokenBuffer.String())
}
