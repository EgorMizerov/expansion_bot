package fleet2

import (
	"bytes"
	"text/template"
)

type URL string

func (url URL) SetPage(page int) URL {
	data, buf := struct{ Page int }{Page: page}, new(bytes.Buffer)
	template.Must(template.New("").Parse(string(url))).Execute(buf, data)
	return URL(buf.String())
}

func (url URL) String() string {
	return string(url)
}
