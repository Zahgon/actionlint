package actionlint

import "regexp"

var deprecatedCommandsPattern = regexp.MustCompile(`(?:::(save-state|set-output|set-env)\s+name=[a-zA-Z][a-zA-Z_-]*::\S+|::(add-path)::\S+)`)

// RuleDeprecatedCommands is a rule checker to detect deprecated workflow commands. Currently
// 'set-state', 'set-output', `set-env' and 'add-path' are detected as deprecated.
//
// - https://github.blog/changelog/2020-10-01-github-actions-deprecating-set-env-and-add-path-commands/
// - https://github.blog/changelog/2022-10-11-github-actions-deprecating-save-state-and-set-output-commands/
type RuleDeprecatedCommands struct {
	RuleBase
}

// NewRuleDeprecatedCommands creates a new RuleDeprecatedCommands instance.
func NewRuleDeprecatedCommands() *RuleDeprecatedCommands { _ = "STUB: not implemented"; return nil }

// VisitStep is callback when visiting Step node.
func (rule *RuleDeprecatedCommands) VisitStep(n *Step) error { _ = "STUB: not implemented"; return nil }
