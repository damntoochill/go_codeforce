package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
)

func readInt(in *bufio.Reader) int {
	str,_ := in.ReadString('\n')
	str = strings.ReplaceAll(str,"\n","")
	str = strings.ReplaceAll(str,"\r","")
	n,_ := strconv.Atoi(str)
	return n
}

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
	tc := readInt(in)

	for tc>0{
		n := readInt(in)
		arr := readIntArray(in)
		smallest := arr[0]
		largest := arr[0]
		for i:=0;i<n;i++{
			if arr[i] < smallest{
				smallest = arr[i]
			}
			if arr[i] > largest{
				largest = arr[i]
			}
		}
		fmt.Println(largest-smallest)
		tc--
	}
}