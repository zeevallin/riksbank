package templates

import (
	"fmt"
	"html/template"
	"strings"

	"github.com/gobuffalo/packr"
	"github.com/zeeraw/riksbank/gen/scraper"
)

type Template struct {
	PackageName string
	Sections    scraper.Sections
}

var FuncMap = template.FuncMap{
	"formatSeriesConstant":      formatSeriesConstant,
	"formatIotaAssignment":      formatIotaAssignment,
	"formatSeriesParameterName": formatSeriesParameterName,
	"formatGroupName":           formatGroupName,
}

// Box returns the template packer box
func Box() packr.Box {
	return packr.NewBox(".")
}

func formatSeriesConstant(series *scraper.Series) string {
	name := strings.ToUpper(series.ID)
	name = strings.Replace(name, "/", "", -1)
	return name
}

func formatIotaAssignment(series *scraper.Series, i int) string {
	if i == 0 {
		return fmt.Sprintf(" Series = iota")
	}
	return ""
}

func formatSeriesParameterName(series *scraper.Series) string {
	return fmt.Sprintf("g%d-%s", series.GroupID, series.ID)
}

var GroupNames = map[int]string{
	2:  "PrimeRates",
	3:  "OtherRates",
	5:  "StockholmInterbankOfferedRates",
	6:  "StateSecurities",
	7:  "StateObligations",
	8:  "StateFixInterests",
	9:  "PropertyObligations",
	10: "CompanyCertificates",

	97: "EuroMarket3MonthsRates",
	98: "EuroMarket6MonthsRates",

	99:  "International5YearsStateObligations",
	100: "International10YearsStateObligations",

	130: "ExchangeRates",
	138: "SpecialDrawingRights",

	12:  "TCWIndex",
	151: "KIXIndex",

	155: "FuturesExchangeRates",
}

func formatGroupName(group *scraper.Group) string {
	if n, ok := GroupNames[group.ID]; ok {
		return n
	}
	return fmt.Sprintf("Group%d", group.ID)
}
