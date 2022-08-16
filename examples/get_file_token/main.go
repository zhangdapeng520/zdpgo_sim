package main

import (
	"fmt"
	"github.com/zhangdapeng520/zdpgo_sim"
)

func main() {
	filePath := "/data/tmp/RuoYi/ruoyi-admin/src/main/java/com/ruoyi/web/controller/common/CommonController.java"

	// 获取分割前token
	token, err := zdpgo_sim.GetFileToken(filePath)
	if err != nil {
		panic(err)
	}
	fmt.Println(token)

	// 获取按换行符分割后的token
	token, err = zdpgo_sim.GetFileTokenFromArr(filePath)
	if err != nil {
		panic(err)
	}
	fmt.Println(token)
}
