package main

import (
	"encoding/json"
	"bytes"
	"fmt"
	"log"
)

//37. 将 JSON 中的数字解码为 interface 类型
//在 encode/decode JSON 数据时，Go 默认会将数值当做 float64 处理，比如下边的代码会造成 panic：

func checkError(err error) {
	if err != nil{
		log.Fatalln(err)
	}
}

/*
func main() {
	var data = []byte(`{"status": 200}`)
	var result map[string]interface{}

	if err := json.Unmarshal(data, &result); err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("%T\n", result["status"])    // float64
	var status = result["status"].(int)    // 类型断言错误
	fmt.Println("Status value: ", status)
}
*/

//panic: interface conversion: interface {} is float64, not int
//如果你尝试 decode 的 JSON 字段是整型，你可以：
//
//将 int 值转为 float 统一使用
//将 decode 后需要的 float 值转为 int 使用
// 将 decode 的值转为 int 使用

/*
func main() {
	var data = []byte(`{"status": 200}`)
	var result map[string]interface{}

	if err := json.Unmarshal(data, &result); err != nil {
		log.Fatalln(err)
	}
	var status2 = result["status"]
	fmt.Printf("Status value: %f, %T\n", status2, status2)
	var status = uint64(result["status"].(float64))
	fmt.Printf("Status value: %d, %T", status, status)
}
*/

//使用 Decoder 类型来 decode JSON 数据，明确表示字段的值类型
//// 指定字段类型
/*
func main() {
	var data = []byte(`{"status": 200}`)
	var result map[string]interface{}

	var decoder = json.NewDecoder(bytes.NewReader(data))
	decoder.UseNumber()

	if err := decoder.Decode(&result); err != nil {
		log.Fatalln(err)
	}

	var status2 = result["status"]
	fmt.Printf("Status value: %d, %T\n", status2, status2)
	var status, err = result["status"].(json.Number).Int64()
	fmt.Println("-----", err)
	fmt.Printf("Status value: %d, %T", status, status)
}
*/

//
//// 你可以使用 string 来存储数值数据，在 decode 时再决定按 int 还是 float 使用
//// 将数据转为 decode 为 string
/*
func main() {
	var data = []byte(`{"status": 200}`)
	var result map[string]interface{}
	var decoder = json.NewDecoder(bytes.NewReader(data))
	decoder.UseNumber()
	if err := decoder.Decode(&result); err != nil {
		log.Fatalln(err)
	}
	var status uint64
	err := json.Unmarshal([]byte(result["status"].(json.Number).String()), &status);
	checkError(err)
	fmt.Printf("Status value:%s, %T\n", result["status"].(json.Number).String(), result["status"].(json.Number).String())
	fmt.Printf("Status value:%d, %T ", status, status)
}
*/

//​- 使用 struct 类型将你需要的数据映射为数值型
//
//// struct 中指定字段类型
/*
func main() {
	var data = []byte(`{"status": 200}`)
	var result struct {
		Status uint64 `json:"status"`
	}

	err := json.NewDecoder(bytes.NewReader(data)).Decode(&result)
	checkError(err)
	fmt.Printf("Result: %#v, %T", result, result)
}
*/

//可以使用 struct 将数值类型映射为 json.RawMessage 原生数据类型
//适用于如果 JSON 数据不着急 decode 或 JSON 某个字段的值类型不固定等情况：
//
// 状态名称可能是 int 也可能是 string，指定为 json.RawMessage 类型
func main() {
	records := [][]byte{
		[]byte(`{"status":200, "tag":"one"}`),
		[]byte(`{"status":"ok", "tag":"two"}`),
	}

	for idx, record := range records {
		var result struct {
			StatusCode uint64
			StatusName string
			Status     json.RawMessage `json:"status"`
			Tag        string          `json:"tag"`
		}

		err := json.NewDecoder(bytes.NewReader(record)).Decode(&result)
		checkError(err)

		var name string
		err = json.Unmarshal(result.Status, &name)
		if err == nil {
			result.StatusName = name
		}

		var code uint64
		err = json.Unmarshal(result.Status, &code)
		if err == nil {
			result.StatusCode = code
		}

		fmt.Printf("[%v] result => %+v\n", idx, result)
	}
}

