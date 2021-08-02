package strconvPackage

import (
	"fmt"
	"strconv"
)

func AtoiTest(){
	s1:="100"
	il,err:=strconv.Atoi(s1)
	if err!=nil{
		fmt.Println("can't convert to int")
	}else {
		//type int value 100
		fmt.Printf("type:%T value:%#v",il,il)
	}
}

func ItoaTest(){
	i2:=200
	s2:=strconv.Itoa(i2)
	//type:string value "200"
	fmt.Printf("type:%T value:%#v",s2,s2)
}

//ParseInt()   func ParseInt(s string, base int, bitSize int) (i int64, err error)
//返回字符串表示的整数值，接受正负号。
//base指定进制（2到36），如果base为0，则会从字符串前置判断，”0x”是16进制，”0”是8进制，否则是10进制；
//bitSize指定结果必须能无溢出赋值的整数类型，0、8、16、32、64 分别代表 int、int8、int16、int32、int64；
//返回的err是*NumErr类型的，如果语法有误，err.Error = ErrSyntax；如果结果超出类型范围err.Error = ErrRange。
func ParseIntTest(){
    s3:="10000"
	integer, err := strconv.ParseInt(s3, 10, 32)
	if err!=nil{
		fmt.Printf("type:%T value:%#v",integer,integer)
	}else {
		fmt.Printf("err:%+v",err)
	}
}



