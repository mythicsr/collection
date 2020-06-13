package collection

import "github.com/clipperhouse/typewriter"

var collection = &typewriter.Template{
	Name: "collection",
	Text: `// {{.SliceName}} is a collection of type {{.Type}}. Use it where you would use []{{.Type}}.
type {{.SliceName}} []{{.Type}}
`,
}
