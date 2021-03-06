package collection

import "github.com/clipperhouse/typewriter"

var sum = &typewriter.Template{
	Name: "Sum",
	Text: `
func (rcv {{.SliceName}}) Sum() (result {{.Type}}) {
	for _, v := range rcv {
		result += v
	}
	return
}
`,
	TypeConstraint: typewriter.Constraint{Numeric: true},
}

var sumT = &typewriter.Template{
	Name: "Sum",
	Text: `
func (rcv {{.SliceName}}) Sum{{.TypeParameter.LongName}}(fn func({{.Type}}) {{.TypeParameter}}) (result {{.TypeParameter}}) {
	for _, v := range rcv {
		result += fn(v)
	}
	return
}
`,
	TypeParameterConstraints: []typewriter.Constraint{
		// exactly one type parameter is required, and it must be numeric
		{Numeric: true},
	},
}
