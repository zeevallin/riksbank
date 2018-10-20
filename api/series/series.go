// Code generated DO NOT EDIT
package series

// Series is an identifier for an interest or exchange rate
// https://www.riksbank.se/sv/statistik/sok-rantor--valutakurser/oppet-api/serier-for-webbservices/
type Series int

func (s Series) String() string {
	return GetName(s)
}

const (
	// Riksbanksräntor

	// Riksbankens styrräntor

	// SECBDEPOEFF Inlåningsränta (1994-06-01 - )
	SECBDEPOEFF Series = iota
	// SECBLENDEFF Utlåningsränta (1994-06-01 - )
	SECBLENDEFF
	// SECBREPOEFF Reporänta (1994-06-01 - )
	SECBREPOEFF
	// SECBMARGEFF Marginalränta (1987-01-30 - 1994-05-31)
	SECBMARGEFF

	// Andra Riksbanksräntor

	// SECBREFEFF Referensränta (2002-07-01 - )
	SECBREFEFF Series = iota
	// SECBDISCEFF Diskonto (1907-11-09 - 2002-06-28)
	SECBDISCEFF

	// Svenska marknadsräntor

	// Stibor fixing

	// SEDPTNSTIBOR STIBOR T/N (Tomorrow/Next) (1998-06-30 - )
	SEDPTNSTIBOR Series = iota
	// SEDP1WSTIBOR STIBOR 1-veckas löptid (1987-09-28 - )
	SEDP1WSTIBOR
	// SEDP1MSTIBOR STIBOR 1-månads löptid (1987-01-02 - )
	SEDP1MSTIBOR
	// SEDP2MSTIBOR STIBOR 2-månaders löptid (1994-09-06 - )
	SEDP2MSTIBOR
	// SEDP3MSTIBOR STIBOR 3-månaders löptid (1987-01-02 - )
	SEDP3MSTIBOR
	// SEDP6MSTIBOR STIBOR 6-månaders löptid (1987-01-02 - )
	SEDP6MSTIBOR
	// SEDP9MSTIBOR STIBOR 9-månaders löptid (1994-09-06 - 2013-03-01)
	SEDP9MSTIBOR
	// SEDP12MSTIBOR STIBOR 12-månaders löptid (1994-09-06 - 2013-03-01)
	SEDP12MSTIBOR

	// Statsskuldväxlar

	// SETB1MBENCHC Svensk statsskuldväxel 1-månads löptid (1983-01-03 - )
	SETB1MBENCHC Series = iota
	// SETB3MBENCH Svensk statsskuldväxel 3-månaders löptid (1983-01-03 - )
	SETB3MBENCH
	// SETB6MBENCH Svensk statsskuldväxel 6-månaders löptid (1984-01-02 - )
	SETB6MBENCH
	// SETB12MBENCH Svensk statskuldsväxel 12-månaders löptid (1984-01-02 - 2011-04-21)
	SETB12MBENCH

	// Statsobligationer

	// SEGVB2YC Svensk statsobligation 2-års löptid (1987-01-07 - )
	SEGVB2YC Series = iota
	// SEGVB5YC Svensk statsobligation 5-års löptid (1985-01-02 - )
	SEGVB5YC
	// SEGVB7YC Svensk statsobligation 7-års löptid (1987-01-02 - )
	SEGVB7YC
	// SEGVB10YC Svensk statsobligation 10-års löptid (1987-01-02 - )
	SEGVB10YC

	// Statsfixräntor

	// SETB3MSTFIX Svensk statsfixränta 3-månaders löptid (1991-08-13 - )
	SETB3MSTFIX Series = iota
	// SETB6MSTFIX Svensk statsfixränta 6-månaders löptid (1990-07-02 - )
	SETB6MSTFIX
	// SEGVB5MSTFIXC Svensk statsfixränta 5-års löptid (1990-07-02 - )
	SEGVB5MSTFIXC

	// Bostadsobligationer

	// SEMB2YCACOMB Svensk bostadsobligation 2-års löptid (1994-01-03 - )
	SEMB2YCACOMB Series = iota
	// SEMB5YCACOMB Svensk bostadsobligation 5-års löptid (1986-06-16 - )
	SEMB5YCACOMB

	// Företagscertifikat

	// SECP3MK1FIX Svenskt företagscertifikat 3-månaders löptid (1989-06-01 - 2005-10-06)
	SECP3MK1FIX Series = iota
	// SECP6MK1FIX Svenskt företagscertifikat 6-månaders löptid (1991-03-14 - 2005-10-03)
	SECP6MK1FIX

	// Internationella marknadsräntor

	// Euromarknadsräntor 3-månaders löptid

	// EUDP3MUSD Euormarknadsränta USA 3-månaders löptid (1979-11-28 - )
	EUDP3MUSD Series = iota
	// EUDP3MJPY Euromarknadsränta Japan 3-månaders löptid (1979-11-29 - )
	EUDP3MJPY
	// EUDP3MGBP Euromarknadsränta Storbritannien 3-månaders löptid (1979-12-19 - )
	EUDP3MGBP
	// EUDP3MEUR Euromarknadsränta EUR 3-månaders löptid (1999-01-04 - )
	EUDP3MEUR
	// EUDP3MNOK Euromarknadsränta Norge 3-månaders löptid (1981-11-12 - )
	EUDP3MNOK
	// EUDP3MDKK Euromarknadsränta Danmark 3-månaders löptid (1981-11-12 - )
	EUDP3MDKK

	// Euromarknadsräntor 6-månaders löptid

	// EUDP6MUSD Euromarknadsränta USA 6-månaders löptid (1979-11-28 - )
	EUDP6MUSD Series = iota
	// EUDP6MJPY Euromarknadsränta Japan 6-månaders löptid (1979-11-29 - )
	EUDP6MJPY
	// EUDP6MGBP Euromarknadsränta Storbritannien 6-månaders löptid (1979-11-28 - )
	EUDP6MGBP
	// EUDP6MEUR Euromarknadsränta EUR 6-månaders löptid (1999-01-04 - )
	EUDP6MEUR
	// EUDP6MNOK Euromarknadsränta Norge 6-månaders löptid (1981-11-12 - )
	EUDP6MNOK
	// EUDP6MDKK Euromarknadsränta Danmark 6-månaders löptid (1981-11-12 - )
	EUDP6MDKK

	// Internationella statsobligationer 5-års lötid

	// USGVB5Y statsobligation USA 5-års löptid (1987-02-02 - )
	USGVB5Y Series = iota
	// JPGVB5Y statsobligation Japan 5-års löptid (1987-03-23 - )
	JPGVB5Y
	// DEGVB5Y statsobligation Tyskland 5-års löptid (1987-02-09 - )
	DEGVB5Y
	// NLGVB5Y statsobligation Holland 5-års löptid (1987-02-09 - )
	NLGVB5Y
	// FRGVB5Y statsobligation Frankrike 5-års löptid (1988-02-08 - )
	FRGVB5Y
	// GBGVB5Y statsobligation Storbritannien 5-års löptid (1987-01-02 - )
	GBGVB5Y
	// EMGVB5Y statsobligation EUR 5-års löptid (1999-01-04 - )
	EMGVB5Y

	// Internationella statsobligationer 10-års lötid

	// USGVB10Y statsobligation USA 10-års löptid (1991-01-02 - )
	USGVB10Y Series = iota
	// JPGVB10Y statsobligation Japan 10-års löptid (1987-01-05 - )
	JPGVB10Y
	// DEGVB10Y statsobligation Tyskland 10-års löptid (1987-02-09 - )
	DEGVB10Y
	// NLGVB10Y statsobligation Holland 10-års löptid (1987-02-09 - )
	NLGVB10Y
	// FRGVB10Y statsobligation Frankrike 10-års löptid (1988-02-08 - )
	FRGVB10Y
	// GBGVB10Y statsobligation Storbritannien 10-års löptid (1987-01-02 - )
	GBGVB10Y
	// EMGVB10Y statsobligation EUR 10-års löptid (1999-01-04 - )
	EMGVB10Y
	// NOGVB10Y statsobligation Norge 10-års löptid (1990-05-31 - )
	NOGVB10Y
	// DKGVB10Y statsobligation Danmark10-års löptid (1982-01-04 - )
	DKGVB10Y
	// FIGVB10Y statsobligation Finland 10-års löptid (1990-04-02 - )
	FIGVB10Y

	// Valutakurser

	// Valutor mot svenska kronor

	// SEKATSPMI ATS Österrike, shilling (1993-01-04 - 2002-02-28)
	SEKATSPMI Series = iota
	// SEKAUDPMI AUD Australien, dollar (1993-01-04 - )
	SEKAUDPMI
	// SEKBEFPMI BEF Belgien, franc (1993-01-04 - 2002-02-28)
	SEKBEFPMI
	// SEKBRLPMI BRL Brasilien, real (2005-09-06 - )
	SEKBRLPMI
	// SEKCADPMI CAD Canada, dollar (1993-01-04 - )
	SEKCADPMI
	// SEKCHFPMI CHF Schweiz, franc (1993-01-04 - )
	SEKCHFPMI
	// SEKCNYPMI CNY Kina, yuan renminbi (1994-03-01 - )
	SEKCNYPMI
	// SEKCYPPMI CYP Cypern, pund (1998-01-02 - 2007-12-28)
	SEKCYPPMI
	// SEKCZKPMI CZK Tjeckien, kronor (1998-01-01 - )
	SEKCZKPMI
	// SEKDEMPMI DEM Tyskland, mark (1993-01-04 - 2002-02-28)
	SEKDEMPMI
	// SEKDKKPMI DKK Danmark, kronor (1993-01-04 - )
	SEKDKKPMI
	// SEKEEKPMI EEK Estland, kronor (1998-01-01 - 2010-12-30)
	SEKEEKPMI
	// SEKESPPMI ESP Spanien, pesetas (1993-01-04 - 2002-02-28)
	SEKESPPMI
	// SEKEURPMI EUR Euroland, euro (1993-01-04 - )
	SEKEURPMI
	// SEKFIMPMI FIM Finland, mark (1993-01-04 - 2002-02-28)
	SEKFIMPMI
	// SEKFRFPMI FRF Frankrike, franc (1993-01-04 - 2002-02-28)
	SEKFRFPMI
	// SEKGBPPMI GBP Storbritannien, pund (1993-01-04 - )
	SEKGBPPMI
	// SEKGRDPMI GRD Grekland, drachmer (1993-01-04 - 2002-02-28)
	SEKGRDPMI
	// SEKHKDPMI HKD Hong Kong, dollar (1994-03-01 - )
	SEKHKDPMI
	// SEKHUFPMI HUF Ungern, forint (1998-01-01 - )
	SEKHUFPMI
	// SEKIDRPMI IDR Indonesien, rupee (1998-01-01 - )
	SEKIDRPMI
	// SEKIEPPMI IEP Irland, pund (1993-01-04 - 2002-02-28)
	SEKIEPPMI
	// SEKINRPMI INR Indien, rupee (1994-03-01 - )
	SEKINRPMI
	// SEKISKPMI ISK Island, kronor (1993-01-04 - )
	SEKISKPMI
	// SEKITLPMI ITL Italien, lire (1993-01-04 - 2002-02-28)
	SEKITLPMI
	// SEKJPYPMI JPY Japan, yen (1993-01-04 - )
	SEKJPYPMI
	// SEKKRWPMI KRW Syd Korea, won (2005-09-09 - )
	SEKKRWPMI
	// SEKKWDPMI KWD Kuwait, dinar (1994-03-01 - 2005-02-28)
	SEKKWDPMI
	// SEKLTLPMI LTL Litauen, litas (1998-01-01 - 2014-12-30)
	SEKLTLPMI
	// SEKLVLPMI LVL Lettland, lav (1998-01-01 - 2013-12-30)
	SEKLVLPMI
	// SEKMADPMI MAD Marocko, dirham (1998-01-01 - )
	SEKMADPMI
	// SEKMXNPMI MXN Mexiko, nuevo peso (1998-01-01 - )
	SEKMXNPMI
	// SEKMYRPMI MYR Malaysia, ringitt (1994-03-01 - 2005-02-28)
	SEKMYRPMI
	// SEKNLGPMI NLG Nederländerna, gulden (1993-01-04 - 2002-02-28)
	SEKNLGPMI
	// SEKNOKPMI NOK Norge, kronor (1993-01-04 - )
	SEKNOKPMI
	// SEKNZDPMI NZD Nya Zeeland, dollar (1994-03-01 - )
	SEKNZDPMI
	// SEKPLNPMI PLN Polen, zloty (1998-01-01 - )
	SEKPLNPMI
	// SEKPTEPMI PTE Portugal, escudo (1993-01-04 - 2002-02-28)
	SEKPTEPMI
	// SEKRUBPMI RUB Ryssland, rubel (1993-01-04 - )
	SEKRUBPMI
	// SEKSARPMI SAR Saudiarabien, riyal (1994-03-04 - )
	SEKSARPMI
	// SEKSGDPMI SGD Singapore, dollar (1994-03-01 - )
	SEKSGDPMI
	// SEKSITPMI SIT Slovenien, tolar (1998-01-01 - 2006-12-29)
	SEKSITPMI
	// SEKSKKPMI SKK Slovakien, koruna (2005-04-29 - 2009-01-01)
	SEKSKKPMI
	// SEKTHBPMI THB Thailand, baht (1998-01-01 - )
	SEKTHBPMI
	// SEKTRLPMI TRL Turkiet, lira (1998-01-01 - 2004-12-30)
	SEKTRLPMI
	// SEKTRYPMI TRY Turkiet, ny lira (2005-01-03 - )
	SEKTRYPMI
	// SEKUSDPMI USD Förenta Staterna, dollar (1993-01-04 - )
	SEKUSDPMI
	// SEKZARPMI ZAR Sydafrika, rand (1994-03-01 - )
	SEKZARPMI

	// SDR särskilda dragningsrätter

	// SEKSDR SDR Särskilda dragningsrätter (SEK/SDR) (1993-01-04 - )
	SEKSDR Series = iota

	// Svenskt TCW-index

	// SEKTCW92 Svenskt TCW-index (1981-09-21 - )
	SEKTCW92 Series = iota

	// Svenskt KIX-index

	// SEKKIX92 Svenskt KIX-index (1992-11-18 - )
	SEKKIX92 Series = iota

	// Valutaterminskurser

	// SEKUSDFO3MFIX SEK/USD valutatermin 3-månaders löptid (1980-01-02 - )
	SEKUSDFO3MFIX Series = iota
	// SEKUSDFO6MFIX SEK/USD valutatermin 6-månaders löptid (1980-01-02 - )
	SEKUSDFO6MFIX
)

