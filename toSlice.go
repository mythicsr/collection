package collection

import "github.com/clipperhouse/typewriter"

var toSlice = &typewriter.Template{
	Name: "ToSlice",
	Text: `
func (rcv {{.SliceName}}) ToSlice() []{{.Type}} {
	return []{{.Type}}(rcv)
}
`}
