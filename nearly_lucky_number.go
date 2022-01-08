package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
)

func main(){
	in := bufio.NewReader(os.Stdin)
	str,_ := in.ReadString('\n')
	str = strings.ReplaceAll(str,"\n","")
	str = strings.ReplaceAll(str,"\r","")
	var count int
	for i:=0;i<len(str);i++{
		if string(str[i]) == "4" || string(str[i]) == "7"{
			count += 1
		}
	}

	if count == 7 || count  == 4{
		fmt.Println("YES")
		return
	}
	fmt.Println("NO")
}	