package main

import (
	"fmt"
	"go-academy/utils"
)

func main() {
	for i := 0; i <= 50; i++ {
		eyesUp, outcome := utils.RollDicePair()
		fmt.Println(eyesUp, outcome)
	}
}
