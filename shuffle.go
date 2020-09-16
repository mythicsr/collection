package collection

import "github.com/clipperhouse/typewriter"

var shuffle = &typewriter.Template{
	Name: "Shuffle",
	Text: `
// 洗牌，打乱
func (rcv {{.SliceName}}) Shuffle() {{.SliceName}} {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(rcv), func(i, j int) { rcv[i], rcv[j] = rcv[j], rcv[i] })
	return rcv
}
`}
