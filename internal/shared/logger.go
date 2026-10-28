package shared

import (
	"fmt"
	"io"

	"github.com/octovel/crate/internal/cli/models"
)

const (
	Reset  = "\033[0m"
	Gray   = "\033[90m"
	Blue   = "\033[34m"
	Yellow = "\033[33m"
	Green  = "\033[32m"
	Red    = "\033[31m"
)

type Logger struct {
	Stdout io.Writer
	Stderr io.Writer
}

func NewLogger(ctx *models.Context) *Logger {
	return &Logger{
		Stdout: ctx.Stdout,
		Stderr: ctx.Stderr,
	}
}

func (l *Logger) Inf(msg string) {
	fmt.Fprintf(l.Stdout, "%s[%sinf%s]%s %s\n",
		Gray, Blue, Gray,
		Reset,
		msg,
	)
}

func (l *Logger) Suc(msg string) {
	fmt.Fprintf(l.Stdout, "%s[%ssuc%s]%s %s\n",
		Gray, Green, Gray,
		Reset,
		msg,
	)
}

func (l *Logger) Wrn(msg string) {
	fmt.Fprintf(l.Stderr, "%s[%swrn%s]%s %s\n",
		Gray, Yellow, Gray,
		Reset,
		msg,
	)
}

func (l *Logger) Err(msg string) {
	fmt.Fprintf(l.Stderr, "%s[%serr%s]%s %s\n",
		Gray, Red, Gray,
		Reset,
		msg,
	)
}
