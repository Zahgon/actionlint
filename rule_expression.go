package actionlint

//go:generate go run ./scripts/generate-availability ./availability.go

type typedExpr struct {
	ty  ExprType
	pos Pos
}

// RuleExpression is a rule checker to check expression syntax in string values of workflow syntax.
// It checks syntax and semantics of the expressions including type checks and functions/contexts
// definitions. For more details see
// - https://docs.github.com/en/actions/learn-github-actions/contexts
// - https://docs.github.com/en/actions/learn-github-actions/expressions
type RuleExpression struct {
	RuleBase
	matrixTy         *ObjectType
	stepsTy          *ObjectType
	needsTy          *ObjectType
	secretsTy        *ObjectType
	inputsTy         *ObjectType
	dispatchInputsTy *ObjectType
	jobsTy           *ObjectType
	workflow         *Workflow
	localActions     *LocalActionsCache
	localWorkflows   *LocalReusableWorkflowCache
}

// NewRuleExpression creates new RuleExpression instance.
func NewRuleExpression(actionsCache *LocalActionsCache, workflowCache *LocalReusableWorkflowCache) *RuleExpression {
	_ = "STUB: not implemented"
	return nil
}

// VisitWorkflowPre is callback when visiting Workflow node before visiting its children.
func (rule *RuleExpression) VisitWorkflowPre(n *Workflow) error {
	_ = "STUB: not implemented"
	return nil
}

// Set `inputs` context object before checking inputs since input's default values can refer `inputs` context.
//   inputs:
//     input1:
//       type: string
//     input2:
//       type: string
//       default: ${{ inputs.input1 }}

// Check default value before setting type to `ity` because referring myself should cause an error.
//   inputs:
//     recursive:
//       type: string
//       default: ${{ inputs.recursive }}

// ok

// ok

// When no secret is passed, secrets may be inherited from a caller of the workflow.
// So `secrets` context must be typed as { string => string }. `e.Secrets` is nil when `secrets:` does not
// exist. When `e.Secrets` is an empty map, `secrets:` exists but it has no child.

// o.Value will be checked in VisitWorkflowPost

// VisitWorkflowPost is callback when visiting Workflow node after visiting its children
func (rule *RuleExpression) VisitWorkflowPost(n *Workflow) error {
	_ = "STUB: not implemented"
	return nil
}

// VisitJobPre is callback when visiting Job node before visiting its children.
func (rule *RuleExpression) VisitJobPre(n *Job) error {
	_ = "STUB: not implemented"
	// Type of needs must be resolved before resolving type of matrix because `needs` context can
	// be used in matrix configuration.
	return nil
}

// Set matrix type at start of VisitJobPre() because matrix values are available in
// jobs.<job_id> section. For example:
//   jobs:
//     foo:
//       strategy:
//         matrix:
//           os: [ubuntu-latest, macos-latest, windows-latest]
//       runs-on: ${{ matrix.os }}

// Check and guess type of the matrix

// OK

// Note: Types in "jobs.<job_id>.strategy.matrix" were checked `checkMatrix`

// VisitJobPost is callback when visiting Job node after visiting its children
func (rule *RuleExpression) VisitJobPost(n *Job) error {
	_ = "STUB: not implemented"
	// 'environment' and 'outputs' sections are evaluated after all steps are run
	return nil
}

// VisitStep is callback when visiting Step node.
func (rule *RuleExpression) VisitStep(n *Step) error { _ = "STUB: not implemented"; return nil }

// env: at step level can refer 'env' context (#158)

// Step ID is case insensitive

// Get type of `outputs.<output name>`
func (rule *RuleExpression) getActionOutputsType(spec *String) *ObjectType {
	_ = "STUB: not implemented"
	return nil
}

// github-script action allows to set any outputs through calling `core.setOutput` directly.
// So any `outputs.*` properties should be accepted (#104)

// When the action run at this step is a popular action, we know what outputs are set by it.
// Set the output names to `steps.{step_id}.outputs.{name}`.

