package collection

import "github.com/clipperhouse/typewriter"

var rpush = &typewriter.Template{
	Name: "RPUSH",
	Text: `
func (rcv *{{.SliceName}}) RPUSH(v {{.Type}}) {{.SliceName}} {
	*rcv = append(*rcv, v)
	return *rcv
}
`}
