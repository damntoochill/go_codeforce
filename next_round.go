package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readArrInt64(in *bufio.Reader) []int64 {
	line, _ := in.ReadString('\n')
	line = strings.ReplaceAll(line, "\r", "")
	line = strings.ReplaceAll(line, "\n", "")
	numbs := strings.Split(line, " ")
	arr := make([]int64, len(numbs))
	for i, n := range numbs {
		val, _ := strconv.ParseInt(n, 10, 64)
		arr[i] = val
	}
	return arr
}

func main() {
	in := bufio.NewReader(os.Stdin)
	tc := readArrInt64(in)
	//n := tc[0]
	k := tc[1]

	scores := readArrInt64(in)
	cutoff := scores[k-1]
	var sum int
	for _, score := range scores{
		if score>0 && score>=cutoff{
			sum += 1
		}
	}
	fmt.Println(sum)
}