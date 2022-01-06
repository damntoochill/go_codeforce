package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main(){
	in := bufio.NewReader(os.Stdin)
	str, _ := in.ReadString('\n')
	str = strings.ReplaceAll(str,"\r","")
	str = strings.ReplaceAll(str,"\n","")
	strarray := strings.Split(str,"+")
	sort.Strings(strarray)
	str = strings.Join(strarray,"+")
	fmt.Println(str)
}