package bsmx

import (
	"encoding/xml"
	"errors"
	"fmt"
	"log"
	"strings"
)

type IntBool bool

func (b *IntBool) UnmarshalXML(dec *xml.Decoder, start xml.StartElement) error {
	var value string
	err := dec.DecodeElement(&value, &start)
	if err != nil {
		return err
	}

	value = strings.TrimSpace(value)

	if value == "0" {
		*b = false
	} else if value == "1" {
		*b = true
	} else {
		return errors.New(fmt.Sprintf("'%s' is not a valid IntBool value", value))
	}

	return nil
}

func (b IntBool) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	// if d.IsZero() {
	// 	return nil
	// }
	if b {
		return e.EncodeElement("1", start)
	} else {
		return e.EncodeElement("0", start)
	}
}

type String string

func (s String) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	// v := html.EscapeString(string(s))
	v := string(s)
	err := e.EncodeElement([]byte(v), start)
	if err != nil {
		return err
	}
	return nil
}

func (s *String) UnmarshalXML(dec *xml.Decoder, start xml.StartElement) error {
	var value string
	dec.Entity = map[string]string{}
	// dec.Entity = map[string]string{
	// 	"©":     " &copy;",
	// 	"ndash": "\u2013",
	// }
	dec.Strict = false
	err := dec.DecodeElement(&value, &start)
	if err != nil {
		log.Fatal("123")
		return err
	}

	// value = strings.TrimSpace(value)
	*s = String(value)
	return nil
}

type Float7 float64

func (f Float7) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement("HELLL", start)
	// s := fmt.Sprintf("%.7f", f)
	// return e.EncodeElement(s, start)
}
