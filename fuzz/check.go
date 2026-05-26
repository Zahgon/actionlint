//go:build gofuzz

package actionlint_fuzz

import "github.com/rhysd/actionlint"

func parseWorkflowPanicFree(data []byte) *actionlint.Workflow {
	_ = "STUB: not implemented"
	// Avoid Parse() panicking. It panics when go-yaml panics
	return nil
}

func FuzzCheck(data []byte) int { _ = "STUB: not implemented"; return 0 }
