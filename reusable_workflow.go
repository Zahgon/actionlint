package actionlint

import (
	"io"
	"sync"

	"go.yaml.in/yaml/v4"
)

func expectedMapping(where string, n *yaml.Node) error { _ = "STUB: not implemented"; return nil }

// ReusableWorkflowMetadataInput is an input metadata for validating local reusable workflow file.
type ReusableWorkflowMetadataInput struct {
	// Name is a name of the input defined in the reusable workflow.
	Name string
	// Required is true when 'required' field of the input is set to true and no default value is set.
	Required bool
	// Type is a type of the input. When the input type is unknown, 'any' type is set.
	Type ExprType
}

// UnmarshalYAML implements yaml.Unmarshaler.
func (input *ReusableWorkflowMetadataInput) UnmarshalYAML(n *yaml.Node) error {
	_ = "STUB: not implemented"
	return nil
}

// ReusableWorkflowMetadataInputs is a map from input name to reusable wokflow input metadata. The
// keys are in lower case since input names of workflow calls are case insensitive.
type ReusableWorkflowMetadataInputs map[string]*ReusableWorkflowMetadataInput

// UnmarshalYAML implements yaml.Unmarshaler.
func (inputs *ReusableWorkflowMetadataInputs) UnmarshalYAML(n *yaml.Node) error {
	_ = "STUB: not implemented"
	return nil
}

// Reach here when `v` is null node

// ReusableWorkflowMetadataSecret is a secret metadata for validating local reusable workflow file.
type ReusableWorkflowMetadataSecret struct {
	// Name is a name of the secret in the reusable workflow.
	Name string
	// Required indicates whether the secret is required by its reusable workflow. When this value
	// is true, workflow calls must set this secret unless secrets are not inherited.
	Required bool `yaml:"required"`
}

// ReusableWorkflowMetadataSecrets is a map from secret name to reusable wokflow secret metadata.
// The keys are in lower case since secret names of workflow calls are case insensitive.
type ReusableWorkflowMetadataSecrets map[string]*ReusableWorkflowMetadataSecret

// UnmarshalYAML implements yaml.Unmarshaler.
func (secrets *ReusableWorkflowMetadataSecrets) UnmarshalYAML(n *yaml.Node) error {
	_ = "STUB: not implemented"
	return nil
}

// ReusableWorkflowMetadataOutput is an output metadata for validating local reusable workflow file.
type ReusableWorkflowMetadataOutput struct {
	// Name is a name of the output in the reusable workflow.
	Name string
}

// ReusableWorkflowMetadataOutputs is a map from output name to reusable wokflow output metadata.
// The keys are in lower case since output names of workflow calls are case insensitive.
type ReusableWorkflowMetadataOutputs map[string]*ReusableWorkflowMetadataOutput

// UnmarshalYAML implements yaml.Unmarshaler.
func (outputs *ReusableWorkflowMetadataOutputs) UnmarshalYAML(n *yaml.Node) error {
	_ = "STUB: not implemented"
	return nil
}

// ReusableWorkflowMetadata is metadata to validate local reusable workflows. This struct does not
// contain all metadata from YAML file. It only contains metadata which is necessary to validate
// reusable workflow files by actionlint.
type ReusableWorkflowMetadata struct {
	Inputs  ReusableWorkflowMetadataInputs  `yaml:"inputs"`
	Outputs ReusableWorkflowMetadataOutputs `yaml:"outputs"`
	Secrets ReusableWorkflowMetadataSecrets `yaml:"secrets"`
}

// LocalReusableWorkflowCache is a cache for local reusable workflow metadata files. It avoids find/read/parse
// local reusable workflow YAML files. This cache is dedicated for a single project (repository)
// indicated by 'proj' field. One LocalReusableWorkflowCache instance needs to be created per one
// project.
type LocalReusableWorkflowCache struct {
	mu    sync.RWMutex
	proj  *Project // maybe nil
	cache map[string]*ReusableWorkflowMetadata
	cwd   string
	dbg   io.Writer
}

func (c *LocalReusableWorkflowCache) debug(format string, args ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func (c *LocalReusableWorkflowCache) readCache(key string) (*ReusableWorkflowMetadata, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func (c *LocalReusableWorkflowCache) writeCache(key string, val *ReusableWorkflowMetadata) {
	_ = "STUB: not implemented"
	return
}

// FindMetadata finds/parses a reusable workflow metadata located by the 'spec' argument. When project
// is not set to 'proj' field or the spec does not start with "./", this method immediately returns with nil.
//
// Note that an error is not cached. At first search, let's say this method returned an error since
// the reusable workflow is invalid. In this case, calling this method with the same spec later will
// not return the error again. It just will return nil. This behavior prevents repeating to report
// the same error from multiple places.
//
// Calling this method is thread-safe.
func (c *LocalReusableWorkflowCache) FindMetadata(spec string) (*ReusableWorkflowMetadata, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Remember the workflow file was not found

// Remember the workflow file was invalid

func (c *LocalReusableWorkflowCache) convWorkflowPathToSpec(p string) (string, bool) {
	_ = "STUB: not implemented"
	return "", false
}

// Unreachable

// WriteWorkflowCallEvent writes reusable workflow metadata by converting from WorkflowCallEvent AST
// node. The 'wpath' parameter is a path to the workflow file of the AST, which is a relative to the
// project root directory or an absolute path.
// This method does nothing when (1) no project is set, (2) it could not convert the workflow path
// to workflow call spec, (3) some cache for the workflow is already existing.
// This method is thread safe.
func (c *LocalReusableWorkflowCache) WriteWorkflowCallEvent(wpath string, event *WorkflowCallEvent) {
	_ = "STUB: not implemented"
	// Convert workflow path to workflow call spec
	return
}

func parseReusableWorkflowMetadata(src []byte) (*ReusableWorkflowMetadata, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Unreachable

// on:
//   workflow_call:

// on: workflow_call

// on: [workflow_call]

// NewLocalReusableWorkflowCache creates a new LocalReusableWorkflowCache instance for the given
// project. 'cwd' is a current working directory as an absolute file path. The 'Local' means that
// the cache instance is project-local. It is not available across multiple projects.
func NewLocalReusableWorkflowCache(proj *Project, cwd string, dbg io.Writer) *LocalReusableWorkflowCache {
	_ = "STUB: not implemented"
	return nil
}

func newNullLocalReusableWorkflowCache(dbg io.Writer) *LocalReusableWorkflowCache {
	_ = "STUB: not implemented"
	// Null cache. Cache never hits. It is used when project is not found
	return nil
}

// LocalReusableWorkflowCacheFactory is a factory object to create a LocalReusableWorkflowCache
// instance per project.
type LocalReusableWorkflowCacheFactory struct {
	caches map[string]*LocalReusableWorkflowCache
	cwd    string
	dbg    io.Writer
}

// NewLocalReusableWorkflowCacheFactory creates a new LocalReusableWorkflowCacheFactory instance.
func NewLocalReusableWorkflowCacheFactory(cwd string, dbg io.Writer) *LocalReusableWorkflowCacheFactory {
	_ = "STUB: not implemented"
	return nil
}

// GetCache returns a new or existing LocalReusableWorkflowCache instance per project. When a instance
// was already created for the project, this method returns the existing instance. Otherwise it creates
// a new instance and returns it.
func (f *LocalReusableWorkflowCacheFactory) GetCache(p *Project) *LocalReusableWorkflowCache {
	_ = "STUB: not implemented"
	return nil
}
