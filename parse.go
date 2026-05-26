package actionlint

import (
	"iter"

	"go.yaml.in/yaml/v4"
)

// https://pkg.go.dev/go.yaml.in/yaml/v4#Kind
func nodeKindName(k yaml.Kind) string { _ = "STUB: not implemented"; return "" }

// Should be unreachable because we resolve all aliases before parsing

func posAt(n *yaml.Node) *Pos { _ = "STUB: not implemented"; return nil }

func newString(n *yaml.Node) *String { _ = "STUB: not implemented"; return nil }

// workflowMappingEntry represents a key-value entry in YAML mapping.
type workflowMappingEntry struct {
	// id is a key in lower case for comparing case-insensitive keys.
	id  string
	key *String
	val *yaml.Node
}

type delayedSprintf struct {
	result string
	// Note: Currently only one string arg is sufficient and it's faster than keeping generic interface{} args.
	// `arg` must not be empty when it is used. Empty value means the argument is unused.
	arg string
}

func sprintf(fmt, arg string) delayedSprintf {
	_ = "STUB: not implemented"
	return *new(delayedSprintf)
}

func (l *delayedSprintf) String() string {
	_ = "STUB: not implemented"

	// This delayed formatting reduces the number of allocations on parsing workflow by 4.94%
	return ""
}

type parser struct {
	errors []*Error
}

func (p *parser) error(n *yaml.Node, m string) { _ = "STUB: not implemented"; return }

func (p *parser) errorAt(pos *Pos, m string) { _ = "STUB: not implemented"; return }

