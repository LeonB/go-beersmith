package units

import (
	"encoding/xml"
	"fmt"
)

type Altitude float64

func (a Altitude) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	s := fmt.Sprintf("%.7f", a)
	return e.EncodeElement(s, start)
}
