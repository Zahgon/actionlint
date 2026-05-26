package main

import (
	_ "embed"
	"io"
	"log"
	"os"

	"github.com/rhysd/actionlint"
)

// List of known outdated actions which cannot be detected from 'runs' in action.yml
var outdatedActions = []string{
	"actions/labeler@v1",
	"actions/checkout@v1",
	"actions/upload-artifact@v1",
	"actions/download-artifact@v1",
}

type actionOutput struct {
	Spec     string                     `json:"spec"`
	Meta     *actionlint.ActionMetadata `json:"metadata"`
	Outdated bool                       `json:"outdated"`
}

type registry struct {
	Slug    string   `json:"slug"`
	Path    string   `json:"path"`
	Tags    []string `json:"tags"`
	Next    string   `json:"next"`
	FileExt string   `json:"file_ext"`
	// slugs not to check inputs. Some actions allow to specify inputs which are not defined in action.yml.
	// In such cases, actionlint no longer can check the inputs, but it can still check outputs. (#16)
	SkipInputs bool `json:"skip_inputs"`
	// slugs which allows any outputs to be set. Some actions sets outputs 'dynamically'. Those outputs
	// may or may not exist. And they are not listed in action.yml metadata. actionlint cannot check
	// such outputs and fallback into allowing to set any outputs. (#18)
	SkipOutputs bool `json:"skip_outputs"`
}

func (r *registry) rawURL(tag string) string { _ = "STUB: not implemented"; return "" }

func (r *registry) githubURL(tag string) string { _ = "STUB: not implemented"; return "" }

func (r *registry) spec(tag string) string { _ = "STUB: not implemented"; return "" }

// Note: Actions used by top 1000 public repositories at GitHub sorted by number of occurrences:
// https://gist.github.com/rhysd/1db81fa80096b699b9c045f435d0cace

//go:embed popular_actions.json
var defaultPopularActionsJSON []byte

const minNodeRunnerVersion = 20

func isOutdated(spec, runs string) bool { _ = "STUB: not implemented"; return false }

type gen struct {
	stdout      io.Writer
	stderr      io.Writer
	log         *log.Logger
	rawRegistry []byte
}

func newGen(stdout, stderr, dbgout io.Writer) *gen { _ = "STUB: not implemented"; return nil }

func (g *gen) registry() ([]*registry, error) { _ = "STUB: not implemented"; return nil, nil }

func (g *gen) fetchRemote() (map[string]*actionlint.ActionMetadata, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// treosh/lighthouse-ci-action's metadata causes parse error due to typo. Ignore the error not to stop codegen
// https://github.com/treosh/lighthouse-ci-action/pull/153

// Workaround for #416.
// Once this PR is merged, remove this `if` statement and regenerate popular_actions.go.
// https://github.com/dorny/paths-filter/pull/236

// Workaround for #442.
// https://github.com/actions/download-artifact/issues/355

func (g *gen) writeJSONL(out io.Writer, actions map[string]*actionlint.ActionMetadata) error {
	_ = "STUB: not implemented"
	return nil
}

func (g *gen) writeGo(out io.Writer, actions map[string]*actionlint.ActionMetadata) error {
	_ = "STUB: not implemented"
	return nil
}

// Format the generated source with checking Go syntax

func (g *gen) readJSONL(file string) (map[string]*actionlint.ActionMetadata, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (g *gen) detectNewReleaseURLs() ([]string, error) { _ = "STUB: not implemented"; return nil, nil }

// Filter actions which have no next versions

func (g *gen) run(args []string) int { _ = "STUB: not implemented"; return 0 }

// When -h or -help

func main() {
	os.Exit(newGen(os.Stdout, os.Stderr, os.Stderr).run(os.Args))
}
