package collection

import "github.com/clipperhouse/typewriter"

var average = &typewriter.Template{
	Name: "Average",
	Text: `
func (rcv {{.SliceName}}) Average() (result float64, err error) {
	var results float64
	length := rcv.Count(func(i {{.Type}}) bool {
		return true
	})

	l := len(rcv)
	if l == 0 {
		return 0, errors.New("cannot determine Average of zero-length IntSlice")
	}
	for _, v := range rcv {
		results += float64(v)
	}

	return results / float64(length), nil
}
`,
	TypeConstraint: typewriter.Constraint{Numeric: true},
}

var averageT = &typewriter.Template{
	Name: "Average",
	Text: `
func (rcv {{.SliceName}}) Average{{.TypeParameter.LongName}}(fn func({{.Type}}) {{.TypeParameter}}) (result {{.TypeParameter}}, err error) {
	l := len(rcv)
	if l == 0 {
		err = errors.New("cannot determine Average[{{.TypeParameter}}] of zero-length {{.SliceName}}")
		return
	}
	for _, v := range rcv {
		result += fn(v)
	}
	result = result / {{.TypeParameter}}(l)
	return
}
`,
	TypeParameterConstraints: []typewriter.Constraint{
		// exactly one type parameter is required, and it must be numeric
		{Numeric: true},
	},
}
