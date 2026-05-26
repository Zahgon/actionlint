package actionlint

import (
	"io"
	"sync"

	"go.yaml.in/yaml/v4"
)

//go:generate go run ./scripts/generate-popular-actions ./popular_actions.go

// ActionMetadataInput is input metadata in "inputs" section in action.yml metadata file.
// https://docs.github.com/en/actions/creating-actions/metadata-syntax-for-github-actions#inputs
type ActionMetadataInput struct {
	// Name is a name of this input.
	Name string `json:"name"`
	// Required is true when this input is mandatory to run the action.
	Required bool `json:"required"`
	// Deprecated is true when this input is marked as deprecated.
	Deprecated bool `json:"deprecated"`
	// DeprecationMessage is a deprecation message for the deprecated input.
	DeprecationMessage string `json:"deprecation-message"`
}

// ActionMetadataInputs is a map from input ID to its metadata. Keys are in lower case since input
// names are case-insensitive.
// https://docs.github.com/en/actions/creating-actions/metadata-syntax-for-github-actions#inputs
type ActionMetadataInputs map[string]*ActionMetadataInput

// UnmarshalYAML implements yaml.Unmarshaler.
func (inputs *ActionMetadataInputs) UnmarshalYAML(n *yaml.Node) error {
	_ = "STUB: not implemented"
	return nil
}

// Value of `deprecationMessage: ` is `nil`. So we cannot determine the key exists or not by `v.Decode()` even
// if we change the type of `DeprecationMessage` to `*string`.

// OK

// Do not return this non-fatal error immediately so that this method still can return a parsed metadata
// even if the error occurs. This is useful for parsing somewhat broken metadata by generate-popular-actions
// script.

// ActionMetadataOutput is output metadata in "outputs" section in action.yml metadata file.
// https://docs.github.com/en/actions/creating-actions/metadata-syntax-for-github-actions#outputs-for-composite-actions
type ActionMetadataOutput struct {
	Name string `json:"name"`
}

// ActionMetadataOutputs is a map from output ID to its metadata. Keys are in lower case since output
// names are case-insensitive.
// https://docs.github.com/en/actions/creating-actions/metadata-syntax-for-github-actions#outputs-for-composite-actions
type ActionMetadataOutputs map[string]*ActionMetadataOutput

// UnmarshalYAML implements yaml.Unmarshaler.
func (inputs *ActionMetadataOutputs) UnmarshalYAML(n *yaml.Node) error {
	_ = "STUB: not implemented"
	return nil
}

// ActionMetadataRuns is "runs" section of action.yaml. It defines how the action is run.
// https://docs.github.com/en/actions/creating-actions/metadata-syntax-for-github-actions#runs
type ActionMetadataRuns struct {
	// Using is `using` configuration of action.yaml. It defines what runner is used for the action.
	Using string `yaml:"using" json:"using"`
	// Main is `main` configuration of action.yaml for JavaScript action.
	Main string `yaml:"main" json:"main"`
	// Pre is `pre` configuration of action.yaml for JavaScript action.
	Pre string `yaml:"pre" json:"pre"`
	// PreIf is `pre-if` configuration of action.yaml for JavaScript action.
	PreIf string `yaml:"pre-if" json:"pre-if"`
	// Post is `post` configuration of action.yaml for JavaScript action.
	Post string `yaml:"post" json:"post"`
	// PostIf is `post-if` configuration of action.yaml for JavaScript action.
	PostIf string `yaml:"post-if" json:"post-if"`
	// Steps is `steps` configuration of action.yaml for Composite action.
	Steps []any `yaml:"steps" json:"steps"`
	// Image is `image` of action.yaml for Docker action.
	Image string `yaml:"image" json:"image"`
	// PreEntrypoint is `pre-entrypoint` of action.yaml for Docker action.
	PreEntrypoint string `yaml:"pre-entrypoint" json:"pre-entrypoint"`
	// Entrypoint is `entrypoint` of action.yaml for Docker action.
	Entrypoint string `yaml:"entrypoint" json:"entrypoint"`
	// PostEntrypoint is `post-entrypoint` of action.yaml for Docker action.
	PostEntrypoint string `yaml:"post-entrypoint" json:"post-entrypoint"`
	// Args is `args` of action.yaml for Docker action.
	Args []any `yaml:"args" json:"args"`
	// Env is `env` of action.yaml for Docker action.
	Env map[string]any `yaml:"env" json:"env"`
}

// ActionMetadataBranding is "branding" section of action.yaml.
// https://docs.github.com/en/actions/creating-actions/metadata-syntax-for-github-actions#branding
type ActionMetadataBranding struct {
	Icon  string `yaml:"icon"`
	Color string `yaml:"color"`
}

