package requests

import (
	"fmt"
	"html/template"

	"cloud.google.com/go/civil"
	"github.com/gobuffalo/packr"
)

// Box returns the template packer box
func Box() packr.Box {
	return packr.NewBox(".")
}

// FuncMap are the formatting functions for packr
var FuncMap = template.FuncMap{
	"fmtDate": fmtDate,
	"fmtBool": fmtBool,
}

func fmtDate(d civil.Date) string {
	return fmt.Sprintf("%04d-%02d-%02d", d.Year, d.Month, d.Day)
}

func fmtBool(b bool) string {
	if b {
		return "true"
	}
	return "false"
}
