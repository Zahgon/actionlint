package actionlint

var allPermissionScopes = map[string][]string{
	"actions":             {"read", "write", "none"},
	"artifact-metadata":   {"read", "write", "none"},
	"attestations":        {"read", "write", "none"},
	"checks":              {"read", "write", "none"},
	"contents":            {"read", "write", "none"},
	"deployments":         {"read", "write", "none"},
	"discussions":         {"read", "write", "none"},
	"id-token":            {"write", "none"},
	"issues":              {"read", "write", "none"},
	"models":              {"read", "none"},
	"packages":            {"read", "write", "none"},
	"pages":               {"read", "write", "none"},
	"pull-requests":       {"read", "write", "none"},
	"repository-projects": {"read", "write", "none"},
	"security-events":     {"read", "write", "none"},
	"statuses":            {"read", "write", "none"},
}

// RulePermissions is a rule checker to check permission configurations in a workflow.
// https://docs.github.com/en/actions/reference/workflows-and-actions/workflow-syntax#defining-access-for-the-github_token-scopes
type RulePermissions struct {
	RuleBase
}

// NewRulePermissions creates new RulePermissions instance.
func NewRulePermissions() *RulePermissions { _ = "STUB: not implemented"; return nil }

// VisitJobPre is callback when visiting Job node before visiting its children.
func (rule *RulePermissions) VisitJobPre(n *Job) error { _ = "STUB: not implemented"; return nil }

// VisitWorkflowPre is callback when visiting Workflow node before visiting its children.
func (rule *RulePermissions) VisitWorkflowPre(n *Workflow) error {
	_ = "STUB: not implemented"
	return nil
}

func (rule *RulePermissions) checkPermissions(p *Permissions) { _ = "STUB: not implemented"; return }

// OK

// Permission names are case-sensitive
