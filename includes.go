package collection

import "github.com/clipperhouse/typewriter"

var includes = &typewriter.Template{
	Name: "Includes",
	Text: `
func (rcv {{.SliceName}}) Includes(equalFn func({{.Type}}, {{.Type}}) bool, targets ...{{.Type}}) bool {
    if len(targets) == 0 || len(rcv) == 0 {
        return false
    }
	if equalFn == nil {
		equalFn = func(data {{.Type}}, data2 {{.Type}}) bool {
			return reflect.DeepEqual(data, data2)
		}
	}
	_targets := {{.SliceName}}(targets)
	_objs := rcv.DistinctBy(equalFn)
	_arr := _targets.DistinctBy(equalFn)
	found := 0

	for j := 0; j < len(_objs); j++ {
		for i := 0; i < len(_arr); i++ {
			if equalFn(_objs[j], _arr[i]) {
				found++
			}
		}
	}
	return found == len(_objs)
}
`}
