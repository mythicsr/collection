package collection

import "github.com/clipperhouse/typewriter"

var forEach = &typewriter.Template{
	Name: "ForEach",
	Text: `
// ForEach iterates over {{.SliceName}} and executes the passed func against forEach element. See: http://clipperhouse.github.io/gen/#ForEach
func (rcv {{.SliceName}}) ForEach(fn func({{.Type}})) {
	for _, v := range rcv {
		fn(v)
	}
}
`}
