package main

import (
	"fmt"
)

func main() {
  var i int
	_,err := fmt.Scanf("%d",&i)
  if err != nil{
    return
  }
  if i>2 && i%2 == 0{
    fmt.Println("YES")
  }else{
    fmt.Println("NO")
  }
}
