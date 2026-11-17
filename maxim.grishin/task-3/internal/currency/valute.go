package currency

type Valute struct {
	ID        string     `xml:"ID,attr" json:"-"`
	NumCode   int        `xml:"NumCode" json:"num_code"`
	CharCode  string     `xml:"CharCode" json:"char_code"`
	Nominal   int        `xml:"Nominal" json:"-"`
	Name      string     `xml:"Name" json:"-"`
	Value     CommaFloat `xml:"Value" json:"value"`
	VunitRate string     `xml:"VunitRate" json:"-"`
}
