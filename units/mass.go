package units

import (
	"encoding/xml"
	"fmt"
)

type Mass float64

func (m Mass) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	s := fmt.Sprintf("%.7f", m)
	return e.EncodeElement(s, start)
}
