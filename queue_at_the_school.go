package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readIntArray(in *bufio.Reader) []int {
	str, _ := in.ReadString('\n')
	str = strings.ReplaceAll(str, "\n", "")
	str = strings.ReplaceAll(str, "\r", "")
	strarr := strings.Split(str, " ")
	intarr := make([]int, 0)
	for i := range strarr {
		val, _ := strconv.Atoi(strarr[i])
		intarr = append(intarr, val)
	}
	return intarr
}

func main() {
	in := bufio.NewReader(os.Stdin)
	arr := readIntArray(in)
	n := arr[0]
	k := arr[1]
	q, _ := in.ReadString('\n')
	q = strings.ReplaceAll(q, "\n", "")
	q = strings.ReplaceAll(q, "\r", "")

	for l := 0; l < k; l++ {
		var result string
		for i := 0; i < n; i++ {
			j := i + 1
			if string(q[i]) == "B" && i != n-1 {
				if string(q[j]) == "G" {
					result += string(q[j]) + string(q[i])
					i += 1
				} else {
					result += string(q[i])
				}
			} else {
				result += string(q[i])
			}
		}
		q = result
	}
	fmt.Println(q)
}
