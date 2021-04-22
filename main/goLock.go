package main

import "sync"

/*
*  @author liqiqiorz
*  @data 2021/4/22 15:41
 */
var mu sync.Mutex

func getInstance() *singleton {
	mu.Lock() // 如果实例存在没有必要加锁
	defer mu.Unlock()

	if instance == nil {
		instance = &singleton{}
	}
	return instance
}
