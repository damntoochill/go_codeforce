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
	str2,_ := in.ReadString('\n')
	str2 = strings.ReplaceAll(str2,"\n","")
	str2 = strings.ReplaceAll(str2,"\r","")
	for i,j:=0, len(str2)-1; i<len(str)&&j>=0; i,j = i+1,j-1{
		if str[i] != str2[j]{
			fmt.Println("NO")
			return
		}
	}
	fmt.Println("YES")
}