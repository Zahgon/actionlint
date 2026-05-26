package actionlint

import (
	"context"
	"sync"

	"golang.org/x/sync/errgroup"
	"golang.org/x/sync/semaphore"
)

// cmdExecution represents a single command line execution.
type cmdExecution struct {
	cmd           string
	args          []string
	stdin         string
	combineOutput bool
}

func (e *cmdExecution) run() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// Set stdin via an io.Reader so that exec.Cmd pipes the bytes to the child
// after Start(). Writing to cmd.StdinPipe() before Start() relies on the
// kernel pipe buffer being large enough to absorb the whole payload, which
// deadlocks on darwin when multiple workers run concurrently (issue #650).

// Reaches here when exit status is non-zero and stdout is not empty, shellcheck successfully found some errors

// concurrentProcess is a manager to run process concurrently. Since running process consumes OS
// resources, running too many processes concurrently causes some issues. On macOS, making too many
// process makes the parent process hang (see issue #3). And running processes which open files can
// cause the error "pipe: too many files to open". To avoid it, this type manages how many processes
// are run at once.
type concurrentProcess struct {
	ctx  context.Context
	sema *semaphore.Weighted
	wg   sync.WaitGroup
}

// newConcurrentProcess creates a new ConcurrentProcess instance. The `par` argument represents how
// many processes can be run in parallel. It is recommended to use the value returned from
// runtime.NumCPU() for the argument.
func newConcurrentProcess(par int) *concurrentProcess { _ = "STUB: not implemented"; return nil }

func (proc *concurrentProcess) run(eg *errgroup.Group, exec *cmdExecution, callback func([]byte, error) error) {
	_ = "STUB: not implemented"
	return
}

// wait waits all goroutines started by this concurrentProcess instance finish.
func (proc *concurrentProcess) wait() {
	_ = "STUB: not implemented"
	// Wait for all goroutines completing to shutdown
	return
}

// newCommandRunner creates new external command runner for given executable. The executable path
// is resolved in this function.
func (proc *concurrentProcess) newCommandRunner(exe string, combineOutput bool) (*externalCommand, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func resolveExternalCommand(exe string) (string, []string, error) {
	_ = "STUB: not implemented"
	return "", nil, nil
}

// Try to parse the string as a command line instead of a single executable file path.

// externalCommand is struct to run specific command concurrently with concurrentProcess bounding
// number of processes at the same time. This type manages fatal errors while running the command
// by using errgroup.Group. The wait() method must be called at the end for checking if some fatal
// error occurred.
type externalCommand struct {
	proc          *concurrentProcess
	eg            errgroup.Group
	exe           string
	args          []string
	combineOutput bool
}

// run runs the command with given arguments and stdin. The callback function is called after the
// process runs. First argument is stdout and the second argument is an error while running the
// process.
func (cmd *externalCommand) run(args []string, stdin string, callback func([]byte, error) error) {
	_ = "STUB: not implemented"
	return
}

// wait waits until all goroutines for this command finish. Note that it does not wait for
// goroutines for other commands.
func (cmd *externalCommand) wait() error { _ = "STUB: not implemented"; return nil }
