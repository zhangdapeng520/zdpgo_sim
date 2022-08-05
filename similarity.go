package zdpgo_sim

import (
	"github.com/zhangdapeng520/zdpgo_pool_goroutine"
	"github.com/zhangdapeng520/zdpgo_type/maps/safemap"
)

type tokenPaire []string

func (item tokenPaire) compare(other tokenPaire) int {
	return 1
}

// ArrSimilarity 数组相似度
type ArrSimilarity struct {
	Token1     string  `json:"token1,omitempty"`
	Token2     string  `json:"token2,omitempty"`
	Similarity float32 `json:"similarity,omitempty"`
}

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
	zdpgo_pool_goroutine.RunBatchArgTask[string](
		poolSize,
		handleFunc,
		project1SafeMap.Keys())

	// 返回
	return resultMap
}

// GetArrSimilarity 获取两个token数组的相似度
func GetArrSimilarity(
	token1Arr, token2Arr []string,
	poolSize int,
	algorithmFunc Option) []ArrSimilarity {

	// 默认参数
	if poolSize == 0 {
		poolSize = 33333
	}

	// 结果
	var smap = new(safemap.SafeMap[string, ArrSimilarity])
	var result []ArrSimilarity

	// 处理方法
	// 参数：[[token1,token2],[token1,token2],...]
	handleFunc := func(tokensArr []string) {
		token1 := tokensArr[0]
		token2 := tokensArr[1]
		similarity := Compare(token1, token2, algorithmFunc)
		obj := ArrSimilarity{
			Token1:     token1,
			Token2:     token2,
			Similarity: float32(similarity),
		}
		smap.Set(token1+token2, obj)
	}

	// 构造参数
	var args [][]string

	// 以token1为准
	for _, token1 := range token1Arr {
		for _, token2 := range token2Arr {
			args = append(args, []string{token1, token2})
		}
	}

	// 并行计算
	zdpgo_pool_goroutine.RunBatchArgTask[[]string](
		poolSize,
		handleFunc,
		args)

	// 返回
	result = smap.Values()
	return result
}
