package actionlint

func ordinal(i int) string { _ = "STUB: not implemented"; return "" }

// parseFormatFuncSpecifiers parses the format string passed to `format()` calls.
// https://docs.github.com/en/actions/writing-workflows/choosing-what-your-workflow-does/evaluate-expressions-in-workflows-and-actions#format
func parseFormatFuncSpecifiers(f string, n int) map[int]struct{} {
	_ = "STUB: not implemented"
	return nil
}

// Before specifier

// Opening brace

// Inside specifier

// Escaped '{'

// Empty specifier '{}'

// Closing brace.

// After specifier

// Parsing specifier needs to be delayed while '}' continues.

// Odd number of closing braces means the specifier closed. For example '{0}}}' contains
// a specifier but '{0}}}}' doesn't.

// When the input ends while '}' continues at end of specifier

// Functions

// FuncSignature is a signature of function, which holds return and arguments types.
type FuncSignature struct {
	// Name is a name of the function.
	Name string
	// Ret is a return type of the function.
	Ret ExprType
	// Params is a list of parameter types of the function. The final element of this list might
	// be repeated as variable length arguments.
	Params []ExprType
	// VariableLengthParams is a flag to handle variable length parameters. When this flag is set to
	// true, it means that the last type of params might be specified multiple times (including zero
	// times). Setting true implies length of Params is more than 0.
	VariableLengthParams bool
	// IsConstFunc is true when the function returns a constant when all parameters are constants.
	IsConstFunc bool
}

func (sig *FuncSignature) String() string { _ = "STUB: not implemented"; return "" }

// BuiltinFuncSignatures is a set of all builtin function signatures. All function names are in
// lower case because function names are compared in case insensitive.
// https://docs.github.com/en/actions/learn-github-actions/expressions#functions
var BuiltinFuncSignatures = map[string][]*FuncSignature{
	"contains": {
		{
			Name: "contains",
			Ret:  BoolType{},
			Params: []ExprType{
				StringType{},
				StringType{},
			},
			IsConstFunc: true,
		},
		{
			Name: "contains",
			Ret:  BoolType{},
			Params: []ExprType{
				&ArrayType{Elem: AnyType{}},
				AnyType{},
			},
			IsConstFunc: true,
		},
	},
	"startswith": {{
		Name: "startsWith",
		Ret:  BoolType{},
		Params: []ExprType{
			StringType{},
			StringType{},
		},
		IsConstFunc: true,
	}},
	"endswith": {{
		Name: "endsWith",
		Ret:  BoolType{},
		Params: []ExprType{
			StringType{},
			StringType{},
		},
		IsConstFunc: true,
	}},
	"format": {{
		Name: "format",
		Ret:  StringType{},
		Params: []ExprType{
			StringType{},
			AnyType{}, // variable length
		},
		VariableLengthParams: true,
		IsConstFunc:          true,
	}},
	"join": {
		{
			Name: "join",
			Ret:  StringType{},
			Params: []ExprType{
				&ArrayType{Elem: StringType{}},
				StringType{},
			},
			IsConstFunc: true,
		},
		// When the second parameter is omitted, values are concatenated with ','.
		{
			Name: "join",
			Ret:  StringType{},
			Params: []ExprType{
				&ArrayType{Elem: StringType{}},
			},
			IsConstFunc: true,
		},
	},
	"tojson": {{
		Name: "toJSON",
		Ret:  StringType{},
		Params: []ExprType{
			AnyType{},
		},
		IsConstFunc: true,
	}},
	"fromjson": {{
		Name: "fromJSON",
		Ret:  AnyType{},
		Params: []ExprType{
			StringType{},
		},
	}},
	"hashfiles": {{
		Name: "hashFiles",
		Ret:  StringType{},
		Params: []ExprType{
			StringType{},
		},
		VariableLengthParams: true,
	}},
	"success": {{
		Name:   "success",
		Ret:    BoolType{},
		Params: []ExprType{},
	}},
	"always": {{
		Name:   "always",
		Ret:    BoolType{},
		Params: []ExprType{},
	}},
	"cancelled": {{
		Name:   "cancelled",
		Ret:    BoolType{},
		Params: []ExprType{},
	}},
	"failure": {{
		Name:   "failure",
		Ret:    BoolType{},
		Params: []ExprType{},
	}},
	"case": {{
		Name: "case",
		Ret:  AnyType{},
		Params: []ExprType{
			BoolType{},
			AnyType{},
			AnyType{},
		},
		VariableLengthParams: true,
		IsConstFunc:          true,
	}},
}

