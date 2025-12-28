package main

type command struct {
	name string
	args []string
}

func NewCommand(name string, args []string) command {
	return command{
		name: name,
		args: args,
	}
}
