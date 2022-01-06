package main

import (
	"bufio"
	"fmt"
	"math"
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

func main(){
	in := bufio.NewReader(os.Stdin)
	tc := readArrInt64(in)
	m := float64(tc[0])
	n := float64(tc[1])

	a := math.Floor((m*n)/2)

	fmt.Println(a)
}