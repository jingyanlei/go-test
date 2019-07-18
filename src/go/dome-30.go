package main

import (
	"fmt"
	"encoding/json"
)

//30. 不导出的 struct 字段无法被 encode
//以小写字母开头的字段成员是无法被外部直接访问的，所以 struct 在进行 json、xml、gob 等格式的 encode 操作时，这些私有字段会被忽略，导出时得到零值：

type MyData struct {
	One int
	two string
}

func main() {
	in := MyData{1, "two"}
	fmt.Printf("%#v\n", in)    // main.MyData{One:1, two:"two"}

	encoded, _ := json.Marshal(in)
	fmt.Println(string(encoded))    // {"One":1}    // 私有字段 two 被忽略了

	var out MyData
	json.Unmarshal(encoded, &out)
	fmt.Printf("%#v\n", out)     // main.MyData{One:1, two:""}
}
