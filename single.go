package collection

import "github.com/clipperhouse/typewriter"

var single = &typewriter.Template{
	Name: "Single",
	Text: `
// 是否只存在一个元素满足 matchFn
func (rcv {{.SliceName}}) Single(matchFn func({{.Type}}) bool) (result {{.Type}}, err error) {
	var candidate {{.Type}}
	found := false
	for _, v := range rcv {
		if matchFn(v) {
			if found {
				err = errors.New("multiple {{.SliceName}} elements return true for passed func")
				return
			}
			candidate = v
			found = true
		}
	}
	if found {
		result = candidate
	} else {
		err = errors.New("no {{.SliceName}} elements return true for passed func")
	}
	return
}
`}
