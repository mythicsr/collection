package collection

import "github.com/clipperhouse/typewriter"

var count = &typewriter.Template{
	Name: "Count",
	Text: `
// 满足 matchFn 的元素个数
func (rcv {{.SliceName}}) Count(matchFn func({{.Type}}) bool) (result int) {
	for _, v := range rcv {
		if matchFn(v) {
			result++
		}
	}
	return
}
`}
