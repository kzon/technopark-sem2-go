package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	input, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}
	result, err := Calculate(string(input))
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(result)
}
