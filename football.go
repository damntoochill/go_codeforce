package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	str, _ := in.ReadString('\n')
	str = strings.ReplaceAll(str, "\r", "")
	str = strings.ReplaceAll(str, "\n", "")
	var sum int = 1
	for i := 0; i < len(str)-1; i++ {
		j := i + 1
		if str[j] == str[i] {
			sum += 1
			if sum == 7 {
				fmt.Println("YES")
				return
			}
		}else{
			sum = 1
		}
	}
	fmt.Println("NO")
}