// Global variables

// BuiltinGlobalVariableTypes defines types of all global variables. All context variables are
// documented at https://docs.github.com/en/actions/learn-github-actions/contexts
var BuiltinGlobalVariableTypes = map[string]ExprType{
	// https://docs.github.com/en/actions/writing-workflows/choosing-what-your-workflow-does/accessing-contextual-information-about-workflow-runs#github-context
	"github": NewStrictObjectType(map[string]ExprType{
		"action":                    StringType{},
		"action_path":               StringType{}, // Note: Composite actions only
		"action_ref":                StringType{},
		"action_repository":         StringType{},
		"action_status":             StringType{}, // Note: Composite actions only
		"actor":                     StringType{},
		"actor_id":                  StringType{},
		"api_url":                   StringType{},
		"artifact_cache_size_limit": NumberType{}, // Note: Undocumented
		"base_ref":                  StringType{},
		"env":                       StringType{},
		"event":                     NewEmptyObjectType(), // Note: Stricter type check for this payload would be possible
		"event_name":                StringType{},
		"event_path":                StringType{},
		"graphql_url":               StringType{},
		"head_ref":                  StringType{},
		"job":                       StringType{},
		"output":                    StringType{}, // Note: Undocumented
		"path":                      StringType{},
		"ref":                       StringType{},
		"ref_name":                  StringType{},
		"ref_protected":             BoolType{},
		"ref_type":                  StringType{},
		"repository":                StringType{},
		"repository_id":             StringType{},
		"repository_owner":          StringType{},
		"repository_owner_id":       StringType{},
		"repository_visibility":     StringType{}, // Note: Undocumented
		"repositoryurl":             StringType{}, // repositoryUrl
		"retention_days":            NumberType{},
		"run_attempt":               StringType{},
		"run_id":                    StringType{},
		"run_number":                StringType{},
		"secret_source":             StringType{},
		"server_url":                StringType{},
		"sha":                       StringType{},
		"state":                     StringType{}, // Note: Undocumented
		"step_summary":              StringType{}, // Note: Undocumented
		"token":                     StringType{},
		"triggering_actor":          StringType{},
		"workflow":                  StringType{},
		"workflow_ref":              StringType{},
		"workflow_sha":              StringType{},
		"workspace":                 StringType{},
	}),
	// https://docs.github.com/en/actions/learn-github-actions/contexts#env-context
	"env": NewMapObjectType(StringType{}), // env.<env_name>
	// https://docs.github.com/en/actions/learn-github-actions/contexts#job-context
	"job": NewStrictObjectType(map[string]ExprType{
		"check_run_id": NumberType{},
		"container": NewStrictObjectType(map[string]ExprType{
			"id":      StringType{},
			"network": StringType{},
		}),
		"services": NewMapObjectType(
			NewStrictObjectType(map[string]ExprType{
				"id":      StringType{}, // job.services.<service id>.id
				"network": StringType{},
				"ports":   NewMapObjectType(StringType{}),
			}),
		),
		"status": StringType{},
	}),
	// https://docs.github.com/en/actions/learn-github-actions/contexts#steps-context
	"steps": NewEmptyStrictObjectType(), // This value will be updated contextually
	// https://docs.github.com/en/actions/learn-github-actions/contexts#runner-context
	"runner": NewStrictObjectType(map[string]ExprType{
		"name":        StringType{},
		"os":          StringType{},
		"arch":        StringType{},
		"temp":        StringType{},
		"tool_cache":  StringType{},
		"debug":       StringType{},
		"environment": StringType{}, // https://github.com/github/docs/issues/32443
	}),
	// https://docs.github.com/en/actions/learn-github-actions/contexts#secrets-context
	"secrets": NewMapObjectType(StringType{}),
	// https://docs.github.com/en/actions/learn-github-actions/contexts#strategy-context
	"strategy": NewObjectType(map[string]ExprType{
		"fail-fast":    BoolType{},
		"job-index":    NumberType{},
		"job-total":    NumberType{},
		"max-parallel": NumberType{},
	}),
	// https://docs.github.com/en/actions/learn-github-actions/contexts#matrix-context
	"matrix": NewEmptyStrictObjectType(), // This value will be updated contextually
	// https://docs.github.com/en/actions/learn-github-actions/contexts#needs-context
	"needs": NewEmptyStrictObjectType(), // This value will be updated contextually
	// https://docs.github.com/en/actions/learn-github-actions/contexts#inputs-context
	// https://docs.github.com/en/actions/learn-github-actions/reusing-workflows
	"inputs": NewEmptyStrictObjectType(),
	// https://docs.github.com/en/actions/learn-github-actions/contexts#vars-context
	"vars": NewMapObjectType(StringType{}), // vars.<var_name>
}

