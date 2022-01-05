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

func readLineNumbs(in *bufio.Reader) []string {
	line, _ := in.ReadString('\n')
	line = strings.ReplaceAll(line, "\r", "")
	line = strings.ReplaceAll(line, "\n", "")
	numbs := strings.Split(line, " ")
	return numbs
}

func readLine(in *bufio.Reader) string {
	line, _ := in.ReadString('\n')
	line = strings.ReplaceAll(line, "\r", "")
	line = strings.ReplaceAll(line, "\n", "")
	return line
}

func main() {
	in := bufio.NewReader(os.Stdin)
	tc := readInt(in)

	for i := 0; i < tc; i++ {
		n := readLine(in)
		if len(n) > 10 {
			w := fmt.Sprintf("%s%d%s",string(n[0]), len(n)-2, string(n[len(n)-1]))
			fmt.Println(w)
		} else {
			fmt.Println(n)
		}
	}
}
