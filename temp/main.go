package main

import (
	"fmt"

	"github.com/Henry-Sarabia/tinker"
)

func main() {
	t := tinker.New()
	i := t.Item()
	fmt.Println("TEXT: ", i.Text)
	// fmt.Println("ITEM")
	// fmt.Println(i)
}
