package units

import (
	"encoding/xml"
	"time"

	"github.com/aodin/date"
)

type Date date.Date

func (d *Date) UnmarshalXML(dec *xml.Decoder, start xml.StartElement) error {
	var value string
	err := dec.DecodeElement(&value, &start)
	if err != nil {
		return err
	}

	if value == "" {
		return nil
	}

	layout := "2006-01-02"
	d2, err := time.Parse(layout, value)
	if err == nil {
		*d = Date{Time: d2}
	}
	return err
}