// Names are the parameter names of every series
var Names = map[Series]string{
	SECBDEPOEFF:   "g2-SECBDEPOEFF",
	SECBLENDEFF:   "g2-SECBLENDEFF",
	SECBREPOEFF:   "g2-SECBREPOEFF",
	SECBMARGEFF:   "g2-SECBMARGEFF",
	SECBREFEFF:    "g3-SECBREFEFF",
	SECBDISCEFF:   "g3-SECBDISCEFF",
	SEDPTNSTIBOR:  "g5-SEDPT/NSTIBOR",
	SEDP1WSTIBOR:  "g5-SEDP1WSTIBOR",
	SEDP1MSTIBOR:  "g5-SEDP1MSTIBOR",
	SEDP2MSTIBOR:  "g5-SEDP2MSTIBOR",
	SEDP3MSTIBOR:  "g5-SEDP3MSTIBOR",
	SEDP6MSTIBOR:  "g5-SEDP6MSTIBOR",
	SEDP9MSTIBOR:  "g5-SEDP9MSTIBOR",
	SEDP12MSTIBOR: "g5-SEDP12MSTIBOR",
	SETB1MBENCHC:  "g6-SETB1MBENCHC",
	SETB3MBENCH:   "g6-SETB3MBENCH",
	SETB6MBENCH:   "g6-SETB6MBENCH",
	SETB12MBENCH:  "g6-SETB12MBENCH",
	SEGVB2YC:      "g7-SEGVB2YC",
	SEGVB5YC:      "g7-SEGVB5YC",
	SEGVB7YC:      "g7-SEGVB7YC",
	SEGVB10YC:     "g7-SEGVB10YC",
	SETB3MSTFIX:   "g8-SETB3MSTFIX",
	SETB6MSTFIX:   "g8-SETB6MSTFIX",
	SEGVB5MSTFIXC: "g8-SEGVB5MSTFIXC",
	SEMB2YCACOMB:  "g9-SEMB2YCACOMB",
	SEMB5YCACOMB:  "g9-SEMB5YCACOMB",
	SECP3MK1FIX:   "g10-SECP3MK1FIX",
	SECP6MK1FIX:   "g10-SECP6MK1FIX",
	EUDP3MUSD:     "g97-EUDP3MUSD",
	EUDP3MJPY:     "g97-EUDP3MJPY",
	EUDP3MGBP:     "g97-EUDP3MGBP",
	EUDP3MEUR:     "g97-EUDP3MEUR",
	EUDP3MNOK:     "g97-EUDP3MNOK",
	EUDP3MDKK:     "g97-EUDP3MDKK",
	EUDP6MUSD:     "g98-EUDP6MUSD",
	EUDP6MJPY:     "g98-EUDP6MJPY",
	EUDP6MGBP:     "g98-EUDP6MGBP",
	EUDP6MEUR:     "g98-EUDP6MEUR",
	EUDP6MNOK:     "g98-EUDP6MNOK",
	EUDP6MDKK:     "g98-EUDP6MDKK",
	USGVB5Y:       "g99-USGVB5Y",
	JPGVB5Y:       "g99-JPGVB5Y",
	DEGVB5Y:       "g99-DEGVB5Y",
	NLGVB5Y:       "g99-NLGVB5Y",
	FRGVB5Y:       "g99-FRGVB5Y",
	GBGVB5Y:       "g99-GBGVB5Y",
	EMGVB5Y:       "g99-EMGVB5Y",
	USGVB10Y:      "g100-USGVB10Y",
	JPGVB10Y:      "g100-JPGVB10Y",
	DEGVB10Y:      "g100-DEGVB10Y",
	NLGVB10Y:      "g100-NLGVB10Y",
	FRGVB10Y:      "g100-FRGVB10Y",
	GBGVB10Y:      "g100-GBGVB10Y",
	EMGVB10Y:      "g100-EMGVB10Y",
	NOGVB10Y:      "g100-NOGVB10Y",
	DKGVB10Y:      "g100-DKGVB10Y",
	FIGVB10Y:      "g100-FIGVB10Y",
	SEKATSPMI:     "g130-SEKATSPMI",
	SEKAUDPMI:     "g130-SEKAUDPMI",
	SEKBEFPMI:     "g130-SEKBEFPMI",
	SEKBRLPMI:     "g130-SEKBRLPMI",
	SEKCADPMI:     "g130-SEKCADPMI",
	SEKCHFPMI:     "g130-SEKCHFPMI",
	SEKCNYPMI:     "g130-SEKCNYPMI",
	SEKCYPPMI:     "g130-SEKCYPPMI",
	SEKCZKPMI:     "g130-SEKCZKPMI",
	SEKDEMPMI:     "g130-SEKDEMPMI",
	SEKDKKPMI:     "g130-SEKDKKPMI",
	SEKEEKPMI:     "g130-SEKEEKPMI",
	SEKESPPMI:     "g130-SEKESPPMI",
	SEKEURPMI:     "g130-SEKEURPMI",
	SEKFIMPMI:     "g130-SEKFIMPMI",
	SEKFRFPMI:     "g130-SEKFRFPMI",
	SEKGBPPMI:     "g130-SEKGBPPMI",
	SEKGRDPMI:     "g130-SEKGRDPMI",
	SEKHKDPMI:     "g130-SEKHKDPMI",
	SEKHUFPMI:     "g130-SEKHUFPMI",
	SEKIDRPMI:     "g130-SEKIDRPMI",
	SEKIEPPMI:     "g130-SEKIEPPMI",
	SEKINRPMI:     "g130-SEKINRPMI",
	SEKISKPMI:     "g130-SEKISKPMI",
	SEKITLPMI:     "g130-SEKITLPMI",
	SEKJPYPMI:     "g130-SEKJPYPMI",
	SEKKRWPMI:     "g130-SEKKRWPMI",
	SEKKWDPMI:     "g130-SEKKWDPMI",
	SEKLTLPMI:     "g130-SEKLTLPMI",
	SEKLVLPMI:     "g130-SEKLVLPMI",
	SEKMADPMI:     "g130-SEKMADPMI",
	SEKMXNPMI:     "g130-SEKMXNPMI",
	SEKMYRPMI:     "g130-SEKMYRPMI",
	SEKNLGPMI:     "g130-SEKNLGPMI",
	SEKNOKPMI:     "g130-SEKNOKPMI",
	SEKNZDPMI:     "g130-SEKNZDPMI",
	SEKPLNPMI:     "g130-SEKPLNPMI",
	SEKPTEPMI:     "g130-SEKPTEPMI",
	SEKRUBPMI:     "g130-SEKRUBPMI",
	SEKSARPMI:     "g130-SEKSARPMI",
	SEKSGDPMI:     "g130-SEKSGDPMI",
	SEKSITPMI:     "g130-SEKSITPMI",
	SEKSKKPMI:     "g130-SEKSKKPMI",
	SEKTHBPMI:     "g130-SEKTHBPMI",
	SEKTRLPMI:     "g130-SEKTRLPMI",
	SEKTRYPMI:     "g130-SEKTRYPMI",
	SEKUSDPMI:     "g130-SEKUSDPMI",
	SEKZARPMI:     "g130-SEKZARPMI",
	SEKSDR:        "g138-SEKSDR",
	SEKTCW92:      "g12-SEKTCW92",
	SEKKIX92:      "g151-SEKKIX92",
	SEKUSDFO3MFIX: "g155-SEKUSDFO3MFIX",
	SEKUSDFO6MFIX: "g155-SEKUSDFO6MFIX",
}

