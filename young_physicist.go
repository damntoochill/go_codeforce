package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
)

func convertToInt(arr []string) []int{
	intarray := make([]int,0)
	for i := range arr{
		y,_ := strconv.Atoi(arr[i])
		intarray = append(intarray,y)
	}
	return intarray
}
func main(){
	in := bufio.NewReader(os.Stdin)
	str,_ := in.ReadString('\n')
	str = strings.ReplaceAll(str, "\r","")
	str = strings.ReplaceAll(str, "\n", "")
	n, _ := strconv.Atoi(str)
	var sumx, sumy, sumz int
	for i := 0; i < n; i++{
		str1,_ := in.ReadString('\n')
		str1 = strings.ReplaceAll(str1, "\r","")
		str1 = strings.ReplaceAll(str1, "\n", "")
		strarray := strings.Split(str1, " ")
	    intarray := convertToInt(strarray)
		sumx = sumx + intarray[0]
		sumy = sumy + intarray[1]
		sumz = sumz + intarray[2]
	}
	if sumx == 0 && sumy ==0 && sumz ==0{
		fmt.Println("YES")
	} else{
		fmt.Println("NO")
	}
	
}