//go:build wasm

package main

import (
	"syscall/js"

	"github.com/rhysd/actionlint"
)

var (
	window = js.Global().Get("window")
)

func fail(err error, when string) { _ = "STUB: not implemented"; return }

func encodeErrorAsMap(err *actionlint.Error) map[string]interface{} {
	_ = "STUB: not implemented"
	return nil
}

func lint(source string) interface{} { _ = "STUB: not implemented"; return nil }

func runActionlint(_this js.Value, args []js.Value) interface{} {
	_ = "STUB: not implemented"
	return nil
}

func main() {
	window.Set("runActionlint", js.FuncOf(runActionlint))
	window.Call("dismissLoading")
	lint(window.Call("getYamlSource").String()) // Show the first result
	select {}
}
