package collection

import "github.com/clipperhouse/typewriter"

var isAny = &typewriter.Template{
	Name: "IsAny",
	Text: `
// IsAny verifies that one or more elements of {{.SliceName}} return true for the passed func. See: http://clipperhouse.github.io/gen/#IsAny
func (rcv {{.SliceName}}) IsAny(fn func({{.Type}}) bool) bool {
	for _, v := range rcv {
		if fn(v) {
			return true
		}
	}
	return false
}
`}
