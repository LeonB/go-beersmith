package units

import (
	"encoding/xml"
	"fmt"
)

type Percentage float64

func (p Percentage) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	s := fmt.Sprintf("%.7f", p)
	return e.EncodeElement(s, start)
}
