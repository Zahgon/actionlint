package actionlint

import (
	"io"
	"time"
)

// Pass is an interface to traverse a workflow syntax tree
type Pass interface {
	// VisitStep is callback when visiting Step node. It returns internal error when it cannot continue the process
	VisitStep(node *Step) error
	// VisitJobPre is callback when visiting Job node before visiting its children. It returns internal error when it cannot continue the process
	VisitJobPre(node *Job) error
	// VisitJobPost is callback when visiting Job node after visiting its children. It returns internal error when it cannot continue the process
	VisitJobPost(node *Job) error
	// VisitWorkflowPre is callback when visiting Workflow node before visiting its children. It returns internal error when it cannot continue the process
	VisitWorkflowPre(node *Workflow) error
	// VisitWorkflowPost is callback when visiting Workflow node after visiting its children. It returns internal error when it cannot continue the process
	VisitWorkflowPost(node *Workflow) error
}

// Visitor visits syntax tree from root in depth-first order
type Visitor struct {
	passes []Pass
	dbg    io.Writer
}

// NewVisitor creates Visitor instance
func NewVisitor() *Visitor {
	_ = "STUB: not implemented"

	// AddPass adds new pass which is called on traversing a syntax tree
	return nil
}

func (v *Visitor) AddPass(p Pass) { _ = "STUB: not implemented"; return }

// EnableDebug enables debug output when non-nil io.Writer value is given. All debug outputs from
// visitor will be written to the writer.
func (v *Visitor) EnableDebug(w io.Writer) { _ = "STUB: not implemented"; return }

func (v *Visitor) reportElapsedTime(what string, start time.Time) {
	_ = "STUB: not implemented"
	return
}

// Visit visits given syntax tree in depth-first order
func (v *Visitor) Visit(n *Workflow) error { _ = "STUB: not implemented"; return nil }

func (v *Visitor) visitJob(n *Job) error { _ = "STUB: not implemented"; return nil }

func (v *Visitor) visitStep(n *Step) error { _ = "STUB: not implemented"; return nil }
