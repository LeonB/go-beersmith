package units

import (
	"encoding/xml"
	"fmt"
)

type Lintner float64

func (l Lintner) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	s := fmt.Sprintf("%.7f", l)
	return e.EncodeElement(s, start)
}
