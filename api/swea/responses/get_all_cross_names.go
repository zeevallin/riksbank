package responses

import (
	"encoding/xml"
)

// GetAllCrossNamesResponseEnvelope was generated 2018-10-23 13:30:44
type GetAllCrossNamesResponseEnvelope struct {
	XMLName xml.Name `xml:"Envelope"`
	Text    string   `xml:",chardata"`
	Soap    string   `xml:"soap,attr"`
	Body    struct {
		Text                     string `xml:",chardata"`
		GetAllCrossNamesResponse struct {
			Text   string `xml:",chardata"`
			Ns2    string `xml:"ns2,attr"`
			Return []struct {
				Text              string `xml:",chardata"`
				Seriesdescription struct {
					Text string `xml:",chardata"`
				} `xml:"seriesdescription"`
				Seriesid struct {
					Text string `xml:",chardata"`
				} `xml:"seriesid"`
				Seriesname struct {
					Text string `xml:",chardata"`
				} `xml:"seriesname"`
			} `xml:"return"`
		} `xml:"getAllCrossNamesResponse"`
	} `xml:"Body"`
}
