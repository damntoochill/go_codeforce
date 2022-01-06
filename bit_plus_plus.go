package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readInt(in *bufio.Reader) int {
	nStr, _ := in.ReadString('\n')
	nStr = strings.ReplaceAll(nStr, "\r", "")
	nStr = strings.ReplaceAll(nStr, "\n", "")
	n, _ := strconv.Atoi(nStr)
	return n
}

func main() {
	in := bufio.NewReader(os.Stdin)
	tc := readInt(in)
	var sum int
	for i := 0; i < tc; i++ {
		str, _ := in.ReadString('\n')
		str = strings.ReplaceAll(str, "\r", "")
		str = strings.ReplaceAll(str, "\n", "")
		if strings.Contains(str, "+") {
			sum += 1
		} else {
			sum -= 1
		}
	}
	fmt.Println(sum)
}
