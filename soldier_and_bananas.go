package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readIntArray (in *bufio.Reader) []int{
	str,_ := in.ReadString('\n')
	str = strings.ReplaceAll(str,"\n","")
	str = strings.ReplaceAll(str,"\r","")
	strarr := strings.Split(str," ")
	intarr := make([]int,0)
	for i := range strarr{
		val,_ := strconv.Atoi(strarr[i])
		intarr = append(intarr,val )
	}
	return intarr
}
func main(){
	in := bufio.NewReader(os.Stdin)
	arr := readIntArray(in)
	k := arr[0]
	n := arr[1]
	w := arr[2]
	var total int
	for j:=1;j<=w;j++{
		total = total + j*k
	}
	if total-n >0 {
		fmt.Println(total-n)
	}else{
		fmt.Println("0")
	}
	
}