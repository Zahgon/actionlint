package actionlint

// RuleWorkflowCall is a rule checker to check workflow call at jobs.<job_id>.
type RuleWorkflowCall struct {
	RuleBase
	workflowCallEventPos *Pos
	workflowPath         string
	cache                *LocalReusableWorkflowCache
}

// NewRuleWorkflowCall creates a new RuleWorkflowCall instance. 'workflowPath' is a file path to
// the workflow which is relative to a project root directory or an absolute path.
func NewRuleWorkflowCall(workflowPath string, cache *LocalReusableWorkflowCache) *RuleWorkflowCall {
	_ = "STUB: not implemented"
	return nil
}

// VisitWorkflowPre is callback when visiting Workflow node before visiting its children.
func (rule *RuleWorkflowCall) VisitWorkflowPre(n *Workflow) error {
	_ = "STUB: not implemented"
	return nil
}

// Register this reusable workflow in cache so that it does not need to parse this workflow
// file again when this workflow is called by other workflows.

// VisitJobPre is callback when visiting Job node before visiting its children.
func (rule *RuleWorkflowCall) VisitJobPre(n *Job) error { _ = "STUB: not implemented"; return nil }

// When the specification is invalid and it is local reusable workflow call, remember it caused
// an error by setting `nil` to cache. This can prevent redundant 'could not read workflow call'
// error.

func (rule *RuleWorkflowCall) checkWorkflowCallUsesLocal(call *WorkflowCall) {
	_ = "STUB: not implemented"
	return
}

// Validate inputs

// Validate secrets

// Parse ./{path/{filename}
// https://docs.github.com/en/actions/learn-github-actions/reusing-workflows#calling-a-reusable-workflow
func isWorkflowCallUsesLocalFormat(u string) bool { _ = "STUB: not implemented"; return false }

// Cannot container a ref

// Parse {owner}/{repo}/{path to workflow.yml}@{ref}
// https://docs.github.com/en/actions/learn-github-actions/reusing-workflows#calling-a-reusable-workflow
func isWorkflowCallUsesRepoFormat(u string) bool {
	_ = "STUB: not implemented"
	// Repo reference must start with owner
	return false
}

// Eat owner

// Eat repo

// Eat workflow path
