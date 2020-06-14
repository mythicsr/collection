package collection

import "github.com/clipperhouse/typewriter"

var lpop = &typewriter.Template{
	Name: "LPOP",
	Text: `
func (rcv *{{.SliceName}}) LPOP() (Data, {{.SliceName}}, error) {
	var v {{.Type}}
	if len(*rcv) == 0 {
		return v, *rcv, errors.New("LPOP error, no data")
	}
	v = (*rcv)[0]
	*rcv = (*rcv)[1:]
	return v, *rcv, nil
}
`}
