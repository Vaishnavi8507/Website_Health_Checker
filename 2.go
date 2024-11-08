package main
import("strconv"
"strings")
func ex(input string)[]int  {
	numb:=strings.Split(input,",")
	length:=len(numb)
	var num=make([]int,length)
	for index,v:=range numb{
		s:=strings.Trim(v,"")
		num[index],_=strconv.Atoi(s)

	}
	return num
}