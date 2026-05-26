package actionlint

import (
	"io"
)

// LogLevel is log level of logger used in Linter instance.
type LogLevel int

const (
	// LogLevelNone does not output any log output.
	LogLevelNone LogLevel = 0
	// LogLevelVerbose shows verbose log output. This is equivalent to specifying -verbose option
	// to actionlint command.
	LogLevelVerbose = 1
	// LogLevelDebug shows all log output including debug information.
	LogLevelDebug = 2
)

// ColorOptionKind is kind of colorful output behavior.
type ColorOptionKind int

const (
	// ColorOptionKindAuto is kind to determine to colorize errors output automatically. It is
	// determined based on pty and $NO_COLOR environment variable. See document of fatih/color
	// for more details.
	ColorOptionKindAuto ColorOptionKind = iota
	// ColorOptionKindAlways is kind to always colorize errors output.
	ColorOptionKindAlways
	// ColorOptionKindNever is kind never to colorize errors output.
	ColorOptionKindNever
)

// LinterOptions is set of options for Linter instance. This struct is used for NewLinter factory
// function call. The zero value LinterOptions{} represents the default behavior.
type LinterOptions struct {
	// Verbose is flag if verbose log output is enabled.
	Verbose bool
	// Debug is flag if debug log output is enabled.
	Debug bool
	// LogWriter is io.Writer object to use to print log outputs. Note that error outputs detected
	// by the linter are not included in the log outputs.
	LogWriter io.Writer
	// Color is option for colorizing error outputs. See ColorOptionKind document for each enum values.
	Color ColorOptionKind
	// Oneline is flag if one line output is enabled. When enabling it, one error is output per one
	// line. It is useful when reading outputs from programs.
	Oneline bool
	// Shellcheck is executable for running shellcheck external command. It can be command name like
	// "shellcheck" or file path like "/path/to/shellcheck", "path/to/shellcheck". When this value
	// is empty, shellcheck won't run to check scripts in workflow file.
	Shellcheck string
	// Pyflakes is executable for running pyflakes external command. It can be command name like "pyflakes"
	// or file path like "/path/to/pyflakes", "path/to/pyflakes". When this value is empty, pyflakes
	// won't run to check scripts in workflow file.
	Pyflakes string
	// IgnorePatterns is list of regular expression to filter errors. The pattern is applied to error
	// messages. When an error is matched, the error is ignored.
	IgnorePatterns []string
	// ConfigFile is a path to config file. Empty string means no config file path is given. In
	// the case, actionlint will try to read config from .github/actionlint.yaml.
	ConfigFile string
	// Format is a custom template to format error messages. It must follow Go Template format and
	// contain at least one {{ }} placeholder. https://pkg.go.dev/text/template
	Format string
	// StdinFileName is a file name when reading input from stdin. When this value is empty, "<stdin>"
	// is used as the default value.
	StdinFileName string
	// WorkingDir is a file path to the current working directory. When this value is empty, os.Getwd
	// will be used to get a working directory.
	WorkingDir string
	// OnRulesCreated is a hook to add or remove the check rules. This function is called on checking
	// every workflow files. Rules created by Linter instance are passed to the argument and the
	// function should return the modified rules.
	// Note that syntax errors may be reported even if this function returns nil or an empty slice.
	OnRulesCreated func([]Rule) []Rule
	// More options will come here
}

// Linter is struct to lint workflow files.
type Linter struct {
	projects       *Projects
	out            io.Writer
	logOut         io.Writer
	logLevel       LogLevel
	oneline        bool
	shellcheck     string
	pyflakes       string
	ignorePats     IgnorePatterns
	stdin          string
	defaultConfig  *Config
	errFmt         *ErrorFormatter
	cwd            string
	onRulesCreated func([]Rule) []Rule
}

