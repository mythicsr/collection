package collection

import "github.com/clipperhouse/typewriter"

var chunk = &typewriter.Template{
	Name: "Chunk",
	Text: `// Chunk(['a', 'b', 'c', 'd'], 3) => [['a', 'b', 'c'], ['d']]
func (rcv {{.SliceName}}) Chunk(size int) [][]{{.Type}} {
	var result [][]{{.Type}}
	for i := 0; i < len(rcv); {
		s := make([]{{.Type}}, 0)
		for j := 0; j < size && i < len(rcv); j++ {
			s = append(s, rcv[i])
			i++
		}
		result = append(result, s)
	}
	return result
}
`}
