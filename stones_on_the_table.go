package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	_,err := in.ReadString('\n')
	if err != nil{
		return
	}
	var sum int
	str, _ := in.ReadString('\n')
	str = strings.ReplaceAll(str, "\r", "")
	str = strings.ReplaceAll(str, "\n", "")
	for i := 0; i < len(str)-1; i++ {
		j := i + 1
		if str[i] == str[j] {
			sum += 1
		}

	}
	fmt.Println(sum)
}