func (rule *RuleExpression) getWorkflowCallOutputsType(call *WorkflowCall) *ObjectType {
	_ = "STUB: not implemented"
	return nil
}

func (rule *RuleExpression) checkOneExpression(s *String, what, workflowKey string) ExprType {
	_ = "STUB: not implemented"
	// checkString is not available since it checks types for embedding values into a string
	return *new(ExprType)
}

// This case should be unreachable since only one ${{ }} is included is checked by parser

func (rule *RuleExpression) checkObjectTy(ty ExprType, pos *Pos, what string) ExprType {
	_ = "STUB: not implemented"
	return *new(ExprType)
}

func (rule *RuleExpression) checkArrayTy(ty ExprType, pos *Pos, what string) ExprType {
	_ = "STUB: not implemented"
	return *new(ExprType)
}

func (rule *RuleExpression) checkNumberTy(ty ExprType, pos *Pos, what string) ExprType {
	_ = "STUB: not implemented"
	return *new(ExprType)
}

func (rule *RuleExpression) checkObjectExpression(s *String, what, workflowKey string) ExprType {
	_ = "STUB: not implemented"
	return *new(ExprType)
}

func (rule *RuleExpression) checkArrayExpression(s *String, what, workflowKey string) ExprType {
	_ = "STUB: not implemented"
	return *new(ExprType)
}

func (rule *RuleExpression) checkNumberExpression(s *String, what, workflowKey string) ExprType {
	_ = "STUB: not implemented"
	return *new(ExprType)
}

func (rule *RuleExpression) checkEnv(env *Env, workflowKey string) {
	_ = "STUB: not implemented"
	return
}

// When form of "env: ${{...}}"

func (rule *RuleExpression) checkContainer(c *Container, workflowKey, childWorkflowKeyPrefix string) {
	_ = "STUB: not implemented"
	return
}

// e.g. jobs.<job_id>.container.credentials

// e.g. jobs.<job_id>.container.env.<env_id>

func (rule *RuleExpression) checkConcurrency(c *Concurrency, workflowKey string) {
	_ = "STUB: not implemented"
	return
}

func (rule *RuleExpression) checkDefaults(d *Defaults, workflowKey string) {
	_ = "STUB: not implemented"
	return
}

func (rule *RuleExpression) checkWorkflowCall(c *WorkflowCall) { _ = "STUB: not implemented"; return }

func (rule *RuleExpression) checkSnapshot(s *Snapshot) { _ = "STUB: not implemented"; return }

func (rule *RuleExpression) checkWebhookEventFilter(f *WebhookEventFilter) {
	_ = "STUB: not implemented"
	return
}

func (rule *RuleExpression) checkStrings(ss []*String, workflowKey string) {
	_ = "STUB: not implemented"
	return
}

func (rule *RuleExpression) checkIfCondition(str *String, workflowKey string) {
	_ = "STUB: not implemented"
	return
}

// Note:
// https://docs.github.com/en/actions/learn-github-actions/workflow-syntax-for-github-actions#jobsjob_idif
//
// > When you use expressions in an if conditional, you may omit the expression syntax (${{ }})
// > because GitHub automatically evaluates the if conditional as an expression, unless the
// > expression contains any operators. If the expression contains any operators, the expression
// > must be contained within ${{ }} to explicitly mark it for evaluation.
//
// This document is actually wrong. I confirmed that any strings without surrounding in ${{ }}
// are evaluated.
//
// - run: echo 'run'
//   if: '!false'
// - run: echo 'not run'
//   if: '!true'
// - run: echo 'run'
//   if: false || true
// - run: echo 'run'
//   if: true && true
// - run: echo 'not run'
//   if: true && false

// }} is necessary since lexer lexes it as end of tokens

func (rule *RuleExpression) checkTemplateEvaluatedType(ts []typedExpr) {
	_ = "STUB: not implemented"
	return
}

