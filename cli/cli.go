package cli

import (
	"io"
	"errors"
	"log"

	"github.com/chzyer/readline"
	"github.com/google/shlex"
)

var commands = map[string]func(args []string) error {
	"q": quit,
	"e": quit,
	"quit": quit,
	"exit": quit,
}

var errQuit = errors.New("quit")
func quit(args []string) error {
	log.Println("cya next time!")
	return errQuit
}

func Init() (*readline.Instance, error) {
	rl, err := readline.New("albert2> ")
	if err != nil {	panic(err) }
	log.SetOutput(rl.Stdout())
	log.SetFlags(0)

	log.Println("Welcome to albert2!")
	return rl, err
}

func Run(rl *readline.Instance) error {
	for {
		line, err := rl.Readline()
		if err == readline.ErrInterrupt || err == io.EOF {
			quit(nil)
		}

		args, err := shlex.Split(line)
		if err != nil {
			log.Println("parse err:", err)
			// something
		}
		if len(args) == 0 {
			// something
		}

		if cmd, ok := commands[args[0]]; ok {
			err := cmd(args[1:])
			if errors.Is(err, errQuit) { return err}
			if err != nil { log.Println("error:", err)}
		} else {
			log.Println("unknown command:", args[0])
		}
	}
	return nil
}
