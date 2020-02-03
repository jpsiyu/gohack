package ch4

import (
	"html/template"
	"os"
)

var x = `
<html>
	<body>
		Hello {{.}}
	</body>
</html>
`

func TemplateDemo() {
	t, err := template.New("Hello").Parse(x)
	if err != nil {
		panic(err)
	}
	t.Execute(os.Stdout, "<script>alert('hello')</script>")
}