// ActionMetadata represents structure of action.yaml.
// https://docs.github.com/en/actions/creating-actions/metadata-syntax-for-github-actions
type ActionMetadata struct {
	dir  string
	file string
	// Name is "name" field of action.yaml.
	Name string `yaml:"name" json:"name"`
	// Description is "description" field of action.yaml.
	Description string `yaml:"description" json:"-"`
	// Inputs is "inputs" field of action.yaml.
	Inputs ActionMetadataInputs `yaml:"inputs" json:"inputs"`
	// Outputs is "outputs" field of action.yaml. Key is name of output. Description is omitted
	// since actionlint does not use it.
	Outputs ActionMetadataOutputs `yaml:"outputs" json:"outputs"`
	// SkipInputs is flag to specify behavior of inputs check. When it is true, inputs for this
	// action will not be checked.
	SkipInputs bool `yaml:"-" json:"skip_inputs"`
	// SkipOutputs is flag to specify a bit loose typing to outputs object. If it is set to
	// true, the outputs object accepts any properties along with strictly typed props.
	SkipOutputs bool `yaml:"-" json:"skip_outputs"`
	// Runs is "runs" field of action.yaml.
	Runs ActionMetadataRuns `yaml:"runs" json:"runs"`
	// Branding is "branding" field of action.yaml.
	Branding ActionMetadataBranding `yaml:"branding" json:"-"`
}

// Dir returns a directory path of the action.
func (md *ActionMetadata) Dir() string {
	_ = "STUB: not implemented"

	// Path returns a file path of the action's metadata file.
	return ""
}

func (md *ActionMetadata) Path() string { _ = "STUB: not implemented"; return "" }

// LocalActionsCache is cache for local actions' metadata. It avoids repeating to find/read/parse
// local action's metadata file (action.yml).
// This cache is not available across multiple repositories. One LocalActionsCache instance needs
// to be created per one repository.
type LocalActionsCache struct {
	mu    sync.RWMutex
	proj  *Project // might be nil
	cache map[string]*ActionMetadata
	dbg   io.Writer
}

// NewLocalActionsCache creates new LocalActionsCache instance for the given project.
func NewLocalActionsCache(proj *Project, dbg io.Writer) *LocalActionsCache {
	_ = "STUB: not implemented"
	return nil
}

func newNullLocalActionsCache(dbg io.Writer) *LocalActionsCache {
	_ = "STUB: not implemented"
	// Null cache. Cache never hits. It is used when project is not found
	return nil
}

func (c *LocalActionsCache) debug(format string, args ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func (c *LocalActionsCache) readCache(key string) (*ActionMetadata, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func (c *LocalActionsCache) writeCache(key string, val *ActionMetadata) {
	_ = "STUB: not implemented"
	return
}

// FindMetadata finds metadata for given spec. The spec should indicate for local action hence it
// should start with "./". The first return value can be nil even if error did not occur.
// LocalActionCache caches that the action was not found. At first search, it returns an error that
// the action was not found. But at the second search, it does not return an error even if the result
// is nil. This behavior prevents repeating to report the same error from multiple places.
// Calling this method is thread-safe.
func (c *LocalActionsCache) FindMetadata(spec string) (*ActionMetadata, bool, error) {
	_ = "STUB: not implemented"
	return nil, false, nil
}

// Remember action was not found

// Do not complain about the action does not exist (#25, #40).
// It seems a common pattern that the local action does not exist in the repository
// (e.g. Git submodule) and it is cloned at running workflow (due to a private repository).

// Remember action was invalid

// Unwrap type error when a single type error occurs to simplify the error message

func (c *LocalActionsCache) readLocalActionMetadataFile(dir string) ([]byte, string, bool) {
	_ = "STUB: not implemented"
	return nil, "", false
}

// LocalActionsCacheFactory is a factory to create LocalActionsCache instances. LocalActionsCache
// should be created for each repositories. LocalActionsCacheFactory creates new LocalActionsCache
// instance per repository (project).
type LocalActionsCacheFactory struct {
	caches map[string]*LocalActionsCache
	dbg    io.Writer
}

// GetCache returns LocalActionsCache instance for the given project. One LocalActionsCache is
// created per one repository. Created instances are cached and will be used when caches are
// requested for the same projects. This method is not thread safe.
func (f *LocalActionsCacheFactory) GetCache(p *Project) *LocalActionsCache {
	_ = "STUB: not implemented"
	return nil
}

// NewLocalActionsCacheFactory creates a new LocalActionsCacheFactory instance.
func NewLocalActionsCacheFactory(dbg io.Writer) *LocalActionsCacheFactory {
	_ = "STUB: not implemented"
	return nil
}
