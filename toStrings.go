package collection

import "github.com/clipperhouse/typewriter"

var toStrings = &typewriter.Template{
	Name: "ToStrings",
	Text: `
func (rcv {{.SliceName}}) ToStrings() []string {
	var out []string
	for _, v := range rcv {
		out = append(out, cast.ToString(v))
	}
	return out
}
`}
