package units

import (
	"encoding/xml"
	"fmt"
)

type Temperature float64

func (t Temperature) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	s := fmt.Sprintf("%.7f", t)
	return e.EncodeElement(s, start)
}
