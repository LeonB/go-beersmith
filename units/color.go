package units

import (
	"encoding/xml"
	"fmt"
)

type Color float64

func (c Color) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	s := fmt.Sprintf("%.7f", c)
	return e.EncodeElement(s, start)
}

type Quantity struct {
}

// os.Exit(2)

// default unit
// conversion?
// quantity is a interface?
// quantity is struct?
