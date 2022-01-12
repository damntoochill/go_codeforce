package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func readInt(in *bufio.Reader) int {
	str,_ := in.ReadString('\n')
	str = strings.ReplaceAll(str,"\n","")
	str = strings.ReplaceAll(str,"\r","")
	n,_ := strconv.Atoi(str)
	//fmt.Println(n)
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
	//fmt.Println(tc)
	for i:=0;i<tc;i++{
		//fmt.Println(i)
		arr := readIntArray(in)
		a := float64(arr[0])
		b := float64(arr[1])
		c := float64(arr[2])
		//fmt.Println(a,b,c)
		x := (a+c)/2
		y := 2*b - a
		z := 2*b - c
		//fmt.Println(x,y,z)
		if math.Mod(x,b) ==0 && x>0|| math.Mod(y,c) == 0&& y>0 || math.Mod(z,a) ==0 && z>0{
			fmt.Println("YES")
		}else{
			fmt.Println("NO")
		}
	}
}