package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"

	"github.com/c-bata/go-prompt"
)

const (
	Placeholder       = "{}"
	ExpandPlaceholder = "{...}"
)

func main() {
	baseCommand := appendPlaceholderIfNeeded(os.Args[1:])

	stdin, err := readStdin()
	if err != nil {
		panic(err)
	}

	pt := initPrompt()

	for {
		t := pt.Input()
		if t == "" {
			break
		}

		if err := eval(stdin, baseCommand, t); err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
		}
	}
}

func appendPlaceholderIfNeeded(baseCommand []string) []string {
	for _, word := range baseCommand {
		if word == Placeholder || word == ExpandPlaceholder {
			return baseCommand
		}
	}

	return append(baseCommand, Placeholder)
}

func readStdin() ([]byte, error) {
	var stdin []byte
	fi, _ := os.Stdin.Stat()

	if (fi.Mode() & os.ModeCharDevice) == 0 {
		bytes, err := ioutil.ReadAll(os.Stdin)
		if err != nil {
			return nil, err
		}
		stdin = bytes
	}

	return stdin, nil
}

func initPrompt() *prompt.Prompt {
	return prompt.New(
		func(in string) {},
		func(d prompt.Document) []prompt.Suggest {
			s := []prompt.Suggest{}
			return prompt.FilterHasPrefix(s, d.GetWordBeforeCursor(), true)
		},
		prompt.OptionPrefix("> "),
		prompt.OptionAddASCIICodeBind(prompt.ASCIICodeBind{
			ASCIICode: []byte{0x1b, 0x62},
			Fn:        prompt.GoLeftWord,
		}),
		prompt.OptionAddASCIICodeBind(prompt.ASCIICodeBind{
			ASCIICode: []byte{0x1b, 0x66},
			Fn:        prompt.GoRightWord,
		}),
		prompt.OptionAddASCIICodeBind(prompt.ASCIICodeBind{
			ASCIICode: []byte{0x1b, 0x64},
			Fn:        prompt.DeleteWord,
		}),
	)
}

func eval(stdin []byte, baseCommand []string, input string) error {
	name := baseCommand[0]
	args := []string{}

	for _, c := range baseCommand[1:] {
		if c == Placeholder {
			args = append(args, input)
		} else if c == ExpandPlaceholder {
			args = append(args, strings.Split(input, " ")...)
		} else {
			args = append(args, c)
		}
	}

	cmd := exec.Command(name, args...)
	cmd.Stdin = bytes.NewBuffer(stdin)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}
