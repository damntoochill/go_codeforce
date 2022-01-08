package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func getString(in *bufio.Reader) string {
	str,_ := in.ReadString('\n')
	str = strings.ReplaceAll(str,"\r","")
	str = strings.ReplaceAll(str,"\n","")
	return str
}
func main(){
	in := bufio.NewReader(os.Stdin)
	str1 := strings.ToLower(getString(in))
	str2 := strings.ToLower(getString(in))
	if str1 < str2 {
		fmt.Println("-1")
		return
	}else if str1 == str2 {
		fmt.Println("0")
		return
	}else{
		fmt.Println("1")
	}
}