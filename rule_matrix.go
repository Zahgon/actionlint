package actionlint

// RuleMatrix is a rule checker to check 'matrix' field of job.
type RuleMatrix struct {
	RuleBase
}

// NewRuleMatrix creates new RuleMatrix instance.
func NewRuleMatrix() *RuleMatrix { _ = "STUB: not implemented"; return nil }

// VisitJobPre is callback when visiting Job node before visiting its children.
func (rule *RuleMatrix) VisitJobPre(n *Job) error { _ = "STUB: not implemented"; return nil }

// Note:
// Any new value can be set in the section as new combination. It can add new value to existing
// column also.
//
// matrix:
//   os: [ubuntu-latest, macos-latest]
//   include:
//     - os: windows-latest
//       sh: pwsh

func (rule *RuleMatrix) checkDuplicateInRow(row *MatrixRow) { _ = "STUB: not implemented"; return }

// Give up when ${{ }} is specified

func isYAMLValueSubset(v, sub RawYAMLValue) bool {
	_ = "STUB: not implemented"
	// When the filter side is dynamically constructed with some expression, it is not possible to statically check if the filter
	// matches the value. To avoid false positives, assume such filter always matches to the value. (#414)
	// ```
	// matrix:
	//
	//	foo: ['a', 'b']
	//	exclude:
	//	  foo: ${{ fromJSON('...') }}
	//
	// ```
	return false
}

// `exclude` filter can match to objects in matrix as subset of them (#249).
// For example,
//
// matrix:
//   os:
//     - { name: Ubuntu, matrix: ubuntu }
//     - { name: Windows, matrix: windows }
//   arch:
//     - { name: ARM, matrix: arm }
//     - { name: Intel, matrix: intel }
//   exclude:
//     - os: { matrix: windows }
//       arch: { matrix: arm }
//
// The `exclude` filters out `{ os: { name: Windows, matrix: windows }, arch: {name: ARM, matrix: arm } }`

// When some item is constructed with ${{ }} dynamically, give up checking combinations (#261)

func (rule *RuleMatrix) checkExclude(m *Matrix) { _ = "STUB: not implemented"; return }

// Note: do not use quotesBuilder
