package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readInt(in *bufio.Reader) int {
	str,_ := in.ReadString('\n')
	str = strings.ReplaceAll(str,"\n","")
	str = strings.ReplaceAll(str,"\r","")
	n,_ := strconv.Atoi(str)
	return n
}

func readIntArray (in *bufio.Reader) []int {
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
	n := readInt(in)
	var capacity,minCapacity int
	for n>0{
		arr := readIntArray(in)
		exit := arr [0]
		enter := arr[1]
		capacity = capacity + enter - exit
		if capacity > minCapacity{
			minCapacity = capacity
		}
		n--
	}
	fmt.Println(minCapacity)
}