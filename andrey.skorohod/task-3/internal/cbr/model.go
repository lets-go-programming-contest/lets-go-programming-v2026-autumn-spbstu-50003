package cbr

import "encoding/xml"

type Valute struct {
	NumCode  int      `json:"num_code"  xml:"NumCode"`
	CharCode string   `json:"char_code" xml:"CharCode"`
	Value    CBRFloat `json:"value"     xml:"Value"`
}

type ValCurs struct {
	XMLName xml.Name `xml:"ValCurs"`
	Valutes []Valute `xml:"Valute"`
}
