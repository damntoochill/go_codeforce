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
	index := strings.Index(str,"h")
	if index == -1{
		fmt.Println("NO")
		return
	}
	str = str[index+1:]
	index = strings.Index(str,"e")
	if index == -1{
		fmt.Println("NO")
		return
	}
	str = str[index+1:]
	index = strings.Index(str,"l")
	if index == -1{
		fmt.Println("NO")
		return
	}
	str = str[index+1:]
	index = strings.Index(str,"l")
	if index == -1{
		fmt.Println("NO")
		return
	}
	str = str[index+1:]
	index = strings.Index(str,"o")
	if index == -1{
		fmt.Println("NO")
		return
	}
	fmt.Println("YES")
}