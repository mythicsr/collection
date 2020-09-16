package collection

import "github.com/clipperhouse/typewriter"

var removeAll = &typewriter.Template{
	Name: "RemoveAll",
	Text: `
func (rcv {{.SliceName}}) RemoveAll(equalFn func({{.Type}}) bool) {{.SliceName}} {
	var out {{.SliceName}}
	for _, v := range rcv {
		if !equalFn(v) {
			out = append(out, v)
		}
	}
	return out
}
`}
