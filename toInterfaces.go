package collection

import "github.com/clipperhouse/typewriter"

var toInterfaces = &typewriter.Template{
	Name: "ToInterfaces",
	Text: `
func (rcv {{.SliceName}}) ToInterfaces() []interface{} {
	var result []interface{}
	for _, v := range rcv {
		result = append(result, v)
	}
	return result
}
`}