// Semantics checker

// ExprSemanticsChecker is a semantics checker for expression syntax. It checks types of values
// in given expression syntax tree. It additionally checks other semantics like arguments of
// format() built-in function. To know the details of the syntax, see
//
// - https://docs.github.com/en/actions/learn-github-actions/contexts
// - https://docs.github.com/en/actions/learn-github-actions/expressions
type ExprSemanticsChecker struct {
	funcs                 map[string][]*FuncSignature
	vars                  map[string]ExprType
	errs                  []*ExprError
	varsCopied            bool
	githubVarCopied       bool
	untrusted             *UntrustedInputChecker
	availableContexts     []string
	availableSpecialFuncs []string
	configVars            []string
}

// NewExprSemanticsChecker creates new ExprSemanticsChecker instance. When checkUntrustedInput is
// set to true, the checker will make use of possibly untrusted inputs error.
func NewExprSemanticsChecker(checkUntrustedInput bool, configVars []string) *ExprSemanticsChecker {
	_ = "STUB: not implemented"
	return nil
}

func errorAtExpr(e ExprNode, msg string) *ExprError { _ = "STUB: not implemented"; return nil }

func errorfAtExpr(e ExprNode, format string, args ...interface{}) *ExprError {
	_ = "STUB: not implemented"
	return nil
}

