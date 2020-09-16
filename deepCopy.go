package collection

import "github.com/clipperhouse/typewriter"

var deepCopy = &typewriter.Template{
	Name: "DeepCopy",
	Text: `
func (rcv {{.SliceName}}) DeepCopy() {{.SliceName}} {
	return deepcopy.Copy(rcv).({{.SliceName}})
}
`}
