package main

// This is a script to generate a Go source that contains all activity types of all webhook events.
// Run the following command from the root of this repository to apply manually.
// This script is usually run via `go generate`.
// ```
// go run ./scripts/generate-webhook-events input.html all_webhooks.go
// ```

import (
	"fmt"
	"io"
	"log"
	"os"

	"golang.org/x/net/html"
)

const theURL = "https://docs.github.com/en/actions/reference/workflows-and-actions/events-that-trigger-workflows"

var dbg = log.New(io.Discard, "", log.LstdFlags)

// Parse the activity types of each webhook event. The keys of the map are names of the webhook events
// like "pull_request", and the values are arrays of names of their activity types.
// The HTML input is assumed to be fetched from the following page.
// https://docs.github.com/en/actions/reference/workflows-and-actions/events-that-trigger-workflows#pull_request
func parse(src []byte) (map[string][]string, error) { _ = "STUB: not implemented"; return nil, nil }

func findNode(root *html.Node, pred func(*html.Node) bool) *html.Node {
	_ = "STUB: not implemented"
	return nil
}

func attr(n *html.Node, name string) string { _ = "STUB: not implemented"; return "" }

func hasClass(n *html.Node, want string) bool { _ = "STUB: not implemented"; return false }

func walk(n *html.Node, visit func(*html.Node)) { _ = "STUB: not implemented"; return }

func text(n *html.Node) string { _ = "STUB: not implemented"; return "" }

func eventNameOfHeading(h *html.Node) string { _ = "STUB: not implemented"; return "" }

func children(n *html.Node, tag string) []*html.Node { _ = "STUB: not implemented"; return nil }

func firstChildByTag(n *html.Node, tag string) *html.Node { _ = "STUB: not implemented"; return nil }

func parseTable(hook string, table *html.Node) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func code(n *html.Node) []string { _ = "STUB: not implemented"; return nil }

func write(parsed map[string][]string, out io.Writer) error { _ = "STUB: not implemented"; return nil }

func fetch(url string) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func run(args []string, stdout, dbgout io.Writer, srcURL string) error {
	_ = "STUB: not implemented"
	return nil
}

func main() {
	if err := run(os.Args[1:], os.Stdout, os.Stderr, theURL); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
