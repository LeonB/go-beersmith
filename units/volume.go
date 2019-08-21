package units

import (
	"encoding/xml"
	"fmt"
)

type Volume float64

func (v Volume) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	s := fmt.Sprintf("%.7f", v)
	return e.EncodeElement(s, start)
}

type Volume2 Volume

func (v Volume2) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(float64(v), start)
}
