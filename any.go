package collection

import "github.com/clipperhouse/typewriter"

var isAny = &typewriter.Template{
	Name: "IsAny",
	Text: `
// 是否存在元素满足 matchFn
func (rcv {{.SliceName}}) IsAny(matchFn func({{.Type}}) bool) bool {
	for _, v := range rcv {
		if matchFn(v) {
			return true
		}
	}
	return false
}
`}