// GetName will get the parameter name of a series
func GetName(s Series) string {
	return Names[s]
}

// PrimeRates Riksbankens styrräntor
var PrimeRates = map[Series]struct{}{
	SECBDEPOEFF: struct{}{},
	SECBLENDEFF: struct{}{},
	SECBREPOEFF: struct{}{},
	SECBMARGEFF: struct{}{},
}

// PrimeRatesNames returns the names of all series in the group
func PrimeRatesNames() []string {
	names := make([]string, len(PrimeRates))
	var i = 0
	for k, _ := range PrimeRates {
		names[i] = GetName(k)
		i++
	}
	return names
}

// OtherRates Andra Riksbanksräntor
var OtherRates = map[Series]struct{}{
	SECBREFEFF:  struct{}{},
	SECBDISCEFF: struct{}{},
}

// OtherRatesNames returns the names of all series in the group
func OtherRatesNames() []string {
	names := make([]string, len(OtherRates))
	var i = 0
	for k, _ := range OtherRates {
		names[i] = GetName(k)
		i++
	}
	return names
}

// StockholmInterbankOfferedRates Stibor fixing
var StockholmInterbankOfferedRates = map[Series]struct{}{
	SEDPTNSTIBOR:  struct{}{},
	SEDP1WSTIBOR:  struct{}{},
	SEDP1MSTIBOR:  struct{}{},
	SEDP2MSTIBOR:  struct{}{},
	SEDP3MSTIBOR:  struct{}{},
	SEDP6MSTIBOR:  struct{}{},
	SEDP9MSTIBOR:  struct{}{},
	SEDP12MSTIBOR: struct{}{},
}

