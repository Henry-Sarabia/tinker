package main

import (
	"fmt"

	"github.com/Henry-Sarabia/tinker"
)

func main() {
	t := tinker.New()
	i := t.Item()
	fmt.Println(i.Text)
	// fmt.Println("ITEM")
	// fmt.Println(i)
}
