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
	if len(os.Args) <= 1 {
		panic("Argument is empty.")
	}

	baseCommand := appendPlaceholderIfNeeded(os.Args[1:])

	stdin, err := readStdin()
	if err != nil {
		panic(err)
	}

	pt, err := initPrompt(baseCommand)
	if err != nil {
		panic(err)
	}

	for {
		t := pt.Input()
		if t == "" {
			break
		}

		if err := saveHistory(baseCommand, t); err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
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

func initPrompt(baseCommand []string) (*prompt.Prompt, error) {
	if err := initHistory(baseCommand); err != nil {
		return nil, err
	}
	history, err := readHistory(baseCommand)
	if err != nil {
		return nil, err
	}

	return prompt.New(
		func(in string) {},
		func(d prompt.Document) []prompt.Suggest {
			s := []prompt.Suggest{}
			return prompt.FilterHasPrefix(s, d.GetWordBeforeCursor(), true)
		},
		prompt.OptionPrefix("> "),
		prompt.OptionHistory(history),
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
			Fn: func(buf *prompt.Buffer) {
				buf.Delete(buf.Document().FindEndOfCurrentWordWithSpace())
			},
		}),
	), nil
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
