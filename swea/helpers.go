package swea

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"strings"

	"github.com/zeeraw/riksbank/swea/requests"
)

func isTrue(s string) bool {
	switch strings.ToUpper(s) {
	case "Y", "YES", "TRUE", "1":
		return true
	}
	return false
}

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
	bts, err := box.MustBytes(fmt.Sprintf("%s.xml", name))
	if err != nil {
		log.Fatal("cannot load table template")
	}
	return template.Must(template.New(name).Funcs(requests.FuncMap).Parse(string(bts)))
}