// StockholmInterbankOfferedRatesNames returns the names of all series in the group
func StockholmInterbankOfferedRatesNames() []string {
	names := make([]string, len(StockholmInterbankOfferedRates))
	var i = 0
	for k, _ := range StockholmInterbankOfferedRates {
		names[i] = GetName(k)
		i++
	}
	return names
}

// StateSecurities Statsskuldväxlar
var StateSecurities = map[Series]struct{}{
	SETB1MBENCHC: struct{}{},
	SETB3MBENCH:  struct{}{},
	SETB6MBENCH:  struct{}{},
	SETB12MBENCH: struct{}{},
}

// StateSecuritiesNames returns the names of all series in the group
func StateSecuritiesNames() []string {
	names := make([]string, len(StateSecurities))
	var i = 0
	for k, _ := range StateSecurities {
		names[i] = GetName(k)
		i++
	}
	return names
}

// StateObligations Statsobligationer
var StateObligations = map[Series]struct{}{
	SEGVB2YC:  struct{}{},
	SEGVB5YC:  struct{}{},
	SEGVB7YC:  struct{}{},
	SEGVB10YC: struct{}{},
}

// StateObligationsNames returns the names of all series in the group
func StateObligationsNames() []string {
	names := make([]string, len(StateObligations))
	var i = 0
	for k, _ := range StateObligations {
		names[i] = GetName(k)
		i++
	}
	return names
}

