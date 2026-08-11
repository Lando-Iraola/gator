package main

import (
	"fmt"
	"landovargas/blog-aggregator/internal"
	"landovargas/blog-aggregator/internal/config"
	"log"
	"os"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}
	fmt.Printf("Read config: %+v\n", cfg)

	var state config.State
	state.Config = &cfg

	var commands config.Commands
	commands.Commands = make(map[string]func(*config.State, config.Command) error)
	commands.Register("login", internal.HandlerLogins)

	if len(os.Args) < 2 {
		fmt.Println("No arguments were given!")
		os.Exit(1)
	}

	var cmd config.Command
	cmd.Name = os.Args[1]
	cmd.Arguments = os.Args[2:]

	err = commands.Run(&state, cmd)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	os.Exit(0)
}
