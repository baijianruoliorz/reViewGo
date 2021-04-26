package main

/*
*  @author liqiqiorz
*  @data 2021/4/26 11:31
 */

//漏桶法限流很好理解，假设我们有一个水桶按固定的速率向下方滴落一滴水，无论有多少请求，
//请求的速率有多大，都按照固定的速率流出，对应到系统中就是按照固定的速率处理请求。

//关于漏桶的实现，uber团队有一个开源的github.com/uber-go/ratelimit库。
//这个库的使用方法比较简单，Take() 方法会返回漏桶下一次滴水的时间。

import (
	"fmt"
	"time"

	"go.uber.org/ratelimit"
)

func main() {
	rl := ratelimit.New(100) // per second

	prev := time.Now()
	for i := 0; i < 10; i++ {
		now := rl.Take()
		fmt.Println(i, now.Sub(prev))
		prev = now
	}
}
