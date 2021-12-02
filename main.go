package main

import (
	"io"
	"os"
	"os/exec"
	"strings"

	"github.com/peterh/liner"
)

const (
	Placeholder = "{}"
)

func main() {
	baseCommand := appendPlaceholderIfNeeded(os.Args[1:])

	line := initLinerState()
	defer line.Close()

	for {
		input, canceled, aborted, err := read(line)
		if err != nil {
			panic(err)
		}
		if canceled {
			os.Exit(0)
		}
		if aborted {
			continue
		}

		if err := eval(baseCommand, input); err != nil {
			panic(err)
		}
	}
}

func appendPlaceholderIfNeeded(baseCommand []string) []string {
	if strings.Contains(strings.Join(baseCommand, " "), Placeholder) {
		return baseCommand
	}

	return append(baseCommand, Placeholder)
}

func initLinerState() *liner.State {
	line := liner.NewLiner()

	line.SetCtrlCAborts(true)

	return line
}

func read(line *liner.State) (input string, canceled bool, aborted bool, err error) {
	if input, err = line.Prompt("> "); err == nil {
		line.AppendHistory((input))

		return input, false, false, nil
	} else if err == io.EOF {
		return "", true, false, nil
	} else if err == liner.ErrPromptAborted {
		return "", false, true, nil
	} else {
		return "", false, false, err
	}
}

func eval(baseCommand []string, input string) error {
	name := baseCommand[0]
	args := []string{}

	words := strings.Split(input, " ")
	for _, c := range baseCommand[1:] {
		if c == Placeholder {
			args = append(args, words...)
		} else {
			args = append(args, c)
		}
	}

	cmd := exec.Command(name, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}
