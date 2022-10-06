package template

import (
	"fmt"
	"io"
	"os"

	"github.com/valyala/fasttemplate"
)

// package mecca provides a .mec template parser. Mec templates are
// like handlebar templates.

type Parser struct {
	state  map[string]interface{}
	output io.Writer
}

func NewParser(state map[string]interface{}, output io.Writer) Parser {
	return Parser{
		state:  state,
		output: output,
	}
}

func (p Parser) Parse(name string) error {
	templateFile, err := os.ReadFile(name)
	if err != nil {
		return fmt.Errorf("parsing %v failed: %w", name, err)
	}

	t := fasttemplate.New(string(templateFile), "[", "]")
	_, err = t.Execute(p.output, p.state)
	if err != nil {
		return fmt.Errorf("parsing %v failed: %w", name, err)
	}

	return nil
}
