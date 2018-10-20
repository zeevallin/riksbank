package main

import (
	"bytes"
	"go/format"
	"html/template"
	"log"
	"os"
	"path/filepath"

	"github.com/zeeraw/riksbank/gen/scraper"
	"github.com/zeeraw/riksbank/gen/templates"
)

func main() {

	sections := scraper.Scrape()

	box := templates.Box()

	tbytes, err := box.MustBytes("series.go.tmpl")
	if err != nil {
		log.Fatal("cannot load table template")
	}
	t := template.Must(template.New("series").Funcs(templates.FuncMap).Parse(string(tbytes)))

	view := templates.Template{
		PackageName: "series",
		Sections:    sections,
	}
	buf := new(bytes.Buffer)
	err = t.Execute(buf, view)
	if err != nil {
		log.Fatal(err)
	}

	formatted, err := format.Source(buf.Bytes())
	if err != nil {
		log.Fatal(err)
	}

	fpath := filepath.Join(os.Args[1], "series.go")

	f, err := os.Create(fpath)
	defer f.Close()
	if err != nil {
		log.Fatal(err)
	}
	bytes.NewBuffer(formatted).WriteTo(f)
}
