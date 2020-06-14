package collection

import "github.com/clipperhouse/typewriter"

var reduceT = &typewriter.Template{
	Name: "Reduce",
	Text: `
// Reduce{{.TypeParameter.LongName}} iterates over {{.SliceName}}, operating on forEach element while maintaining ‘state’. fn = func(pre, {{.Type}}) int, See: http://clipperhouse.github.io/gen/#Reduce
func (rcv {{.SliceName}}) Reduce{{.TypeParameter.LongName}}(fn func({{.TypeParameter}}, {{.Type}}) {{.TypeParameter}}) (result {{.TypeParameter}}) {
	for _, v := range rcv {
		result = fn(result, v)
	}
	return
}
`,
	TypeParameterConstraints: []typewriter.Constraint{
		// exactly one type parameter is required, but no constraints on that type
		{},
	},
}
