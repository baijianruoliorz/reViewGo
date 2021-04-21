package main

import (
	"fmt"
	"sync"
)

/*
*  @author liqiqiorz
*  @data 2021/4/21 21:07
 */
var wg sync.WaitGroup

//go review start! 9
//go review start! 7
//go review start! 8
//go review start! 5
//go review start! 2
//go review start! 0
//go review start! 3
//go review start! 1
//go review start! 4
//go review start! 6
func hello(i int) {
	defer wg.Done()
	fmt.Println("go review start!", i)
}
func main() {
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go hello(i)
	}
	wg.Wait()
}
