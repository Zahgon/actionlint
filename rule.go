package actionlint

import (
	"io"
)

// RuleBase is a struct to be a base of rule structs. Embed this struct to define default methods
// automatically
type RuleBase struct {
	name   string
	desc   string
	errs   []*Error
	dbg    io.Writer
	config *Config
}

// NewRuleBase creates a new RuleBase instance. It should be embedded to your own
// rule instance.
func NewRuleBase(name string, desc string) RuleBase {
	_ = "STUB: not implemented"
	return *new(RuleBase)
}

// VisitStep is callback when visiting Step node.
func (r *RuleBase) VisitStep(node *Step) error {
	_ = "STUB: not implemented"

	// VisitJobPre is callback when visiting Job node before visiting its children.
	return nil
}

func (r *RuleBase) VisitJobPre(node *Job) error {
	_ = "STUB: not implemented"

	// VisitJobPost is callback when visiting Job node after visiting its children.
	return nil
}

func (r *RuleBase) VisitJobPost(node *Job) error {
	_ = "STUB: not implemented"

	// VisitWorkflowPre is callback when visiting Workflow node before visiting its children.
	return nil
}

func (r *RuleBase) VisitWorkflowPre(node *Workflow) error {
	_ = "STUB: not implemented"

	// VisitWorkflowPost is callback when visiting Workflow node after visiting its children.
	return nil
}

func (r *RuleBase) VisitWorkflowPost(node *Workflow) error {
	_ = "STUB: not implemented"

	// Error creates a new error from the source position and the error message and stores it in the
	// rule instance. The errors can be accessed by Errs method.
	return nil
}

func (r *RuleBase) Error(pos *Pos, msg string) { _ = "STUB: not implemented"; return }

// Errorf reports a new error with the source position and the formatted error message and stores it
// in the rule instance. The errors can be accessed by Errs method.
func (r *RuleBase) Errorf(pos *Pos, format string, args ...interface{}) {
	_ = "STUB: not implemented"
	return
}

// Debug prints debug log to the output. The output is specified by the argument of EnableDebug method.
// By default, no output is set so debug log is not printed.
func (r *RuleBase) Debug(format string, args ...interface{}) { _ = "STUB: not implemented"; return }

// Errs returns errors found by the rule.
func (r *RuleBase) Errs() []*Error {
	_ = "STUB: not implemented"

	// Name returns the name of the rule.
	return nil
}

func (r *RuleBase) Name() string {
	_ = "STUB: not implemented"

	// Description returns the description of the rule.
	return ""
}

func (r *RuleBase) Description() string {
	_ = "STUB: not implemented"

	// EnableDebug enables debug output from the rule. Given io.Writer instance is used to print debug
	// information to console. Setting nil means disabling debug output.
	return ""
}

func (r *RuleBase) EnableDebug(out io.Writer) {
	_ = "STUB: not implemented"

	// SetConfig populates user configuration of actionlint to the rule. When no config is set, rules
	// should behave as if the default configuration is set.
	return
}

func (r *RuleBase) SetConfig(cfg *Config) {
	_ = "STUB: not implemented"

	// Config returns the user configuration of actionlint. When no config was set to this rule by SetConfig,
	// this method returns nil.
	return
}

func (r *RuleBase) Config() *Config {
	_ = "STUB: not implemented"

	// Rule is an interface which all rule structs must meet.
	return nil
}

type Rule interface {
	Pass
	Errs() []*Error
	Name() string
	Description() string
	EnableDebug(out io.Writer)
	SetConfig(cfg *Config)
	Config() *Config
}
