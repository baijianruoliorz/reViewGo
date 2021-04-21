package main

import "fmt"

/*
*  @author liqiqiorz
*  @data 2021/4/21 21:31
 */
//对一个关闭的通道再发送值就会导致panic。
//对一个关闭的通道进行接收会一直获取值直到通道为空。
//对一个关闭的并且没有值的通道执行接收操作会得到对应类型的零值。
//关闭一个已经关闭的通道会导致panic。
func recv(c chan int) {
	ret := <-c
	fmt.Println("接受成功", ret)
}
func main() {
	var ch chan int
	fmt.Println(ch)
	//无缓冲的通道

	//ch:=make(chan int)
	//go 协程可以写前面的~
	go recv(ch)
	ch <- 10
	fmt.Println("发送成功")
	//无缓冲通道有一个发送值就必须有一个接受值 否则会出现error
	//fatal error: all goroutines are asleep - deadlock!
	//因为我们使用ch := make(chan int)创建的是无缓冲的通道，无缓冲的通道只有在有人
	//接收值的时候才能发送值。就像你住的小区没有快递柜和代收点，快递员给你打电话
	//必须要把这个物品送到你的手中，简单来说就是无缓冲的通道必须有接收才能发送。
	// 代码会再ch<-10这一行形成死锁
	//可以启动一个goroutine来接受这个值

	//缓冲的通道
	ch1 := make(chan int, 1)
	ch1 <- 10
	fmt.Println("发送成功")
}
