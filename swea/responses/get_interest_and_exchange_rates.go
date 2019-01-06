package responses

import "encoding/xml"

// GetInterestAndExchangeRatesResponseEnvelope was generated 2019-01-05 23:58:09
type GetInterestAndExchangeRatesResponseEnvelope struct {
	XMLName xml.Name `xml:"Envelope"`
	Text    string   `xml:",chardata"`
	Soap    string   `xml:"soap,attr"`
	Body    struct {
		Text                                string `xml:",chardata"`
		GetInterestAndExchangeRatesResponse struct {
			Text   string `xml:",chardata"`
			Ns2    string `xml:"ns2,attr"`
			Return struct {
				Text     string `xml:",chardata"`
				Datefrom struct {
					Text string `xml:",chardata"`
				} `xml:"datefrom"`
				Dateto struct {
					Text string `xml:",chardata"`
				} `xml:"dateto"`
				Groups []struct {
					Text    string `xml:",chardata"`
					Groupid struct {
						Text string `xml:",chardata"`
					} `xml:"groupid"`
					Groupname struct {
						Text string `xml:",chardata"`
					} `xml:"groupname"`
					Series []struct {
						Text     string `xml:",chardata"`
						Seriesid struct {
							Text string `xml:",chardata"`
						} `xml:"seriesid"`
						Seriesname struct {
							Text string `xml:",chardata"`
						} `xml:"seriesname"`
						Unit struct {
							Text string `xml:",chardata"`
							Xsi  string `xml:"xsi,attr"`
							Nil  string `xml:"nil,attr"`
						} `xml:"unit"`
						Resultrows []struct {
							Text string `xml:",chardata"`
							Date struct {
								Text string `xml:",chardata"`
							} `xml:"date"`
							Period struct {
								Text string `xml:",chardata"`
							} `xml:"period"`
							Min struct {
								Text string `xml:",chardata"`
								Xsi  string `xml:"xsi,attr"`
								Nil  string `xml:"nil,attr"`
							} `xml:"min"`
							Average struct {
								Text string `xml:",chardata"`
								Xsi  string `xml:"xsi,attr"`
								Nil  string `xml:"nil,attr"`
							} `xml:"average"`
							Max struct {
								Text string `xml:",chardata"`
								Xsi  string `xml:"xsi,attr"`
								Nil  string `xml:"nil,attr"`
							} `xml:"max"`
							Ultimo struct {
								Text string `xml:",chardata"`
								Xsi  string `xml:"xsi,attr"`
								Nil  string `xml:"nil,attr"`
							} `xml:"ultimo"`
							Value struct {
								Text string `xml:",chardata"`
								Xsi  string `xml:"xsi,attr"`
								Nil  string `xml:"nil,attr"`
							} `xml:"value"`
						} `xml:"resultrows"`
					} `xml:"series"`
				} `xml:"groups"`
			} `xml:"return"`
		} `xml:"getInterestAndExchangeRatesResponse"`
	} `xml:"Body"`
}
