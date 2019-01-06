package responses

import "encoding/xml"

// GetInterestAndExchangeNamesResponseEnvelope was generated 2019-01-06 03:21:26
type GetInterestAndExchangeNamesResponseEnvelope struct {
	XMLName xml.Name `xml:"Envelope"`
	Text    string   `xml:",chardata"`
	Soap    string   `xml:"soap,attr"`
	Body    struct {
		Text                                string `xml:",chardata"`
		GetInterestAndExchangeNamesResponse struct {
			Text   string `xml:",chardata"`
			Ns2    string `xml:"ns2,attr"`
			Return []struct {
				Text     string `xml:",chardata"`
				Datefrom struct {
					Text string `xml:",chardata"`
				} `xml:"datefrom"`
				Dateto struct {
					Text string `xml:",chardata"`
					Xsi  string `xml:"xsi,attr"`
					Nil  string `xml:"nil,attr"`
				} `xml:"dateto"`
				Description struct {
					Text string `xml:",chardata"`
				} `xml:"description"`
				Groupid struct {
					Text string `xml:",chardata"`
				} `xml:"groupid"`
				Longdescription struct {
					Text string `xml:",chardata"`
				} `xml:"longdescription"`
				Seriesid struct {
					Text string `xml:",chardata"`
				} `xml:"seriesid"`
				Shortdescription struct {
					Text string `xml:",chardata"`
				} `xml:"shortdescription"`
				Source struct {
					Text string `xml:",chardata"`
				} `xml:"source"`
				Type struct {
					Text string `xml:",chardata"`
				} `xml:"type"`
			} `xml:"return"`
		} `xml:"getInterestAndExchangeNamesResponse"`
	} `xml:"Body"`
}
