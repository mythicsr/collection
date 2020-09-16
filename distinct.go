package collection

import "github.com/clipperhouse/typewriter"

var distinct = &typewriter.Template{
	Name: "Distinct",
	Text: `
func (rcv {{.SliceName}}) Distinct() (result {{.SliceName}}) {
	appended := make(map[{{.Type}}]bool)
	for _, v := range rcv {
		if !appended[v] {
			result = append(result, v)
			appended[v] = true
		}
	}
	return result
}
`,
	TypeConstraint: typewriter.Constraint{Comparable: true},
}
