package cbr

import (
	"encoding/xml"
	"fmt"
	"strconv"
	"strings"
)

type CBRFloat float64

func (float *CBRFloat) UnmarshalXML(decod *xml.Decoder, start xml.StartElement) error {
	var str string

	if err := decod.DecodeElement(&str, &start); err != nil {
		return fmt.Errorf("decode Value element: %w", err)
	}

	str = strings.TrimSpace(strings.ReplaceAll(str, ",", "."))

	value, err := strconv.ParseFloat(str, 64)
	if err != nil {
		return fmt.Errorf("parse Value %q: %w", str, err)
	}

	*float = CBRFloat(value)

	return nil
}
