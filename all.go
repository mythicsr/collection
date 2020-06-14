package collection

import "github.com/clipperhouse/typewriter"

var isAll = &typewriter.Template{
	Name: "IsAll",
	Text: `
// IsAll verifies that isAll elements of {{.SliceName}} return true for the passed func. See: http://clipperhouse.github.io/gen/#IsAll
func (rcv {{.SliceName}}) IsAll(fn func({{.Type}}) bool) bool {
	for _, v := range rcv {
		if !fn(v) {
			return false
		}
	}
	return true
}
`}
