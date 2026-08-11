package main

import (
	"fmt"
	"landovargas/blog-aggregator/internal/config"
)

func main() {
	err := config.SetUser("Lando")

	if err != nil {
		fmt.Println("Error setting username: ", err)
	}

	cfg, err := config.Read()
	if err != nil {
		fmt.Println("Error reading file: ", err)
	}
	fmt.Println(cfg)
}
