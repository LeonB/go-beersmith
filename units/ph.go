package units

import (
	"encoding/xml"
	"fmt"
)

type PH float64

func (ph PH) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	s := fmt.Sprintf("%.7f", ph)
	return e.EncodeElement(s, start)
}
