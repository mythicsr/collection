package collection

import "github.com/clipperhouse/typewriter"

var remove = &typewriter.Template{
	Name: "Remove",
	Text: `
func (rcv *{{.SliceName}}) Remove(fn func(Data) bool) ({{.SliceName}}, error) {
	_, index, err := rcv.FindOne(fn)
	if err != nil {
		return *rcv, err
	}
	*rcv = append((*rcv)[:index], (*rcv)[index+1:]...)
	return *rcv, nil
}
`}
