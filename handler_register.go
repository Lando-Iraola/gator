package main

import (
	"context"
	"fmt"
	"time"

	"landovargas/blog-aggregator/internal/database"

	"github.com/google/uuid"
)

func handlerRegister(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <name>", cmd.Name)
	}
	name := cmd.Args[0]

	var newUser database.CreateUserParams
	var id uuid.NullUUID

	id.UUID = uuid.New()
	id.Valid = true

	newUser.ID = id
	newUser.CreatedAt = time.Now()
	newUser.UpdatedAt = time.Now()
	newUser.Name = name
	user, err := s.db.CreateUser(context.Background(), newUser)
	if err != nil {
		return fmt.Errorf("couldn't set current user: %w", err)
	}

	s.cfg.SetUser(user.Name)

	fmt.Println("User created successfully!")
	fmt.Println(user)
	return nil
}
