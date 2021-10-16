package dsl

type FunctionType struct {
	Type           int
	ParameterCount int
	Function       interface{}
}

var FunctionDefinition map[string]FunctionType

// FunctionLogistic 逻辑函数定义
// 当出现这个的时候，该函数的参数数量必须要与 ParameterCount 一致
const FunctionLogistic = 0

// FunctionNormal 一般函数定义
// 当出现这个的时候， 该函数的参数数量必须要与 ParameterCount 一致
const FunctionNormal = 0

func init() {
	FunctionDefinition["has"] = FunctionType{
		Type:           FunctionLogistic,
		ParameterCount: 1,
		Function: func() {

		},
	}

	FunctionDefinition["and"] = FunctionType{
		Type:           FunctionLogistic,
		ParameterCount: 1,
		Function: func() {

		},
	}

	FunctionDefinition["or"] = FunctionType{
		Type:           FunctionLogistic,
		ParameterCount: 1,
		Function: func() {

		},
	}

	FunctionDefinition["equals"] = FunctionType{
		Type:           FunctionLogistic,
		ParameterCount: 1,
		Function: func() {

		},
	}

	FunctionDefinition["greaterThan"] = FunctionType{
		Type:           FunctionLogistic,
		ParameterCount: 1,
		Function: func() {

		},
	}

	FunctionDefinition["greaterOrEqual"] = FunctionType{
		Type:           FunctionLogistic,
		ParameterCount: 1,
		Function: func() {

		},
	}

	FunctionDefinition["lessThan"] = FunctionType{
		Type:           FunctionLogistic,
		ParameterCount: 1,
		Function: func() {

		},
	}

	FunctionDefinition["lessOrEqual"] = FunctionType{
		Type:           FunctionLogistic,
		ParameterCount: 1,
		Function: func() {

		},
	}

	FunctionDefinition["contains"] = FunctionType{
		Type:           FunctionNormal,
		ParameterCount: 1,
		Function: func() {

		},
	}

	FunctionDefinition["startsWith"] = FunctionType{
		Type:           FunctionNormal,
		ParameterCount: 1,
		Function: func() {

		},
	}

	FunctionDefinition["endsWith"] = FunctionType{
		Type:           FunctionNormal,
		ParameterCount: 1,
		Function: func() {

		},
	}

	FunctionDefinition["any"] = FunctionType{
		Type:           FunctionNormal,
		ParameterCount: 1,
		Function: func() {

		},
	}

	FunctionDefinition["has"] = FunctionType{
		Type:           FunctionNormal,
		ParameterCount: 1,
		Function: func() {

		},
	}
}
