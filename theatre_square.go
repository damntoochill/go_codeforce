package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	line, _ := in.ReadString('\n')
	line = strings.ReplaceAll(line,"\r","")
	line = strings.ReplaceAll(line,"\n","")
	strarray := strings.Split(line," ")
	m, _ := strconv.ParseFloat(strarray[0],64)
	n, _ := strconv.ParseFloat(strarray[1],64)
	a, _ := strconv.ParseFloat(strarray[2],64)

	var x,y float64
	if a>=n && a>=m{
		fmt.Println("1")
		return
	}
	if a<n && a<m{
		x = math.Ceil(m/a)
		y = math.Ceil(n/a)
	} else if a<=n && a>=m{
		x = 1
		y = math.Ceil(n/a)
	} else {
		x = math.Ceil(m/a)
		y = 1	
	}	

	fmt.Println(strings.TrimRight(strings.TrimRight(fmt.Sprintf("%.2f", x*y), "0"), ".")) 

}