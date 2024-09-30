package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("welcome malhar")
	fmt.Println("please rate our app in 1 to 5")

	reader := bufio.NewReader(os.Stdin)
	input,_ := reader.ReadString('\n')
	fmt.Println("thans for rating ,",input)
	
	numRating ,err := strconv.ParseFloat(strings.TrimSpace(input),64)

	if err !=nil{
		fmt.Println(err)

	}else{
		fmt.Println("Added 1 to your ratinfL",numRating+1)
	}
}