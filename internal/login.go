package internal

import (
	"errors"
	"fmt"
	"landovargas/blog-aggregator/internal/config"
)

func HandlerLogins(s *config.State, cmd config.Command) error {
	if len(cmd.Arguments) != 1 {
		return errors.New("Expected a single argument, username")
	}

	err := s.Config.SetUser(cmd.Arguments[0])
	if err != nil {
		return err
	}

	fmt.Printf("User name: " + cmd.Arguments[0] + " has been set")
	return nil
}
