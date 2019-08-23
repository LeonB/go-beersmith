package units

import (
	"encoding/xml"
	"strings"
	"time"
)

type Time struct {
	time.Time
}

func (t *Time) UnmarshalXML(dec *xml.Decoder, start xml.StartElement) error {
	var value string
	err := dec.DecodeElement(&value, &start)
	if err != nil {
		return err
	}

	value = strings.TrimSpace(value)

	if value == "" {
		return nil
	}

	layout := "2006-01-02 15:04:05"
	t2, err := time.Parse(layout, value)
	if err == nil {
		*t = Time{Time: t2}
	}
	return err
}

func (t Time) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(t.Format("2006-01-02 15:04:05"), start)
}