func (p *parser) errorfAt(pos *Pos, format string, args ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func (p *parser) errorf(n *yaml.Node, format string, args ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func (p *parser) resolveAliases(root *yaml.Node) { _ = "STUB: not implemented"; return }

// For recursive call

// Note: Unknown anchors are detected by go-yaml parser so we don't need to detect them by ourselves.

// Resolved

// Don't resolve the recursive alias because it causes stack overflow on parsing the tree as
// `RawYAMLValue`. (#610)

func (p *parser) unexpectedKey(s *String, sec string, expected []string) {
	_ = "STUB: not implemented"
	return
}

func (p *parser) checkNotEmpty(sec string, len int, n *yaml.Node) bool {
	_ = "STUB: not implemented"
	return false
}

func (p *parser) checkSequence(sec string, n *yaml.Node, allowEmpty bool) bool {
	_ = "STUB: not implemented"
	return false
}

func (p *parser) checkString(n *yaml.Node, allowEmpty bool) bool {
	_ = "STUB: not implemented"
	// Do not check n.Tag is !!str because we don't need to check the node is string strictly.
	// In almost all cases, other nodes (like 42) are handled as string with its string representation.
	return false
}

func (p *parser) missingExpression(n *yaml.Node, expecting string) {
	_ = "STUB: not implemented"
	return
}

func (p *parser) parseExpression(n *yaml.Node, expecting string) *String {
	_ = "STUB: not implemented"
	return nil
}

func (p *parser) mayParseExpression(n *yaml.Node) *String { _ = "STUB: not implemented"; return nil }

func (p *parser) parseString(n *yaml.Node, allowEmpty bool) *String {
	_ = "STUB: not implemented"
	return nil
}

func (p *parser) parseStringSequence(sec string, n *yaml.Node, allowEmpty bool, allowElemEmpty bool) []*String {
	_ = "STUB: not implemented"
	return nil
}

func (p *parser) parseStringOrStringSequence(sec string, n *yaml.Node, allowEmpty bool, allowElemEmpty bool) []*String {
	_ = "STUB: not implemented"
	return nil
}

// In the case of 'foo:'

func (p *parser) parseBool(n *yaml.Node) *Bool { _ = "STUB: not implemented"; return nil }

func (p *parser) parseInt(n *yaml.Node) *Int { _ = "STUB: not implemented"; return nil }

func (p *parser) parseFloat(n *yaml.Node) *Float { _ = "STUB: not implemented"; return nil }

func (p *parser) parseMapping(where delayedSprintf, n *yaml.Node, allowEmpty, caseSensitive bool) iter.Seq[workflowMappingEntry] {
	_ = "STUB: not implemented"
	return nil
}

// Keys of mappings are sometimes case insensitive. For example, following matrix is invalid.
//   matrix:
//     foo: [1, 2, 3]
//     FOO: [1, 2, 3]
// To detect case insensitive duplicate keys, we use lowercase keys

func (p *parser) parseSectionMapping(section string, n *yaml.Node, allowEmpty, caseSensitive bool) iter.Seq[workflowMappingEntry] {
	_ = "STUB: not implemented"
	return nil
}

func (p *parser) parseMappingAt(where string, n *yaml.Node, allowEmpty, caseSensitive bool) iter.Seq[workflowMappingEntry] {
	_ = "STUB: not implemented"
	return nil
}

func (p *parser) parseScheduleEvent(pos *Pos, n *yaml.Node) *ScheduledEvent {
	_ = "STUB: not implemented"
	return nil
}

// https://docs.github.com/en/actions/reference/workflows-and-actions/workflow-syntax#onworkflow_dispatchinputs
func (p *parser) parseWorkflowDispatchEventInput(name *String, n *yaml.Node) *DispatchInput {
	_ = "STUB: not implemented"
	return nil
}

// - https://docs.github.com/en/actions/reference/workflows-and-actions/events-that-trigger-workflows#workflow_dispatch
// - https://docs.github.com/en/actions/reference/workflows-and-actions/workflow-syntax#onworkflow_dispatch
func (p *parser) parseWorkflowDispatchEvent(pos *Pos, n *yaml.Node) *WorkflowDispatchEvent {
	_ = "STUB: not implemented"
	return nil
}

// https://docs.github.com/en/actions/reference/workflows-and-actions/events-that-trigger-workflows#repository_dispatch
func (p *parser) parseRepositoryDispatchEvent(pos *Pos, n *yaml.Node) *RepositoryDispatchEvent {
	_ = "STUB: not implemented"
	return nil
}

// Note: Omitting 'types' is ok. In the case, all types trigger the workflow

// https://docs.github.com/en/actions/using-workflows/workflow-syntax-for-github-actions#using-filters
func (p *parser) parseWebhookEventFilter(name *String, n *yaml.Node) *WebhookEventFilter {
	_ = "STUB: not implemented"
	return nil
}

func (p *parser) parseWebhookEvent(name *String, n *yaml.Node) *WebhookEvent {
	_ = "STUB: not implemented"
	return nil
}

// Note: 'tags', 'tags-ignore', 'branches', 'branches-ignore' can be empty. Since there are
// some cases where setting empty values to them is necessary.
//
// > If only define only tags filter (tags/tags-ignore) or only branches filter
// > (branches/branches-ignore) for on.push, the workflow won’t run for events affecting the
// > undefined Git ref.
//
// https://github.community/t/using-on-push-tags-ignore-and-paths-ignore-together/16931

// Note: Glob pattern cannot be empty, but it is checked by 'glob' rule with better error
// message. So parser allows empty patterns here.

// https://docs.github.com/en/actions/learn-github-actions/workflow-syntax-for-github-actions#onworkflow_callinputs
func (p *parser) parseWorkflowCallEventInput(id string, name *String, n *yaml.Node) *WorkflowCallEventInput {
	_ = "STUB: not implemented"
	return nil
}

// https://docs.github.com/en/actions/reference/workflows-and-actions/workflow-syntax#example-of-onworkflow_callsecrets
func (p *parser) parseWorkflowCallEventSecret(name *String, n *yaml.Node) *WorkflowCallEventSecret {
	_ = "STUB: not implemented"
	return nil
}

// https://docs.github.com/en/actions/reference/workflows-and-actions/workflow-syntax#example-of-onworkflow_calloutputs
func (p *parser) parseWorkflowCallEventOutput(name *String, n *yaml.Node) *WorkflowCallEventOutput {
	_ = "STUB: not implemented"
	return nil
}

// - https://docs.github.com/en/actions/reference/workflows-and-actions/events-that-trigger-workflows#workflow-reuse-events
// - https://docs.github.com/en/actions/reference/workflows-and-actions/workflow-syntax#onworkflow_call
// - https://docs.github.com/en/actions/learn-github-actions/reusing-workflows
func (p *parser) parseWorkflowCallEvent(pos *Pos, n *yaml.Node) *WorkflowCallEvent {
	_ = "STUB: not implemented"
	return nil
}

func (p *parser) parseImageVersionEvent(pos *Pos, n *yaml.Node) *ImageVersionEvent {
	_ = "STUB: not implemented"
	return nil
}

func (p *parser) parseEventWithNoConfig(n *yaml.Node) Event {
	_ = "STUB: not implemented"
	return *new(Event)
}

func (p *parser) parseEvents(n *yaml.Node) []Event { _ = "STUB: not implemented"; return nil }

// https://docs.github.com/en/actions/learn-github-actions/workflow-syntax-for-github-actions#permissions
func (p *parser) parsePermissions(pos *Pos, n *yaml.Node) *Permissions {
	_ = "STUB: not implemented"
	return nil
}

// XXX: Is the permission scope case insensitive?

// https://docs.github.com/en/actions/learn-github-actions/workflow-syntax-for-github-actions#env
func (p *parser) parseEnv(n *yaml.Node) *Env { _ = "STUB: not implemented"; return nil }

// https://docs.github.com/en/actions/learn-github-actions/workflow-syntax-for-github-actions#defaults
func (p *parser) parseDefaults(pos *Pos, n *yaml.Node) *Defaults {
	_ = "STUB: not implemented"
	return nil
}

// https://docs.github.com/en/actions/learn-github-actions/workflow-syntax-for-github-actions#jobsjob_idconcurrency
func (p *parser) parseConcurrency(pos *Pos, n *yaml.Node) *Concurrency {
	_ = "STUB: not implemented"
	return nil
}

// https://docs.github.com/en/actions/learn-github-actions/workflow-syntax-for-github-actions#jobsjob_idenvironment
func (p *parser) parseEnvironment(pos *Pos, n *yaml.Node) *Environment {
	_ = "STUB: not implemented"
	return nil
}

// https://docs.github.com/en/actions/learn-github-actions/workflow-syntax-for-github-actions#jobsjob_idoutputs
func (p *parser) parseOutputs(n *yaml.Node) map[string]*Output {
	_ = "STUB: not implemented"
	return nil
}

func (p *parser) parseRawYAMLValue(n *yaml.Node) RawYAMLValue {
	_ = "STUB: not implemented"
	return *new(RawYAMLValue)
}

// https://docs.github.com/en/actions/learn-github-actions/workflow-syntax-for-github-actions#example-including-additional-values-into-combinations
// https://docs.github.com/en/actions/learn-github-actions/workflow-syntax-for-github-actions#example-excluding-configurations-from-a-matrix
func (p *parser) parseMatrixCombinations(sec string, n *yaml.Node) *MatrixCombinations {
	_ = "STUB: not implemented"
	return nil
}

// https://docs.github.com/en/actions/learn-github-actions/workflow-syntax-for-github-actions#jobsjob_idstrategymatrix
func (p *parser) parseMatrix(pos *Pos, n *yaml.Node) *Matrix { _ = "STUB: not implemented"; return nil }

// https://docs.github.com/en/actions/using-workflows/workflow-syntax-for-github-actions#jobsjob_idstrategymax-parallel
func (p *parser) parseMaxParallel(n *yaml.Node) *Int { _ = "STUB: not implemented"; return nil }

// https://docs.github.com/en/actions/learn-github-actions/workflow-syntax-for-github-actions#jobsjob_idstrategy
func (p *parser) parseStrategy(pos *Pos, n *yaml.Node) *Strategy {
	_ = "STUB: not implemented"
	return nil
}

func (p *parser) parseCredentials(pos *Pos, n *yaml.Node) *Credentials {
	_ = "STUB: not implemented"
	return nil
}

// https://docs.github.com/en/actions/learn-github-actions/workflow-syntax-for-github-actions#jobsjob_idcontainer
func (p *parser) parseContainer(sec string, pos *Pos, n *yaml.Node) *Container {
	_ = "STUB: not implemented"
	return nil
}

// When you only specify a container image, you can omit the image keyword.

func (p *parser) parseServices(n *yaml.Node) *Services { _ = "STUB: not implemented"; return nil }

// XXX: Is the key case-insensitive?

// https://docs.github.com/en/actions/using-workflows/workflow-syntax-for-github-actions#jobsjob_idtimeout-minutes
func (p *parser) parseTimeoutMinutes(n *yaml.Node) *Float { _ = "STUB: not implemented"; return nil }

func (p *parser) parseStepExecAction(entries []workflowMappingEntry, isDocker bool) *ExecAction {
	_ = "STUB: not implemented"
	return nil
}

// https://docs.github.com/en/actions/learn-github-actions/workflow-syntax-for-github-actions#jobsjob_idstepswithentrypoint

// https://docs.github.com/en/actions/learn-github-actions/workflow-syntax-for-github-actions#jobsjob_idstepswithargs

// do nothing

// Note: `ret.Uses` is never `nil` because `parseStep` checks `uses` key in advance

func (p *parser) parseStepExecRun(entries []workflowMappingEntry) *ExecRun {
	_ = "STUB: not implemented"
	return nil
}

// do nothing

// Note: `ret.Run` is never `nil` because `parseStep` checks `run` key in advance

// https://docs.github.com/en/actions/learn-github-actions/workflow-syntax-for-github-actions#jobsjob_idsteps
func (p *parser) parseStep(n *yaml.Node) *Step { _ = "STUB: not implemented"; return nil }

// Note: Unexpected keys are checked in parseStepExecAction or parseStepExecRun later

// https://docs.github.com/en/actions/learn-github-actions/workflow-syntax-for-github-actions#jobsjob_idsteps
func (p *parser) parseSteps(n *yaml.Node) []*Step { _ = "STUB: not implemented"; return nil }

// https://docs.github.com/en/actions/using-workflows/workflow-syntax-for-github-actions#jobsjob_idruns-on
func (p *parser) parseRunsOn(n *yaml.Node) *Runner { _ = "STUB: not implemented"; return nil }

func (p *parser) parseSnapshot(pos *Pos, n *yaml.Node) *Snapshot {
	_ = "STUB: not implemented"
	return nil
}

// https://docs.github.com/en/actions/learn-github-actions/workflow-syntax-for-github-actions#jobsjob_id
func (p *parser) parseJob(id *String, n *yaml.Node) *Job { _ = "STUB: not implemented"; return nil }

// Only below keys are allowed on reusable workflow call
// https://docs.github.com/en/actions/learn-github-actions/reusing-workflows#supported-keywords-for-jobs-that-call-a-reusable-workflow
//   - jobs.<job_id>.name
//   - jobs.<job_id>.uses
//   - jobs.<job_id>.with
//   - jobs.<job_id>.with.<input_id>
//   - jobs.<job_id>.secrets
//   - jobs.<job_id>.secrets.<secret_id>
//   - jobs.<job_id>.needs
//   - jobs.<job_id>.if
//   - jobs.<job_id>.permissions

// https://docs.github.com/en/actions/using-workflows/reusing-workflows#supported-keywords-for-jobs-that-call-a-reusable-workflow

// needs: job1

// needs: [job1, job2]

// `secrets: inherit` special case
// https://docs.github.com/en/actions/using-workflows/workflow-syntax-for-github-actions#onworkflow_callsecretsinherit

// When not a reusable call

// https://docs.github.com/en/actions/learn-github-actions/workflow-syntax-for-github-actions#jobs
func (p *parser) parseJobs(n *yaml.Node) map[string]*Job { _ = "STUB: not implemented"; return nil }

// https://docs.github.com/en/actions/learn-github-actions/workflow-syntax-for-github-actions
func (p *parser) parse(n *yaml.Node) *Workflow { _ = "STUB: not implemented"; return nil }

// func dumpYAML(n *yaml.Node, level int) {
// 	fmt.Printf("%s%s (%s, %d,%d): %q\n", strings.Repeat(". ", level), nodeKindName(n.Kind), n.Tag, n.Line, n.Column, n.Value)
// 	for _, c := range n.Content {
// 		dumpYAML(c, level+1)
// 	}
// }

func handleYAMLUnmarshalError(err error) []*Error { _ = "STUB: not implemented"; return nil }

// Fallback. I believe this line should be unreachable

// Parse parses given source as byte sequence into workflow syntax tree. It returns all errors
// detected while parsing the input. It means that detecting one error does not stop parsing. Even
// if one or more errors are detected, parser will try to continue parsing and finding more errors.
func Parse(b []byte) (*Workflow, []*Error) { _ = "STUB: not implemented"; return nil, nil }

// Uncomment for checking YAML tree
// dumpYAML(&n, 0)
