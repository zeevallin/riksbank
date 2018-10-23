package swea

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"

	"github.com/zeeraw/riksbank/api/swea/requests"
)

const (
	scheme = "http"
	host   = "swea.riksbank.se"
	path   = "/sweaWS/services/SweaWebServiceHttpSoap12Endpoint"
)

func request(t *template.Template, req interface{}) (buf *bytes.Buffer, err error) {
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

func call(body io.Reader, v interface{}) (err error) {
	u := &url.URL{
		Path:   path,
		Scheme: scheme,
		Host:   host,
	}

	res, err := http.Post(u.String(), "text/xml", body)
	if err != nil {
		return
	}
	defer res.Body.Close()

	bts, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return
	}

	err = xml.Unmarshal(bts, v)
	return
}
