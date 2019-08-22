package units

import (
	"encoding/xml"
	"fmt"
)

type Price float64

func (p Price) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	s := fmt.Sprintf("%.7f", p)
	return e.EncodeElement(s, start)
}
