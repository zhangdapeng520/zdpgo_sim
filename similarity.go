package zdpgo_sim

import (
	"github.com/zhangdapeng520/zdpgo_pool_goroutine"
	"github.com/zhangdapeng520/zdpgo_type/maps/safemap"
)

// GetProjectSimilarity 获取两个项目之间的相似度
func GetProjectSimilarity(project1SafeMap, project2SafeMap *safemap.SafeMap[string, string],
	poolSize int, algorithmFunc Option) *safemap.SafeMap[string, float32] {

	// 默认参数
	if poolSize == 0 {
		poolSize = 33333
	}

	// 结果map
	resultMap := new(safemap.SafeMap[string, float32])

	// 处理方法
	handleFunc := func(key string) {
		// 以项目1为准，分别比较每个文件的相似度
		// 如果项目2中不存在该文件，则相似度为0
		// 如果项目1中存在该文件，则使用指定算法计算相似度
		if project2SafeMap.Has(key) {
			token1 := project1SafeMap.Get(key)
			token2 := project2SafeMap.Get(key)
			simValue := Compare(token1, token2, algorithmFunc)
			resultMap.Set(key, float32(simValue))
		} else {
			resultMap.Set(key, 0)
		}
	}

	// 并行计算
	zdpgo_pool_goroutine.RunBatchArgTask[string](10000, handleFunc, project1SafeMap.Keys())

	// 返回
	return resultMap
}
