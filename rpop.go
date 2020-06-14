package collection

import "github.com/clipperhouse/typewriter"

var rpop = &typewriter.Template{
	Name: "RPOP",
	Text: `
func (rcv *{{.SliceName}}) RPOP() (Data, {{.SliceName}}, error) {
	var v {{.Type}}
	if len(*rcv) == 0 {
		return v, *rcv, errors.New("RPOP error, no data")
	}
	end := len(*rcv) - 1
	v = (*rcv)[end]
	*rcv = (*rcv)[:end]
	return v, *rcv, nil
}
`}
