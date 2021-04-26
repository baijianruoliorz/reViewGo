package main

import (
	"encoding/json"
	"fmt"
)

/*
*  @author liqiqiorz
*  @data 2021/4/26 9:57
 */

type Person struct {
	Name   string
	Age    int64
	Weight float64
}

func main() {
	p1 := Person{
		Name:   "yxr",
		Age:    18,
		Weight: 70,
	}
	//struct->json string 序列化
	b, err := json.Marshal(p1)
	if err != nil {
		fmt.Printf("json.Marshal failed,err:%v\n", err)
		return
	}
	fmt.Printf("str:%s\n", b)
	//json string->struct  反序列化

}
