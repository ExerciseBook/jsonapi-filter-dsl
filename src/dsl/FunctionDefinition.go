package dsl

type FunctionType struct {
	Type           int
	ParameterCount int
	Function       interface{}
}

var FunctionDefinition = make(map[string]FunctionType)

// FunctionLogistic 逻辑函数定义
// 当出现这个的时候，该函数的参数数量必须要与 ParameterCount 一致
const FunctionLogistic = 0

// FunctionNormal 一般函数定义
// 当出现这个的时候， 该函数的参数数量必须要与 ParameterCount 一致
const FunctionNormal = 0

func init() {
	// 啥都没有
}
