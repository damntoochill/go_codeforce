package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main(){
	in := bufio.NewReader(os.Stdin)
	str,_ := in.ReadString('\n')
	str = strings.ReplaceAll(str,"\r","")
	str = strings.ReplaceAll(str,"\n","")
	n,_ := strconv.Atoi(str)
	var total int
	for i:=5;i>=1;i--{
		total = total + n/i
		n = n%i
		if n==0{break}
	}
	fmt.Println(total)
}