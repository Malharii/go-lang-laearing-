package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("welcome to time study of goLang ")

	presenttime := 	time.Now()
	fmt.Println(presenttime)
  fmt.Println(presenttime.Format("01-02-2006 Monday"))
}