// StateFixInterests Statsfixräntor
var StateFixInterests = map[Series]struct{}{
	SETB3MSTFIX:   struct{}{},
	SETB6MSTFIX:   struct{}{},
	SEGVB5MSTFIXC: struct{}{},
}

// StateFixInterestsNames returns the names of all series in the group
func StateFixInterestsNames() []string {
	names := make([]string, len(StateFixInterests))
	var i = 0
	for k, _ := range StateFixInterests {
		names[i] = GetName(k)
		i++
	}
	return names
}

// PropertyObligations Bostadsobligationer
var PropertyObligations = map[Series]struct{}{
	SEMB2YCACOMB: struct{}{},
	SEMB5YCACOMB: struct{}{},
}

// PropertyObligationsNames returns the names of all series in the group
func PropertyObligationsNames() []string {
	names := make([]string, len(PropertyObligations))
	var i = 0
	for k, _ := range PropertyObligations {
		names[i] = GetName(k)
		i++
	}
	return names
}

// CompanyCertificates Företagscertifikat
var CompanyCertificates = map[Series]struct{}{
	SECP3MK1FIX: struct{}{},
	SECP6MK1FIX: struct{}{},
}

// CompanyCertificatesNames returns the names of all series in the group
func CompanyCertificatesNames() []string {
	names := make([]string, len(CompanyCertificates))
	var i = 0
	for k, _ := range CompanyCertificates {
		names[i] = GetName(k)
		i++
	}
	return names
}

