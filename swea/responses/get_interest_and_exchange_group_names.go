package responses

import "encoding/xml"

// GetInterestAndExchangeGroupNamesResponseEnvelope was generated 2019-01-06 02:21:46
type GetInterestAndExchangeGroupNamesResponseEnvelope struct {
	XMLName xml.Name `xml:"Envelope"`
	Text    string   `xml:",chardata"`
	Soap    string   `xml:"soap,attr"`
	Body    struct {
		Text                                     string `xml:",chardata"`
		GetInterestAndExchangeGroupNamesResponse struct {
			Text   string `xml:",chardata"`
			Ns2    string `xml:"ns2,attr"`
			Return []struct {
				Text             string `xml:",chardata"`
				Groupdescription struct {
					Text string `xml:",chardata"`
					Xsi  string `xml:"xsi,attr"`
					Nil  string `xml:"nil,attr"`
				} `xml:"groupdescription"`
				Groupid struct {
					Text string `xml:",chardata"`
				} `xml:"groupid"`
				Groupname struct {
					Text string `xml:",chardata"`
				} `xml:"groupname"`
				Languageid struct {
					Text string `xml:",chardata"`
				} `xml:"languageid"`
				Parentgroupid struct {
					Text string `xml:",chardata"`
					Xsi  string `xml:"xsi,attr"`
					Nil  string `xml:"nil,attr"`
				} `xml:"parentgroupid"`
			} `xml:"return"`
		} `xml:"getInterestAndExchangeGroupNamesResponse"`
	} `xml:"Body"`
}
