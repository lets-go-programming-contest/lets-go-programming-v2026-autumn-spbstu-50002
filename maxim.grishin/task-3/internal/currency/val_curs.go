package currency

import "encoding/xml"

type ValCurs struct {
	XMLName xml.Name `json:"-" xml:"ValCurs"`
	Date    string   `json:"-" xml:"Date,attr"`
	Name    string   `json:"-" xml:"name,attr"`

	Valutes []Valute `xml:"Valute"`
}