// EuroMarket3MonthsRates Euromarknadsräntor 3-månaders löptid
var EuroMarket3MonthsRates = map[Series]struct{}{
	EUDP3MUSD: struct{}{},
	EUDP3MJPY: struct{}{},
	EUDP3MGBP: struct{}{},
	EUDP3MEUR: struct{}{},
	EUDP3MNOK: struct{}{},
	EUDP3MDKK: struct{}{},
}

// EuroMarket3MonthsRatesNames returns the names of all series in the group
func EuroMarket3MonthsRatesNames() []string {
	names := make([]string, len(EuroMarket3MonthsRates))
	var i = 0
	for k, _ := range EuroMarket3MonthsRates {
		names[i] = GetName(k)
		i++
	}
	return names
}

// EuroMarket6MonthsRates Euromarknadsräntor 6-månaders löptid
var EuroMarket6MonthsRates = map[Series]struct{}{
	EUDP6MUSD: struct{}{},
	EUDP6MJPY: struct{}{},
	EUDP6MGBP: struct{}{},
	EUDP6MEUR: struct{}{},
	EUDP6MNOK: struct{}{},
	EUDP6MDKK: struct{}{},
}

// EuroMarket6MonthsRatesNames returns the names of all series in the group
func EuroMarket6MonthsRatesNames() []string {
	names := make([]string, len(EuroMarket6MonthsRates))
	var i = 0
	for k, _ := range EuroMarket6MonthsRates {
		names[i] = GetName(k)
		i++
	}
	return names
}

