package collection

import "github.com/clipperhouse/typewriter"

var keyByT = &typewriter.Template{
	Name: "KeyBy",
	Text: `
// KeyBy{{.TypeParameter.LongName}} groups elements into a map keyed by {{.TypeParameter}}. See: http://clipperhouse.github.io/gen/#KeyBy
func (rcv {{.SliceName}}) KeyBy{{.TypeParameter.LongName}}(fn func({{.Type}}) {{.TypeParameter}}) map[{{.TypeParameter}}]{{.Type}} {
	result := make(map[{{.TypeParameter}}]{{.Type}})
	for _, v := range rcv {
		key := fn(v)
		result[key] = v
	}
	return result
}
`,
	TypeParameterConstraints: []typewriter.Constraint{
		// exactly one type parameter is required, and it must be comparable
		{Comparable: true},
	},
}
