package actionlint

type nodeStatus int

const (
	nodeStatusNew nodeStatus = iota
	nodeStatusActive
	nodeStatusFinished
)

type jobNode struct {
	id       string
	needs    []string
	resolved []*jobNode
	status   nodeStatus
	pos      *Pos
}

type edge struct {
	from *jobNode
	to   *jobNode
}

// RuleJobNeeds is a rule to check 'needs' field in each job configuration. For more details, see
// https://docs.github.com/en/actions/learn-github-actions/workflow-syntax-for-github-actions#jobsjob_idneeds
type RuleJobNeeds struct {
	RuleBase
	nodes map[string]*jobNode
}

// NewRuleJobNeeds creates new RuleJobNeeds instance.
func NewRuleJobNeeds() *RuleJobNeeds { _ = "STUB: not implemented"; return nil }

func contains[T comparable](heystack []T, needle T) bool { _ = "STUB: not implemented"; return false }

// VisitJobPre is callback when visiting Job node before visiting its children.
func (rule *RuleJobNeeds) VisitJobPre(n *Job) error { _ = "STUB: not implemented"; return nil }

// Job ID is key of mapping. Key mapping is stored in lowercase since it is case
// insensitive. So values in 'needs' array must be compared in lowercase.

// VisitWorkflowPost is callback when visiting Workflow node after visiting its children.
func (rule *RuleJobNeeds) VisitWorkflowPost(n *Workflow) error {
	_ = "STUB: not implemented"
	// Resolve nodes
	return nil
}

// Note: Only the first cycle can be detected even if there are multiple cycles in "needs:" configurations.

// Start cycle from the smallest position to make the error message deterministic

func collectCycle(src *jobNode, edges map[*jobNode]*jobNode) bool {
	_ = "STUB: not implemented"
	return false
}

// Detect cyclic dependencies
// https://inzkyk.xyz/algorithms/depth_first_search/detecting_cycles/

func detectFirstCycle(nodes map[string]*jobNode) *edge { _ = "STUB: not implemented"; return nil }

func detectCyclicNode(v *jobNode) *edge { _ = "STUB: not implemented"; return nil }
