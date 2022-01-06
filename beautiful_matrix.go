package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)


func main(){
	in := bufio.NewReader(os.Stdin)
	var x,y float64
	for i := 0; i<5; i++{
		str,_ := in.ReadString('\n')
		str = strings.ReplaceAll(str,"\r","")
		str = strings.ReplaceAll(str,"\n","")
		strarray := strings.Split(str," ")
		for k,j := range strarray{
			if j == "1"{
				x = float64(k)
				y = float64(i)
			}
		}
	}
	steps := math.Abs(x-2) + math.Abs(y-2)
	fmt.Println(steps)
}