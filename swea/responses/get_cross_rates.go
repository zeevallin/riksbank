package responses

import (
	"encoding/xml"
)

// GetCrossRatesResponse was generated 2019-01-05 16:57:28
type GetCrossRatesResponse struct {
	XMLName xml.Name `xml:"Envelope"`
	Text    string   `xml:",chardata"`
	Soap    string   `xml:"soap,attr"`
	Body    struct {
		Text                  string `xml:",chardata"`
		GetCrossRatesResponse struct {
			Text   string `xml:",chardata"`
			Ns2    string `xml:"ns2,attr"`
			Return struct {
				Text            string `xml:",chardata"`
				Datefrom        string `xml:"datefrom"`
				Dateto          string `xml:"dateto"`
				Informationtext string `xml:"informationtext"`
				Groups          struct {
					Text      string `xml:",chardata"`
					Groupid   string `xml:"groupid"`
					Groupname string `xml:"groupname"`
					Series    []struct {
						Text       string `xml:",chardata"`
						Seriesid1  string `xml:"seriesid1"`
						Seriesid2  string `xml:"seriesid2"`
						Seriesname string `xml:"seriesname"`
						Resultrows []struct {
							Text   string `xml:",chardata"`
							Date   string `xml:"date"`
							Period struct {
								Text string `xml:",chardata"`
								Xsi  string `xml:"xsi,attr"`
								Nil  string `xml:"nil,attr"`
							} `xml:"period"`
							Average struct {
								Text string `xml:",chardata"`
								Xsi  string `xml:"xsi,attr"`
								Nil  string `xml:"nil,attr"`
							} `xml:"average"`
							Value string `xml:"value"`
						} `xml:"resultrows"`
					} `xml:"series"`
				} `xml:"groups"`
			} `xml:"return"`
		} `xml:"getCrossRatesResponse"`
	} `xml:"Body"`
}
