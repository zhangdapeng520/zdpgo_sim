package main

import (
	"fmt"
	"time"

	"github.com/zhangdapeng520/zdpgo_sim"
)

func main() {
	testData := []struct {
		projectDir string   // 项目目录
		codeSuffix string   // 源码文件后缀
		ignoreDirs []string // 被忽略的文件夹
		showDetail bool     // 查看查看详细内容
	}{
		{"D:\\zdppy\\django_for_api_4.0", ".py", []string{"venv", ".git", ".idea"}, false},
		{"D:\\tmp\\springboot-bucket", ".java", []string{"out", ".mvn", ".git", ".idea"}, false},
		{"D:\\tmp\\thinkphp", ".php", []string{"out", ".mvn", ".git", ".idea"}, false},
		{"D:\\tmp\\FastCFS", ".c", []string{"out", ".mvn", ".git", ".idea"}, false},
		{"D:\\tmp\\QWidgetDemo", ".cpp", []string{"out", ".mvn", ".git", ".idea"}, false},
	}

	// 连接池数量
	// 这里可以给每个文件分配一个Goroutine，还要根据内存来，1G内存可以8W个（需自行测试）
	poolSize := 10000

	fmt.Println("项目名\t后缀\t文件个数\t耗时")
	for _, tt := range testData {
		startTime := time.Now()

		// 获取token
		projectTokenArrMap, err := zdpgo_sim.GetProjectTokenArr(
			tt.projectDir,
			poolSize,
			tt.codeSuffix,
			tt.ignoreDirs)
		if err != nil {
			fmt.Println("获取项目token失败：", err)
			return
		}

		// 遍历
		if tt.showDetail {
			for _, k := range projectTokenArrMap.Keys() {
				fmt.Println(k)
				for _, token := range projectTokenArrMap.Get(k) {
					fmt.Println(token)
				}
				fmt.Println()
			}
		}

		fmt.Printf("%s\t%s\t%d\t%dms\n",
			tt.projectDir,
			tt.codeSuffix,
			projectTokenArrMap.Len(),
			time.Since(startTime).Milliseconds())
	}

}
