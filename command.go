package actionlint

import (
	"io"
)

// These variables might be modified by ldflags on building release binaries by GoReleaser. Do not modify manually
var (
	version       = ""
	installedFrom = "installed by building from source"
)

const (
	// ExitStatusSuccessNoProblem is the exit status when the command ran successfully with no problem found.
	ExitStatusSuccessNoProblem = 0
	// ExitStatusSuccessProblemFound is the exit status when the command ran successfully with some problem found.
	ExitStatusSuccessProblemFound = 1
	// ExitStatusInvalidCommandOption is the exit status when parsing command line options failed.
	ExitStatusInvalidCommandOption = 2
	// ExitStatusFailure is the exit status when the command stopped due to some fatal error while checking workflows.
	ExitStatusFailure = 3
)

func printUsageHeader(out io.Writer) { _ = "STUB: not implemented"; return }

func getCommandVersion() string { _ = "STUB: not implemented"; return "" }

// Reaches only when actionlint package is built outside module

// Command represents entire actionlint command. Given stdin/stdout/stderr are used for input/output.
type Command struct {
	// Stdin is a reader to read input from stdin
	Stdin io.Reader
	// Stdout is a writer to write output to stdout
	Stdout io.Writer
	// Stderr is a writer to write output to stderr
	Stderr io.Writer
}

func (cmd *Command) runLinter(args []string, opts *LinterOptions, initConfig bool) ([]*Error, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type ignorePatternFlags []string

func (i *ignorePatternFlags) String() string { _ = "STUB: not implemented"; return "" }

func (i *ignorePatternFlags) Set(v string) error { _ = "STUB: not implemented"; return nil }

// Main is main function of actionlint. It takes command line arguments as string slice and returns
// exit status. The args should be entire arguments including the program name, usually given via
// os.Args.
func (cmd *Command) Main(args []string) int { _ = "STUB: not implemented"; return 0 }

// When -h or -help

// Linter found some issues, yay!
