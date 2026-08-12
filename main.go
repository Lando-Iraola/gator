package main

import (
	"database/sql"
	"log"
	"os"

	"landovargas/blog-aggregator/internal/config"
	"landovargas/blog-aggregator/internal/database"

	_ "github.com/lib/pq"
)

type state struct {
	cfg *config.Config
	db  *database.Queries
}

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}

	programState := &state{
		cfg: &cfg,
	}

	cmds := commands{
		registeredCommands: make(map[string]func(*state, command) error),
	}
	cmds.register("login", handlerLogin)
	cmds.register("register", handlerRegister)

	if len(os.Args) < 2 {
		log.Fatal("Usage: cli <command> [args...]")
	}

	cmdName := os.Args[1]
	cmdArgs := os.Args[2:]

	db, err := sql.Open("postgres", programState.cfg.DBURL)
	if err != nil {
		log.Fatalf("Failed to open database connection: ")
	}

	dbQueries := database.New(db)

	programState.db = dbQueries
	err = cmds.run(programState, command{Name: cmdName, Args: cmdArgs})
	if err != nil {
		log.Fatal(err)
	}

}