// NewLinter creates a new Linter instance.
// The out parameter is used to output errors from Linter instance. Set io.Discard if you don't
// want the outputs.
// The opts parameter is LinterOptions instance which configures behavior of linting.
func NewLinter(out io.Writer, opts *LinterOptions) (*Linter, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Allow colorful output on Windows

func (l *Linter) log(args ...interface{}) { _ = "STUB: not implemented"; return }

func (l *Linter) debug(format string, args ...interface{}) { _ = "STUB: not implemented"; return }

func (l *Linter) debugWriter() io.Writer { _ = "STUB: not implemented"; return *new(io.Writer) }

// GenerateDefaultConfig generates default config file at ".github/actionlint.yaml" in the project
// which the given directory path belongs to. When the directory path is empty, the current directory
// will be used instead.
func (l *Linter) GenerateDefaultConfig(dir string) error { _ = "STUB: not implemented"; return nil }

// LintRepository lints YAML workflow files and outputs the errors to given writer. It finds the
// nearest `.github/workflows` directory based on `dir` and applies lint rules to all YAML workflow
// files under the directory. When the directory path is empty, the current working directory will
// be used instead.
func (l *Linter) LintRepository(dir string) ([]*Error, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// LintDir lints all YAML workflow files in the given directory recursively.
func (l *Linter) LintDir(dir string, project *Project) ([]*Error, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// To make output deterministic, sort order of file paths

// LintFiles lints YAML workflow files and outputs the errors to given writer. It applies lint
// rules to all given files. The project parameter can be nil. In the case, a project is detected
// from the file path.
func (l *Linter) LintFiles(filepaths []string, project *Project) ([]*Error, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Each element of ws is accessed by single goroutine so mutex is unnecessary

// This method modifies state of l.projects so it cannot be called in parallel.
// Before entering goroutine, resolve project instance.

// #173

// Bound concurrency on reading files to avoid "too many files to open" error (issue #3)

// Use relative path if possible

// Ensure that all processes finish. `proc.wait()` must be called after `eg.Wait()`.
// Calling `WaitGroup.Add` after `WaitGroup.Wait` can cause a race condition (specifically when
// increasing the group count from 0 to 1 and calling `Wait` and at the same time).
// `WaitGroup.Add` is called in `proc.run()` and `WaitGroup.Wait` is called in `proc.wait()`.
// After traversing all workflows, `proc.run()` is no longer called so `proc.wait()` can be
// called safely.

// LintFile lints one YAML workflow file and outputs the errors to given writer. The project
// parameter can be nil. In the case, the project is detected from the given path.
func (l *Linter) LintFile(path string, project *Project) ([]*Error, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// LintStdin lints the content read from STDIN. The stdin parameter is a reader to read from STDIN,
// which is usually os.Stdin. The file name is determined by LinterOptions.StdinFileName. When the
// option is empty, "<stdin>" is the default value.
func (l *Linter) LintStdin(stdin io.Reader) ([]*Error, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Lint lints YAML workflow file content given as byte slice. The path parameter is used as file
// path where the content came from.
// When nil is passed to the project parameter, it tries to find the project from the path parameter.
func (l *Linter) Lint(path string, content []byte, project *Project) ([]*Error, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (l *Linter) check(
	path string,
	content []byte,
	project *Project,
	proc *concurrentProcess,
	localActions *LocalActionsCache,
	localReusableWorkflows *LocalReusableWorkflowCache,
) ([]*Error, error) {
	_ = "STUB: not implemented"
	// Note: This method is called to check multiple files in parallel.
	// It must be thread safe assuming fields of Linter are not modified while running.
	return nil, nil
}

// `-config-file` option has higher priority than repository config file

// Populate filename in the error

// Alias may duplicate errors

func (l *Linter) filterErrors(errs []*Error, cfgs []PathConfig) []*Error {
	_ = "STUB: not implemented"
	return nil
}

func (l *Linter) printErrors(errs []*Error, src []byte) { _ = "STUB: not implemented"; return }
