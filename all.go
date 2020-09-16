package collection

import "github.com/clipperhouse/typewriter"

var isAll = &typewriter.Template{
	Name: "IsAll",
	Text: `
// 是否所有元素都满足 matchFn
func (rcv {{.SliceName}}) IsAll(matchFn func({{.Type}}) bool) bool {
	for _, v := range rcv {
		if !matchFn(v) {
			return false
		}
	}
	return true
}
`}
