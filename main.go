package main

import (
	controllers "IAM/Controllers"
	"fmt"
)

func main() {
	err := controllers.StartServer()
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Scanln()
	//ch0 := make(chan string)

}
