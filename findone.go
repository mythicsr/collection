package collection

import "github.com/clipperhouse/typewriter"

var findOne = &typewriter.Template{
	Name: "FindOne",
	Text: `
func (rcv {{.SliceName}}) FindOne(fn func({{.Type}}) bool) (result {{.Type}}, index int, err error) {
	for i, v := range rcv {
		if fn(v) {
			index = i
			result = v
			return
		}
	}
	err = errors.New("no {{.SliceName}} elements return true for passed func")
	return
}
`}
