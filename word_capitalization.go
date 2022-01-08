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
	str = strings.ReplaceAll(str, "\r", "")
	str = strings.ReplaceAll(str,"\n","" )
	f := strings.ToUpper(string(str[0]))
	r := string(str[1:])
	fmt.Println(f + r)
}