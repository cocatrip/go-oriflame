package main

import (
	"fmt"

	"github.com/cocatrip/go-oriflame"
)

func main() {
	client := new(oriflame.Client)
	client.Init()

	product, err := client.GetProduct("1276")
	if err != nil {
		fmt.Printf("Error getting product: %v", err)
	}

	fmt.Println(product)
}
