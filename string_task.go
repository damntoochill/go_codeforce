package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
)

func findChar (s rune) bool {
	vowels := []rune{ 'a', 'o', 'y', 'e', 'u', 'i'}
	for _,char := range vowels{
		if s == char{
			return true
		}
	}
	return false
}

func main(){
	var result string
	in := bufio.NewReader(os.Stdin)
	str, _ := in.ReadString('\n')
	str = strings.ReplaceAll(str,"\r","")
	str = strings.ReplaceAll(str,"\n","")
	
	str = strings.ToLower(str)
	for _,char := range str{
		if !findChar(char){
			result = result + "." + string(char)
		}
	}
	fmt.Println(result)
}