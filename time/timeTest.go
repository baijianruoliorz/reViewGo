package main

import (
	"fmt"
	"time"
)

func timeDemo(){
	now := time.Now() //获取当前时间
	fmt.Printf("current time:%v\n", now)
	//current time:2021-07-20 15:04:53.152075 +0800 CST m=+0.000086062
	year := now.Year()     //年
	month := now.Month()   //月
	day := now.Day()       //日
	hour := now.Hour()     //小时
	minute := now.Minute() //分钟
	second := now.Second() //秒
	//2021-07-20 15:04:53
	fmt.Printf("%d-%02d-%02d %02d:%02d:%02d\n", year, month, day, hour, minute, second)
}
//获取时间戳的代码
func timestampDemo() {
	now := time.Now()            //获取当前时间
	timestamp1 := now.Unix()     //时间戳
	timestamp2 := now.UnixNano() //纳秒时间戳
	fmt.Printf("current timestamp1:%v\n", timestamp1)
	fmt.Printf("current timestamp2:%v\n", timestamp2)
}
//使用time.unix可以使时间戳转为时间格式
func timestampDemo2(timestamp int64) {
	timeObj := time.Unix(timestamp, 0) //将时间戳转为时间格式
	fmt.Println(timeObj)
	year := timeObj.Year()     //年
	month := timeObj.Month()   //月
	day := timeObj.Day()       //日
	hour := timeObj.Hour()     //小时
	minute := timeObj.Minute() //分钟
	second := timeObj.Second() //秒
	fmt.Printf("%d-%02d-%02d %02d:%02d:%02d\n", year, month, day, hour, minute, second)
}

//时间间隔
//const (
//	Nanosecond  Duration = 1
//	Microsecond          = 1000 * Nanosecond
//	Millisecond          = 1000 * Microsecond
//	Second               = 1000 * Millisecond
//	Minute               = 60 * Second
//	Hour                 = 60 * Minute
//)
//duration表示的是一段时间间隔，代表两个时间点之间经过的时间，以纳秒为单位
//可表示的最长时间段大约为290年
func main(){
	timeDemo()
}
