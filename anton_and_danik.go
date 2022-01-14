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

func main(){
	in := bufio.NewReader(os.Stdin)
	n := readInt(in)
	str,_ := in.ReadString('\n')
	str = strings.ReplaceAll(str,"\n","")
	str = strings.ReplaceAll(str,"\r","")
	var a,d int
	for i := 0; i<n;i++{
		if string(str[i]) == "A"{
			a += 1
		}else{
			d += 1
		}
	}
	if a>d{
		fmt.Println("Anton")
	}else if a == d{
		fmt.Println("Friendship")
	}else{
		fmt.Println("Danik")
	}
}