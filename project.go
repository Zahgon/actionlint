package actionlint

// Project represents one GitHub project. One Git repository corresponds to one project.
type Project struct {
	root   string
	config *Config
}

func absPath(path string) string { _ = "STUB: not implemented"; return "" }

// findProject creates new Project instance by finding a project which the given path belongs to.
// A project must be a Git repository and have ".github/workflows" directory.
func findProject(path string) (*Project, error) { _ = "STUB: not implemented"; return nil, nil }

// Note: .git may be a file

// NewProject creates a new instance with a file path to the root directory of the repository.
// This function returns an error when failing to parse an actionlint config file in the repository.
func NewProject(root string) (*Project, error) { _ = "STUB: not implemented"; return nil, nil }

// RootDir returns a root directory path of the GitHub project repository.
func (p *Project) RootDir() string {
	_ = "STUB: not implemented"

	// WorkflowsDir returns a ".github/workflows" directory path of the GitHub project repository.
	// This method does not check if the directory exists.
	return ""
}

func (p *Project) WorkflowsDir() string { _ = "STUB: not implemented"; return "" }

// Knows returns true when the project knows the given file. When a file is included in the
// project's directory, the project knows the file.
func (p *Project) Knows(path string) bool {
	_ = "STUB: not implemented"
	// TODO: strings.HasPrefix is not perfect to check file path
	return false
}

// Config returns config object of the GitHub project repository. The config file was read from
// ".github/actionlint.yaml" or ".github/actionlint.yml" when this Project instance was created.
// When no config was found, this method returns nil.
func (p *Project) Config() *Config {
	_ = "STUB: not implemented"
	// Note: Calling this method must be thread safe (#333)
	return nil
}

// Projects represents set of projects. It caches Project instances which was created previously
// and reuses them.
type Projects struct {
	known []*Project
}

// NewProjects creates new Projects instance.
func NewProjects() *Projects {
	_ = "STUB: not implemented"

	// At returns the Project instance which the path belongs to. It returns nil if no project is found
	// from the path.
	return nil
}

func (ps *Projects) At(path string) (*Project, error) { _ = "STUB: not implemented"; return nil, nil }
