package collection

import "github.com/clipperhouse/typewriter"

var lpush = &typewriter.Template{
	Name: "LPUSH",
	Text: `
func (rcv *{{.SliceName}}) LPUSH(v {{.Type}}) {{.SliceName}} {
	length := len(*rcv)
	*rcv = append(*rcv, {{.Type}}{})
	for i := length; i >= 1; i-- {
		(*rcv)[i] = (*rcv)[i-1]
	}
	(*rcv)[0] = v
	return *rcv
}
`}
