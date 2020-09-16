package collection

import "github.com/clipperhouse/typewriter"

var find = &typewriter.Template{
	Name: "Find",
	Text: `
func (rcv {{.SliceName}}) Find(fn func({{.Type}}) bool) (result {{.SliceName}}) {
	for _, v := range rcv {
		if fn(v) {
			result = append(result, v)
		}
	}
	return result
}
`}
