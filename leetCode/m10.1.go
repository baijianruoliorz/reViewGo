package leetCode

import (
	"sort"
	"strings"
)

func groupAnagrams(strs []string)[][]string{
	result:=[][]string{}
	result_tmp:=map[string][]string{}
	for _,item:=range strs{
		tmp:=strings.Split(item,"")
		sort.Strings(tmp)
		key:=strings.Join(tmp,"")
		if _,ok:=result_tmp[strings.Join(tmp,"")];ok{
			result_tmp[key]=append(result_tmp[key],item)
		}else {
			result_tmp[key]=[]string{item}
		}
	}
	for _,v:=range  result_tmp{
		result=append(result,v)
	}
	return result
}