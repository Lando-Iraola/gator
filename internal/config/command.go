package config

import "errors"

type Command struct {
	Name      string
	Arguments []string
}

type Commands struct {
	Commands map[string]func(*State, Command) error
}

func (c *Commands) Run(s *State, cmd Command) error {
	value, ok := c.Commands[cmd.Name]
	if !ok {
		return errors.New("Command " + cmd.Name + " not found")
	}

	err := value(s, cmd)
	if err != nil {
		return err
	}

	return nil
}

func (c *Commands) Register(name string, f func(*State, Command) error) {
	c.Commands[name] = f
}
