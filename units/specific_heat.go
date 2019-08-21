package units

import (
	"encoding/xml"
	"fmt"
)

type SpecificHeat float64

func (sh SpecificHeat) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	s := fmt.Sprintf("%.7f", sh)
	return e.EncodeElement(s, start)
}
