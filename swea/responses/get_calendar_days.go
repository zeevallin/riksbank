package responses

import (
	"encoding/xml"
)

// GetCalendarDaysResponseEnvelope was generated 2018-10-23 13:32:59
type GetCalendarDaysResponseEnvelope struct {
	XMLName xml.Name `xml:"Envelope"`
	Text    string   `xml:",chardata"`
	Soap    string   `xml:"soap,attr"`
	Body    struct {
		Text                    string `xml:",chardata"`
		GetCalendarDaysResponse struct {
			Text   string `xml:",chardata"`
			Ns2    string `xml:"ns2,attr"`
			Return []struct {
				Text    string `xml:",chardata"`
				Bankday struct {
					Text string `xml:",chardata"`
				} `xml:"bankday"`
				Caldate struct {
					Text string `xml:",chardata"`
				} `xml:"caldate"`
				Week struct {
					Text string `xml:",chardata"`
				} `xml:"week"`
				Weekyear struct {
					Text string `xml:",chardata"`
				} `xml:"weekyear"`
			} `xml:"return"`
		} `xml:"getCalendarDaysResponse"`
	} `xml:"Body"`
}
