package main

import (
	"fmt"

	"github.com/y103kim/treedo/database"
)

func main() {
	database := &database.Database{}
	if err := database.Open(); err != nil {
		fmt.Println(err)
	}
}
