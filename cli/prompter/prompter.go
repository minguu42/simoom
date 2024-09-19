package prompter

import (
	"bufio"
	"fmt"
	"io"
)

type Prompter struct {
	stdout  io.Writer
	scanner *bufio.Scanner
}

func New(stdin io.Reader, stdout io.Writer) *Prompter {
	return &Prompter{
		stdout:  stdout,
		scanner: bufio.NewScanner(stdin),
	}
}

func (p *Prompter) Input(prompt string) string {
	for {
		fmt.Fprintf(p.stdout, "\x1b[36m%s\x1b[0m ", prompt)
		p.scanner.Scan()
		if in := p.scanner.Text(); in != "" {
			return in
		}
	}
}
