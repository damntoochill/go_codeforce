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
	m := make(map[rune]int)
	for _, char := range str {
		m[char] += 1
	}
	if len(m)%2 == 0 {
		fmt.Println("CHAT WITH HER!")
	} else {
		fmt.Println("IGNORE HIM!")
	}
}