func (sema *ExprSemanticsChecker) errorf(e ExprNode, format string, args ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func (sema *ExprSemanticsChecker) ensureVarsCopied() { _ = "STUB: not implemented"; return }

// Make shallow copy of current variables map not to pollute global variable

func (sema *ExprSemanticsChecker) ensureGithubVarCopied() { _ = "STUB: not implemented"; return }

// UpdateMatrix updates matrix object to given object type. Since matrix values change according to
// 'matrix' section of job configuration, the type needs to be updated.
func (sema *ExprSemanticsChecker) UpdateMatrix(ty *ObjectType) { _ = "STUB: not implemented"; return }

// UpdateSteps updates 'steps' context object to given object type.
func (sema *ExprSemanticsChecker) UpdateSteps(ty *ObjectType) { _ = "STUB: not implemented"; return }

// UpdateNeeds updates 'needs' context object to given object type.
func (sema *ExprSemanticsChecker) UpdateNeeds(ty *ObjectType) { _ = "STUB: not implemented"; return }

// UpdateSecrets updates 'secrets' context object to given object type.
func (sema *ExprSemanticsChecker) UpdateSecrets(ty *ObjectType) { _ = "STUB: not implemented"; return }

// Merges automatically supplied secrets with manually defined secrets.
// ACTIONS_STEP_DEBUG and ACTIONS_RUNNER_DEBUG seem supplied from caller of the workflow (#130)

// UpdateInputs updates 'inputs' context object to given object type.
func (sema *ExprSemanticsChecker) UpdateInputs(ty *ObjectType) { _ = "STUB: not implemented"; return }

// When both `workflow_call` and `workflow_dispatch` are the triggers of the workflow, `inputs` context can be used
// by both events. To cover both cases, merge `inputs` contexts into one object type. (#263)

// UpdateDispatchInputs updates 'github.event.inputs' and 'inputs' objects to given object type.
// https://docs.github.com/en/actions/reference/workflows-and-actions/events-that-trigger-workflows
func (sema *ExprSemanticsChecker) UpdateDispatchInputs(ty *ObjectType) {
	_ = "STUB: not implemented"
	return

	// Update `github.event.inputs`.
	// Unlike `inputs.*`, type of `github.event.inputs.*` is always string unlike `inputs.*`. We need
	// to create a new type from `ty` (e.g. {foo: boolean, bar: number} -> {foo: string, bar: string})
}

// UpdateJobs updates 'jobs' context object to given object type.
func (sema *ExprSemanticsChecker) UpdateJobs(ty *ObjectType) { _ = "STUB: not implemented"; return }

// SetContextAvailability sets available context names while semantics checks. Some contexts limit
// where they can be used.
// https://docs.github.com/en/actions/learn-github-actions/contexts#context-availability
//
// Elements of 'avail' parameter must be in lower case to check context names in case-insensitive.
//
// If this method is not called before checks, ExprSemanticsChecker considers any contexts are
// available by default.
// Available contexts for workflow keys can be obtained from actionlint.ContextAvailability.
func (sema *ExprSemanticsChecker) SetContextAvailability(avail []string) {
	_ = "STUB: not implemented"
	return
}

func (sema *ExprSemanticsChecker) checkAvailableContext(n *VariableNode) {
	_ = "STUB: not implemented"
	return
}

// SetSpecialFunctionAvailability sets names of available special functions while semantics checks.
// Some functions limit where they can be used.
// https://docs.github.com/en/actions/learn-github-actions/contexts#context-availability
//
// Elements of 'avail' parameter must be in lower case to check function names in case-insensitive.
//
// If this method is not called before checks, ExprSemanticsChecker considers no special function is
// allowed by default. Allowed functions can be obtained from actionlint.SpecialFunctionNames global
// constant.
//
// Available function names for workflow keys can be obtained from actionlint.ContextAvailability.
func (sema *ExprSemanticsChecker) SetSpecialFunctionAvailability(avail []string) {
	_ = "STUB: not implemented"
	return
}

func (sema *ExprSemanticsChecker) checkSpecialFunctionAvailability(n *FuncCallNode) {
	_ = "STUB: not implemented"
	return
}

// This function is not special

func (sema *ExprSemanticsChecker) visitUntrustedCheckerOnEnterNode(n ExprNode) {
	_ = "STUB: not implemented"
	return
}

func (sema *ExprSemanticsChecker) visitUntrustedCheckerOnLeaveNode(n ExprNode) {
	_ = "STUB: not implemented"
	return
}

func (sema *ExprSemanticsChecker) checkVariable(n *VariableNode) ExprType {
	_ = "STUB: not implemented"
	return *new(ExprType)
}

func (sema *ExprSemanticsChecker) checkObjectDeref(n *ObjectDerefNode) ExprType {
	_ = "STUB: not implemented"
	return *new(ExprType)
}

// When element type is any, map the any type to any. Reuse `ty`

// Map element type of delererenced array

func (sema *ExprSemanticsChecker) checkConfigVariables(n *ObjectDerefNode) {
	_ = "STUB: not implemented"
	// https://docs.github.com/en/actions/learn-github-actions/variables#naming-conventions-for-configuration-variables
	return
}

// Note: `n.Property` was already converted to lower case by parser
// Note: First character cannot be number, but it was already checked by parser

func (sema *ExprSemanticsChecker) checkArrayDeref(n *ArrayDerefNode) ExprType {
	_ = "STUB: not implemented"
	return *new(ExprType)
}

// Object filtering is available for objects, not only arrays (#66)

// For map object or loose object at receiver of .*

// For strict object at receiver of .*

func (sema *ExprSemanticsChecker) checkIndexAccess(n *IndexAccessNode) ExprType {
	_ = "STUB: not implemented"
	// Note: Index must be visited before Index to make UntrustedInputChecker work correctly even if
	// the expression has some nest like foo[aaa.bbb].bar. Nest happens in top-down order and
	// properties/indices access check is done in bottom-up order. So, as far as we visit nested
	// index nodes before visiting operand, the index is recursively checked first.
	return *new(ExprType)
}

// Index access with string literal like foo['bar']

// Fallback

func checkFuncSignature(n *FuncCallNode, sig *FuncSignature, args []ExprType) *ExprError {
	_ = "STUB: not implemented"
	return nil
}

// Note: Unlike many languages, this check does not allow 0 argument for the variable length
// parameter since it is useful for checking hashFiles() and format().

func (sema *ExprSemanticsChecker) checkBuiltinFuncCall(n *FuncCallNode, sig *FuncSignature) ExprType {
	_ = "STUB: not implemented"
	return *new(ExprType)
}

// Special checks for specific built-in functions

// -1 means removing first format string argument

// forget it to check unused placeholders

func (sema *ExprSemanticsChecker) checkFuncCall(n *FuncCallNode) ExprType {
	_ = "STUB: not implemented"
	// Check function name in case insensitive. For example, toJson and toJSON are the same function.
	return *new(ExprType)
}

// Check all overloads

// When one of overload pass type check, overload was resolved correctly

// All candidates failed

func (sema *ExprSemanticsChecker) checkNotOp(n *NotOpNode) ExprType {
	_ = "STUB: not implemented"
	return *new(ExprType)
}

func validateCompareOpOperands(op CompareOpNodeKind, l, r ExprType) bool {
	_ = "STUB: not implemented"
	// Comparison behavior: https://docs.github.com/en/actions/learn-github-actions/expressions#operators
	return false
}

// These are coerced to NaN hence the comparison result is always false

// null, bool, array, and object cannot be compared with these operators

func (sema *ExprSemanticsChecker) checkCompareOp(n *CompareOpNode) ExprType {
	_ = "STUB: not implemented"
	return *new(ExprType)
}

// checkWithNarrowing checks type of given expression with type narrowing. Type narrowing narrows
// down the type of the expression by assuming its value. For example, `l && r` is typed as
// `typeof(l) | typeof(r)` usually. However when the expression is assumed to be true, its type can
// be narrowed down to `typeof(r)`.
// This analysis is useful to make type checking more accurate. For example, `some_var && 60 || 20`
// can be typed as `number` instead of `typeof(some_var) | number`. (#384)
func (sema *ExprSemanticsChecker) checkWithNarrowing(n ExprNode, isTruthy bool) ExprType {
	_ = "STUB: not implemented"
	return *new(ExprType)
}

// When `l && r` is true, narrow its type to `typeof(r)`

// When `l || r` is false, narrow its type to `typeof(r)`

func (sema *ExprSemanticsChecker) checkLogicalOp(n *LogicalOpNode) ExprType {
	_ = "STUB: not implemented"
	return *new(ExprType)
}

// When `l` is false in `l && r`, its type is `typeof(l)`. Otherwise `typeof(r)`.
// Narrow the type of LHS expression by assuming its value is falsy.

// When `l` is true in `l || r`, its type is `typeof(l)`. Otherwise `typeof(r).
// Narrow the type of LHS expression by assuming its value is truthy.

func (sema *ExprSemanticsChecker) check(expr ExprNode) ExprType {
	_ = "STUB: not implemented"
	return *new(ExprType)
}

// Call this method in bottom-up order

// Check checks semantics of given expression syntax tree. It returns the type of the expression as
// the first return value when the check was successfully done. And it returns all errors found
// while checking the expression as the second return value.
func (sema *ExprSemanticsChecker) Check(expr ExprNode) (ExprType, []*ExprError) {
	_ = "STUB: not implemented"
	return *new(ExprType), nil
}

// IsConstant returns the given expression is a constant. For example the following expressions
// are constants.
//   - 1 == 2
//   - (!true && 'foo') || 'bar'
//   - startsWith('foobar', 'foo')
//   - format('{} + {} = {}', 1, 2, 3)
func (sema *ExprSemanticsChecker) IsConstant(expr ExprNode) bool {
	_ = "STUB: not implemented"
	return false
}
