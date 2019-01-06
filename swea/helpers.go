package swea

import (
	"bytes"
	"fmt"
	"html/template"
	"log"

	"github.com/zeeraw/riksbank/swea/requests"
)

func build(t *template.Template, req interface{}) (buf *bytes.Buffer, err error) {
	buf = new(bytes.Buffer)
	err = t.Execute(buf, req)
	if err != nil {
		return
	}
	return
}

func tmpl(name string) *template.Template {
	box := requests.Box()
	bts, err := box.Find(fmt.Sprintf("%s.xml", name))
	if err != nil {
		log.Fatal("cannot load table template")
	}
	return template.Must(template.New(name).Funcs(requests.FuncMap).Parse(string(bts)))
}
