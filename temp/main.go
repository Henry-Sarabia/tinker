package main

import (
	"fmt"
	"log"

	"github.com/Henry-Sarabia/tinker"
)

func main() {
	t, err := tinker.New()
	if err != nil {
		log.Fatal(err)
	}
	i, err := t.Item()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(i.Description)
	// fmt.Println("ITEM")
	// fmt.Println(i)
}
