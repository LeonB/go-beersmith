package units

import (
	"encoding/xml"
	"fmt"
)

type Duration float64

func (d Duration) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if d == 0.0 {
		return e.EncodeElement("0", start)
	}

	s := fmt.Sprintf("%.7f", d)
	return e.EncodeElement(s, start)
}
