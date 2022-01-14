package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main(){
	lucky_numbers := []int{4,7,44,47,74,77,444,447,474,477,744,747,774,777}
	in := bufio.NewReader(os.Stdin)
	str,_ := in.ReadString('\n')
	str = strings.ReplaceAll(str,"\n","")
	str = strings.ReplaceAll(str,"\r","")
	n,_ := strconv.Atoi(str)
	for i:= range lucky_numbers{
		if n%lucky_numbers[i] == 0{
			fmt.Println("YES")
			return
		}
	}
	fmt.Println("NO")
}