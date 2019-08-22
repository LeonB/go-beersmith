package units

import (
	"encoding/xml"
	"fmt"
)

type Duration struct {
	float64
	printf string
}

func (d *Duration) SetPrintf(format string) {
	d.printf = format
}

func (d Duration) Printf() string {
	if d.printf == "" {
		return "%.7f"
	}
	return d.printf
}

func (d *Duration) UnmarshalXML(dec *xml.Decoder, start xml.StartElement) error {
	return dec.DecodeElement(&d.float64, &start)
}

func (d Duration) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	s := fmt.Sprintf(d.Printf(), d.float64)
	return e.EncodeElement(s, start)
}
