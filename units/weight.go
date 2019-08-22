package units

import (
	"encoding/xml"
	"fmt"
)

type Weight float64

func (w Weight) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	s := fmt.Sprintf("%.7f", w)
	return e.EncodeElement(s, start)
}
