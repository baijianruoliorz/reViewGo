package main

import (
	"encoding/json"
	"fmt"
)

/*
*  @author liqiqiorz
*  @data 2021/4/26 10:14
 */

//json tag指定字段名
type person struct {
	//指定json序列化名,会使用小写的name
	Name string `json:"name"`
	//指定序列化忽略此字段
	Age    int64 `json:"-"`
	Weight float64
}

//忽略空值字段
//当 struct 中的字段没有值时， json.Marshal() 序列化的时候不会忽略这些字段，而是默认输出字段的类型
//零值（例如int和float类型零值是 0，
//string类型零值是""，对象类型零值是 nil）。如果想要在序列序列化时忽略这些没有值的字段时，可以在对应字段添
//加omitempty tag。
//ex:
// 在tag中添加omitempty忽略空值
// 注意这里 hobby,omitempty 合起来是json tag值，中间用英文逗号分隔
type User struct {
	Name  string   `json:"name"`
	Email string   `json:"email,omitempty"`
	Hobby []string `json:"hobby,omitempty"`
}

//str:{"name":"yxr","email":"","hobby":null}
//转变后:str:{"name":"yxr"}

//优雅处理字符串格式的数字
//有时候，前端在传递来的json数据中可能会使用字符串类型的数字，这个时候可以在结构体tag中添加string
//来告诉json包从字符串中解析相应字段的数据：
type Card struct {
	ID    int64   `json:"id,string"`    // 添加string tag
	Score float64 `json:"score,string"` // 添加string tag
}

func intAndStringDemo() {
	jsonStr1 := `{"id": "1234567","score": "88.50"}`
	var c1 Card
	if err := json.Unmarshal([]byte(jsonStr1), &c1); err != nil {
		fmt.Printf("json.Unmarsha jsonStr1 failed, err:%v\n", err)
		return
	}
	fmt.Printf("c1:%#v\n", c1) // c1:main.Card{ID:1234567, Score:88.5}
}
