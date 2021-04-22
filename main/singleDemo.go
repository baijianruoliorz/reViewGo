package main

import "sync"

/*
*  @author liqiqiorz
*  @data 2021/4/22 15:13
 */

type singleton struct {
}

var instance *singleton

//我们希望利用Go惯用的方式来实现这个单例模式。我们在标准库sync中找到了Once类型。
var once sync.Once

func GetInstance() *singleton {
	once.Do(func() {
		instance = &singleton{}
	})
	return instance
}
