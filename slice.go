package collection

import "github.com/clipperhouse/typewriter"

var collection = &typewriter.Template{
	Name: "collection",
	Text: `// not thread safe, {{.SliceName}} is a collection of type {{.Type}}. Use it where you would use []{{.Type}}.
type {{.SliceName}} []{{.Type}}
`,
}

//默认加Contain 基础函数
