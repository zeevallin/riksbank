package swea

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"strings"

	"cloud.google.com/go/civil"
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
	bts, err := box.Find(fmt.Sprintf("%s.xml", name))
	if err != nil {
		log.Fatal("cannot load table template")
	}
	return template.Must(template.New(name).Funcs(requests.FuncMap).Parse(string(bts)))
}

func parseDate(s string) *civil.Date {
	d, err := civil.ParseDate(strings.TrimSpace(s))
	if err != nil {
		return &civil.Date{}
	}
	return &d
}

func parseDatePeriod(d, p string) (civil.Date, string, error) {
	date, err := civil.ParseDate(strings.TrimSpace(d))
	if err != nil {
		return date, "", err
	}
	var period string
	ptx := strings.TrimSpace(p)
	if ptx != "" {
		period = ptx
	} else {
		period = date.String()
	}

	return date, period, nil
}
