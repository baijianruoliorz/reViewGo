package main

import (
	"encoding/json"
	"fmt"
)

/*
*  @author liqiqiorz
*  @data 2021/4/26 11:00
 */

type UserInfo struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	u1 := UserInfo{Name: "yxr", Age: 18}
	b, _ := json.Marshal(&u1)
	var m map[string]interface{}
	_ = json.Unmarshal(b, &m)
	for k, v := range m {
		fmt.Printf("key:%v value:%v\n", k, v)
	}
}

//输出
//key:name value:yxr
//key:age value:18
//看起来没什么问题，但其实这里是有一个“坑”的。那就是Go语言中的json包在序列化空接口存放的数字类型
//（整型、浮点型等）都会序列化成float64类型。

//第三方库structs
//除了自己实现，Github上也有现成的轮子，例如第三方库：https://github.com/fatih/structs。
//
//它使用的自定义结构体tag是structs:

// UserInfo 用户信息
//type UserInfo struct {
//	Name string `json:"name" structs:"name"`
//	Age  int    `json:"age" structs:"age"`
//}
//用法很简单：
//m3 := structs.Map(&u1)
//for k, v := range m3 {
//fmt.Printf("key:%v value:%v value type:%T\n", k, v, v)
//}
