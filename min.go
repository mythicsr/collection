package collection

import "github.com/clipperhouse/typewriter"

var min = &typewriter.Template{
	Name: "Min",
	Text: `
	func (rcv {{.SliceName}}) Min() (result {{.Type}}, err error) {
		l := len(rcv)
		if l == 0 {
			err = errors.New("cannot determine the Min of an empty collection")
			return
		}
		result = rcv[0]
		for _, v := range rcv {
			if v < result {
				result = v
			}
		}
		return
	}
	`,
	TypeConstraint: typewriter.Constraint{Ordered: true},
}

var minT = &typewriter.Template{
	Name: "Min",
	Text: `
// Min{{.TypeParameter.LongName}} selects the least value of {{.TypeParameter}} in {{.SliceName}}. Returns error on {{.SliceName}} with no elements. See: http://clipperhouse.github.io/gen/#MinCustom
func (rcv {{.SliceName}}) Min{{.TypeParameter.LongName}}(fn func({{.Type}}) {{.TypeParameter}}) (result {{.TypeParameter}}, err error) {
	l := len(rcv)
	if l == 0 {
		err = errors.New("cannot determine Min of zero-length {{.SliceName}}")
		return
	}
	result = fn(rcv[0])
	if l > 1 {
		for _, v := range rcv[1:] {
			f := fn(v)
			if f < result {
				result = f
			}
		}
	}
	return
}
`,
	TypeParameterConstraints: []typewriter.Constraint{
		// exactly one type parameter is required, and it must be ordered
		{Ordered: true},
	},
}
