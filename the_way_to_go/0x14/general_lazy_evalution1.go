package main

import (
	"fmt"
)

type Any interface{}
type EvalFunc func(Any) (Any, Any)

func main() {
	evenFunc := func(state Any) (Any, Any) { // 从 os 开始，每次 +2
		os := state.(int)
		ns := os + 2
		return os, ns
	}
	even := BuildLazyIntEvaluator(evenFunc, 0) // 初始状态 0

	for i := 0; i < 10; i++ { // 输出前 10 个偶数
		fmt.Printf("%vth even: %v\n", i, even())
	}
}

// 工厂函数
func BuildLazyEvaluator(evalFunc EvalFunc, initState Any) func() Any {
	retValChan := make(chan Any)
	loopFunc := func() {
		var actState Any = initState
		var retVal Any
		for {
			retVal, actState = evalFunc(actState)
			retValChan <- retVal
		}
	}
	retFunc := func() Any {
		return <-retValChan
	}
	go loopFunc()
	return retFunc
}

func BuildLazyIntEvaluator(evalFunc EvalFunc, initState Any) func() int {
	ef := BuildLazyEvaluator(evalFunc, initState)
	return func() int {
		return ef().(int)
	}
}
