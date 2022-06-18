package main

import(
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
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

func readInt(in *bufio.Reader) int {
		str,_ := in.ReadString('\n')
		str = strings.ReplaceAll(str,"\n","")
		str = strings.ReplaceAll(str,"\r","")
		n,_ := strconv.Atoi(str)
		return n
	}

	func absInt(x int) int {
		return absDiffInt(x, 0)
	 }
	 
	 func absDiffInt(x, y int) int {
		if x < y {
		   return y - x
		}
		return x - y
	 }

func main(){
	in := bufio.NewReader(os.Stdin)
	t := readInt(in)
	for i := 0; i<t;i++{
		arr := readIntArray(in)
		m := arr[1]
		requiredEnergy := 0
		distances := readIntArray(in)
		for i := range(distances){
			requiredEnergy += distances[i]
		}
		if m>=requiredEnergy{
			fmt.Println("0")
		}else{
			fmt.Println(requiredEnergy-m)
		}
	}
	
}