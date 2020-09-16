package collection

import "github.com/clipperhouse/typewriter"

var distinctBy = &typewriter.Template{
	Name: "DistinctBy",
	Text: `
func (rcv {{.SliceName}}) DistinctBy(equalFn func({{.Type}}, {{.Type}}) bool) (result {{.SliceName}}) {
if equalFn == nil {
		equalFn = func(data {{.Type}}, data2 {{.Type}}) bool {
			return reflect.DeepEqual(data, data2)
		}
	}
Outer:
	for _, v := range rcv {
		for _, r := range result {
			if equalFn(v, r) {
				continue Outer
			}
		}
		result = append(result, v)
	}
	return result
}
`}
