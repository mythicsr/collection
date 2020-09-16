package collection

import "github.com/clipperhouse/typewriter"

var removeByIndex = &typewriter.Template{
	Name: "RemoveByIndex",
	Text: `
func (rcv {{.SliceName}}) RemoveByIndex(index int) ({{.SliceName}}, error) {
	length := len(rcv)
	if !(0 <= index && index <= length-1) {
		return rcv, errors.New("out of range")
	}
	rcv = append(rcv[:index], (rcv)[index+1:]...)
	return rcv, nil
}
`}
