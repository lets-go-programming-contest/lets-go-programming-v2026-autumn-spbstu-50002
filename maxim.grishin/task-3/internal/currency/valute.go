package currency

type Valute struct {
	ID        string     `json:"-"         xml:"ID,attr"`
	NumCode   int        `json:"num_code"  xml:"NumCode"`
	CharCode  string     `json:"char_code" xml:"CharCode"`
	Nominal   int        `json:"-"         xml:"Nominal"`
	Name      string     `json:"-"         xml:"Name"`
	Value     CommaFloat `json:"value"     xml:"Value"`
	VunitRate string     `json:"-"         xml:"VunitRate"`
}
