package next_result_test

import (
	"encoding/json"
	"fmt"
	result "github.com/wvkity/next-result"
	"testing"
)

func TestOk(t *testing.T) {
	_result := result.NewMultipleWithEmpty[string, string]()
	_result.Put("user", "admin")
	_result.Put("password", "123456")
	fmt.Println(_result)

	// data测试
	var _data result.Data[map[string]any]
	_data = result.NewMultipleWithOne[string, any]("count", 1)
	fmt.Println(_data)

	_pr := result.NewPlainWithData("Hello world!!!")
	fmt.Println(_pr)
}

func TestFailure(t *testing.T) {
	_mr := result.NewMultipleFailureWithMsg[string, any](500355, "无权限访问，请联系管理员!!!")
	_pr := result.NewPlainFailureWithMsg[string](500458, "Token已失效，请重新登陆!!!")
	fmt.Println(_mr)
	fmt.Println(_mr.Success)
	fmt.Println(_pr)
	fmt.Println(_pr.Success)
}

func TestOKJson(t *testing.T) {
	_result := result.NewMultipleFailureWithEmpty[string, string]()
	_result.Put("user", "admin")
	_result.Put("password", "123456")
	_data, _error := json.Marshal(_result)
	if _error != nil {
		fmt.Printf("序列化失败: %v\n", _error)
	}
	fmt.Println(string(_data))

	_mr := result.NewPlainFailureWithMsg[string](500355, "无权限访问，请联系管理员!!!")
	_md, _err := json.Marshal(_mr)
	if _err != nil {
		fmt.Printf("序列化失败: %v\n", _err)
	}
	fmt.Println(string(_md))
}