// International5YearsStateObligations Internationella statsobligationer 5-års lötid
var International5YearsStateObligations = map[Series]struct{}{
	USGVB5Y: struct{}{},
	JPGVB5Y: struct{}{},
	DEGVB5Y: struct{}{},
	NLGVB5Y: struct{}{},
	FRGVB5Y: struct{}{},
	GBGVB5Y: struct{}{},
	EMGVB5Y: struct{}{},
}

// International5YearsStateObligationsNames returns the names of all series in the group
func International5YearsStateObligationsNames() []string {
	names := make([]string, len(International5YearsStateObligations))
	var i = 0
	for k, _ := range International5YearsStateObligations {
		names[i] = GetName(k)
		i++
	}
	return names
}

// International10YearsStateObligations Internationella statsobligationer 10-års lötid
var International10YearsStateObligations = map[Series]struct{}{
	USGVB10Y: struct{}{},
	JPGVB10Y: struct{}{},
	DEGVB10Y: struct{}{},
	NLGVB10Y: struct{}{},
	FRGVB10Y: struct{}{},
	GBGVB10Y: struct{}{},
	EMGVB10Y: struct{}{},
	NOGVB10Y: struct{}{},
	DKGVB10Y: struct{}{},
	FIGVB10Y: struct{}{},
}

// International10YearsStateObligationsNames returns the names of all series in the group
func International10YearsStateObligationsNames() []string {
	names := make([]string, len(International10YearsStateObligations))
	var i = 0
	for k, _ := range International10YearsStateObligations {
		names[i] = GetName(k)
		i++
	}
	return names
}

