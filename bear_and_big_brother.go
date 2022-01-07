package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main(){
	in := bufio.NewReader(os.Stdin)
	str, _ := in.ReadString('\n')
	str = strings.ReplaceAll(str, "\r", "")
	str = strings.ReplaceAll(str, "\n", "")
	strarray := strings.Split(str, " ")

	a, _ := strconv.ParseFloat(strarray[0], 64)
	b, _ := strconv.ParseFloat(strarray[1], 64)

	var i float64 = 1
	for{
		if math.Pow(3, i)*a > math.Pow(2, i)*b{
			break
		}
		i += 1
	}
	fmt.Println(i)
}
