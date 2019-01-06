package requests

import (
	"fmt"
	"html/template"
	"time"

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

func fmtDate(t time.Time) string {
	return fmt.Sprintf("%04d-%02d-%02d", t.Year(), t.Month(), t.Day())
}

func fmtBool(b bool) string {
	if b {
		return "true"
	}
	return "false"
}
