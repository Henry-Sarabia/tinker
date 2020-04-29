package main

import (
	"fmt"

	"github.com/Henry-Sarabia/tinker"
)

func main() {
	t := tinker.New()
	i := t.Item()
	fmt.Println(i.Description)
	// fmt.Println("ITEM")
	// fmt.Println(i)
}
