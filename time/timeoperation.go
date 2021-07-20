package main

import (
	"fmt"
	"time"
)

//add :当遇到时间+时间间隔的需求的时候，go语言的时间对象有提供的add方法如下
//func (t Time) Add(d Duration) Time

//ex:求一个小时之后的时间
//terminal运行的话不需要再进行包设置
//一个小时后
//2021-07-20 16:17:48.581739 +0800 CST m=+3600.000077537
func main(){
	now:=time.Now()
	later:=now.Add(time.Hour)
	fmt.Println(later)
	formatDemo()
	fmt.Println("********")
	analysisTime()
}


// sub 求两个时间之前的差值
//func (t Time) Sub(u Time) Duration
//ex:返回一个时间段t-u。如果结果超出了Duration可以表示的最大值/最小值，
//将返回最大值/最小值。要获取时间点t-d（d为Duration），可以使用t.Add(-d)。

//Equal  比较两个时间是否相同，考虑时区的影响
//判断两个时间是否相同，会考虑时区的影响，因此不同时区标准的时间也可以正确比较。
//本方法和用t==u不同，这种方法还会比较地点和时区信息。

//before
//func (t time) Before(u time)bool
//如果t代表的时间点在u之前，返回真 否则返回假


//After
//func (t time) After(u time) bool
//如果t代表的时间点在u之后，返回真；否则返回假。


//定时器
func tickDemo() {
	ticker := time.Tick(time.Second) //定义一个1秒间隔的定时器
	for i := range ticker {
		fmt.Println(i)//每秒都会执行的任务
	}
}

//时间格式化
//时间类型有一个自带的方法Format进行格式化，
//需要注意的是Go语言中格式化时间模板不是常见的Y-m-d H:M:S
//而是使用Go的诞生时间2006年1月2号15点04分（记忆口诀为2006 1 2 3 4）。
//也许这就是技术人员的浪漫吧。
//补充：如果想格式化为12小时方式，需指定PM。
//2021-07-20 15:48:38.839 Tue Jul
//2021-07-20 03:48:38.839 PM Tue Jul
//2021/07/20 15:48
//15:48 2021/07/20
//2021/07/20

func formatDemo(){
	now:=time.Now()
	// 格式化的模板为Go的出生时间2006年1月2号15点04分 Mon Jan
	// 24小时制
	fmt.Println(now.Format("2006-01-02 15:04:05.000 Mon Jan"))
	// 12小时制
	fmt.Println(now.Format("2006-01-02 03:04:05.000 PM Mon Jan"))
	fmt.Println(now.Format("2006/01/02 15:04"))
	fmt.Println(now.Format("15:04 2006/01/02"))
	fmt.Println(now.Format("2006/01/02"))
}


//解析字符串格式的时间

func analysisTime(){
	now := time.Now()
	fmt.Println(now)
	// 加载时区
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		fmt.Println(err)
		return
	}
	// 按照指定时区和指定格式解析字符串时间
	timeObj, err := time.ParseInLocation("2006/01/02 15:04:05", "2019/08/04 14:15:20", loc)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(timeObj)
	fmt.Println(timeObj.Sub(now))
}

//2021-07-20 15:53:51.375083 +0800 CST m=+0.000385533
//2019-08-04 14:15:20 +0800 CST
//-17185h38m31.375083s


