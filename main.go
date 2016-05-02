package main

import (
	"fmt"
)

func main() {
	err := GetAirState()
	if err != nil {
		fmt.Print(err)
	}
}
