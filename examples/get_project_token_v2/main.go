package main

import (
	"fmt"
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

	// 查看
	fmt.Println(token)
}
