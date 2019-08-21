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
	err := e.EncodeElement(string(s), start)
	if err != nil {
		log.Println("2222")
		return err
	}
	return nil
}

func (s *String) UnmarshalXML(dec *xml.Decoder, start xml.StartElement) error {
	var value string
	dec.Entity = xml.HTMLEntity
	err := dec.DecodeElement(&value, &start)
	if err != nil {
		log.Println("3333")
		return err
	}

	value = strings.TrimSpace(value)
	*s = String(value)
	return nil
}
