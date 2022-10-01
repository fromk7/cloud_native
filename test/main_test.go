package main

import (
	"testing"
)

func add(a, b int) int {
	return a + b
}

// 文件命名以_test结尾，并通过命令 [go test 测试文件名] 运行单元测试
// go test .
func TestIncrease(t *testing.T) {
	t.Log("start testing")
	result := add(1, 2)

	print(result)
}