func (rule *RuleExpression) checkString(str *String, workflowKey string) []typedExpr {
	_ = "STUB: not implemented"
	return nil
}

func (rule *RuleExpression) checkScriptString(str *String, workflowKey string) {
	_ = "STUB: not implemented"
	return
}

func (rule *RuleExpression) checkBool(b *Bool, workflowKey string) {
	_ = "STUB: not implemented"
	return
}

// ok

func (rule *RuleExpression) checkInt(i *Int, workflowKey string) { _ = "STUB: not implemented"; return }

func (rule *RuleExpression) checkFloat(f *Float, workflowKey string) {
	_ = "STUB: not implemented"
	return
}

func (rule *RuleExpression) checkExprsIn(s string, pos *Pos, quoted, checkUntrusted bool, workflowKey string) ([]typedExpr, bool) {
	_ = "STUB: not implemented"
	// TODO: Line number is not correct when the string contains newlines.
	return nil, false
}

// when the string is quoted like 'foo' or "foo", column should be incremented

// 3 means removing "${{"

func (rule *RuleExpression) exprError(err *ExprError, lineBase, colBase int) {
	_ = "STUB: not implemented"
	return
}

func (rule *RuleExpression) checkSemanticsOfExprNode(expr ExprNode, line, col int, checkUntrusted bool, workflowKey string) (ExprType, bool) {
	_ = "STUB: not implemented"
	return *new(ExprType), false
}

func (rule *RuleExpression) checkSemantics(src string, line, col int, checkUntrusted bool, workflowKey string) (ExprType, int, bool) {
	_ = "STUB: not implemented"
	return *new(ExprType), 0, false
}

func (rule *RuleExpression) calcNeedsType(job *Job) *ObjectType {
	_ = "STUB: not implemented"
	// https://docs.github.com/en/actions/learn-github-actions/contexts#needs-context
	return nil
}

func (rule *RuleExpression) populateDependantNeedsTypes(out *ObjectType, job *Job, root *Job) {
	_ = "STUB: not implemented"
	return
}

// ID is case insensitive

// When cyclic dependency exists. This does not happen normally.

// Already added

// Do not collect outputs type from parent of parent recursively. (#151)

func (rule *RuleExpression) checkMatrixExpression(expr *String) *ObjectType {
	_ = "STUB: not implemented"
	return nil
}

// Consider properties in include section elements since 'include' section adds matrix values

func (rule *RuleExpression) checkMatrix(m *Matrix) *ObjectType {
	_ = "STUB: not implemented"
	return nil
}

// Check types of "exclude" but they are not used to guess type of matrix

// When the combination exists in 'matrix' section, merge type with existing one

func (rule *RuleExpression) checkMatrixRow(r *MatrixRow) ExprType {
	_ = "STUB: not implemented"
	return *new(ExprType)
}

// No element

func (rule *RuleExpression) checkWorkflowCallOutputs(outputs map[string]*WorkflowCallEventOutput, jobs map[string]*Job) {
	_ = "STUB: not implemented"
	return
}

// Outputs are not defined in jobs.<job_id> section when it is reusable workflow call.

func (rule *RuleExpression) checkRawYAMLValue(v RawYAMLValue) ExprType {
	_ = "STUB: not implemented"
	return *new(ExprType)
}

func (rule *RuleExpression) checkRawYAMLString(y *RawYAMLString) ExprType {
	_ = "STUB: not implemented"
	return *new(ExprType)
}

// Note that keywords are case sensitive. TRUE, FALSE, NULL are invalid named value.

func convertExprLineColToPos(line, col, lineBase, colBase int) *Pos {
	_ = "STUB: not implemented"
	// Line and column in ExprError are 1-based
	return nil
}

func typeOfActionOutputs(meta *ActionMetadata) *ObjectType {
	_ = "STUB: not implemented"
	// Some action sets outputs dynamically. Such outputs are not defined in action.yml. actionlint
	// cannot check such outputs statically so it allows any props (#18)
	return nil
}