// ExchangeRates Valutor mot svenska kronor
var ExchangeRates = map[Series]struct{}{
	SEKATSPMI: struct{}{},
	SEKAUDPMI: struct{}{},
	SEKBEFPMI: struct{}{},
	SEKBRLPMI: struct{}{},
	SEKCADPMI: struct{}{},
	SEKCHFPMI: struct{}{},
	SEKCNYPMI: struct{}{},
	SEKCYPPMI: struct{}{},
	SEKCZKPMI: struct{}{},
	SEKDEMPMI: struct{}{},
	SEKDKKPMI: struct{}{},
	SEKEEKPMI: struct{}{},
	SEKESPPMI: struct{}{},
	SEKEURPMI: struct{}{},
	SEKFIMPMI: struct{}{},
	SEKFRFPMI: struct{}{},
	SEKGBPPMI: struct{}{},
	SEKGRDPMI: struct{}{},
	SEKHKDPMI: struct{}{},
	SEKHUFPMI: struct{}{},
	SEKIDRPMI: struct{}{},
	SEKIEPPMI: struct{}{},
	SEKINRPMI: struct{}{},
	SEKISKPMI: struct{}{},
	SEKITLPMI: struct{}{},
	SEKJPYPMI: struct{}{},
	SEKKRWPMI: struct{}{},
	SEKKWDPMI: struct{}{},
	SEKLTLPMI: struct{}{},
	SEKLVLPMI: struct{}{},
	SEKMADPMI: struct{}{},
	SEKMXNPMI: struct{}{},
	SEKMYRPMI: struct{}{},
	SEKNLGPMI: struct{}{},
	SEKNOKPMI: struct{}{},
	SEKNZDPMI: struct{}{},
	SEKPLNPMI: struct{}{},
	SEKPTEPMI: struct{}{},
	SEKRUBPMI: struct{}{},
	SEKSARPMI: struct{}{},
	SEKSGDPMI: struct{}{},
	SEKSITPMI: struct{}{},
	SEKSKKPMI: struct{}{},
	SEKTHBPMI: struct{}{},
	SEKTRLPMI: struct{}{},
	SEKTRYPMI: struct{}{},
	SEKUSDPMI: struct{}{},
	SEKZARPMI: struct{}{},
}

// ExchangeRatesNames returns the names of all series in the group
func ExchangeRatesNames() []string {
	names := make([]string, len(ExchangeRates))
	var i = 0
	for k, _ := range ExchangeRates {
		names[i] = GetName(k)
		i++
	}
	return names
}

// SpecialDrawingRights SDR särskilda dragningsrätter
var SpecialDrawingRights = map[Series]struct{}{
	SEKSDR: struct{}{},
}

// SpecialDrawingRightsNames returns the names of all series in the group
func SpecialDrawingRightsNames() []string {
	names := make([]string, len(SpecialDrawingRights))
	var i = 0
	for k, _ := range SpecialDrawingRights {
		names[i] = GetName(k)
		i++
	}
	return names
}

// TCWIndex Svenskt TCW-index
var TCWIndex = map[Series]struct{}{
	SEKTCW92: struct{}{},
}

// TCWIndexNames returns the names of all series in the group
func TCWIndexNames() []string {
	names := make([]string, len(TCWIndex))
	var i = 0
	for k, _ := range TCWIndex {
		names[i] = GetName(k)
		i++
	}
	return names
}

// KIXIndex Svenskt KIX-index
var KIXIndex = map[Series]struct{}{
	SEKKIX92: struct{}{},
}

// KIXIndexNames returns the names of all series in the group
func KIXIndexNames() []string {
	names := make([]string, len(KIXIndex))
	var i = 0
	for k, _ := range KIXIndex {
		names[i] = GetName(k)
		i++
	}
	return names
}

// FuturesExchangeRates Valutaterminskurser
var FuturesExchangeRates = map[Series]struct{}{
	SEKUSDFO3MFIX: struct{}{},
	SEKUSDFO6MFIX: struct{}{},
}

// FuturesExchangeRatesNames returns the names of all series in the group
func FuturesExchangeRatesNames() []string {
	names := make([]string, len(FuturesExchangeRates))
	var i = 0
	for k, _ := range FuturesExchangeRates {
		names[i] = GetName(k)
		i++
	}
	return names
}
