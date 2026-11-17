package currency

import "encoding/xml"

type ValCurs struct {
	XMLName xml.Name `xml:"ValCurs" json:"-"`
	Date    string   `xml:"Date,attr" json:"-"`
	Name    string   `xml:"name,attr" json:"-"`

	Valutes []Valute `xml:"Valute"`
}
