package main

import (
	"fmt"
	"github.com/zhangdapeng520/zdpgo_sim"
)

func main() {
	// 项目目录
	projectDir := "D:\\zdppy\\django_for_api_4.0"

	// 连接池数量
	// 这里可以给每个文件分配一个Goroutine，还要根据内存来，1G内存可以8W个（需自行测试）
	poolSize := 10000

	// 被忽略的文件夹
	ignoreDirs := []string{"venv", ".git", ".idea"}

	// 源码文件后缀
	codeSuffix := ".py"

	// 获取token
	projectTokenMap, err := zdpgo_sim.GetProjectToken(projectDir, poolSize, codeSuffix, ignoreDirs)
	if err != nil {
		fmt.Println("获取项目token失败：", err)
		return
	}

	// 获取两个项目之间的相似度，基于simhash算法
	var algorithmList = []zdpgo_sim.Option{
		zdpgo_sim.SimHash(),
		zdpgo_sim.Cosine(),
		zdpgo_sim.Hamming(),
		zdpgo_sim.Jaro(),
		zdpgo_sim.JaroWinkler(),
		zdpgo_sim.DiceCoefficient(),
	}

	for _, algorithmFunc := range algorithmList {
		similarityMap := zdpgo_sim.GetProjectSimilarity(projectTokenMap, projectTokenMap, poolSize, algorithmFunc)
		for _, k := range similarityMap.Keys() {
			fmt.Println(k)
			fmt.Println(similarityMap.Get(k))
			fmt.Println()
		}
	}
}
