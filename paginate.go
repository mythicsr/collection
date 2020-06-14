package collection

import "github.com/clipperhouse/typewriter"

var paginate = &typewriter.Template{
	Name: "Paginate",
	Text: `
//pageIndex begin with 0
func (rcv {{.SliceName}}) Paginate(pageIndex int, sizePerPage int) {{.SliceName}} {
	chunks := rcv.Chunk(sizePerPage)
	if 0 <= pageIndex && pageIndex <= len(chunks)-1 {
		return chunks[pageIndex]
	}
	return {{.SliceName}}{}
}
`}
