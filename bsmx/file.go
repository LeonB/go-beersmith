package bsmx

import (
	"bytes"
	"encoding/xml"
	"io"
	"os"
	"strings"
)

type File struct {
	XMLName xml.Name

	MOD      string  `xml:"_MOD_,omitempty"`
	Name     string  `xml:"Name,omitempty"`
	Type     Type    `xml:"Type,omitempty"`
	Dirty    IntBool `xml:"Dirty"`
	Owndata  IntBool `xml:"Owndata"`
	TID      int     `xml:"TID,omitempty"`
	Size     int     `xml:"Size,omitempty"`
	XName    string  `xml:"_XName,omitempty"`
	Allocinc int     `xml:"Allocinc,omitempty"`
	// Data       Data      `xml:"Data,omitempty"`
	Archives             Archives             `xml:"Data>Archive"`
	CarbonationProfiles  CarbonationProfiles  `xml:"Data>Carbonation"`
	EquipmentProfiles    EquipmentProfiles    `xml:"Data>Equipment"`
	FermentationProfiles FermentationProfiles `xml:"Data>Age"`
	Grains               Grains               `xml:"Data>Grain"`
	Hops                 Hops                 `xml:"Data>Hops"`
	MashProfiles         MashProfiles         `xml:"Data>Mash"`
	Miscs                Miscs                `xml:"Data>Misc"`
	Recipes              Recipes              `xml:"Data>Recipe"`
	Styles               Styles               `xml:"Data>Style"`
	Tables               Tables               `xml:"Data>Table"`
	WaterProfiles        WaterProfiles        `xml:"Data>Water"`
	Yeasts               Yeasts               `xml:"Data>Yeast"`
	TExpanded            IntBool              `xml:"_TExpanded"`
	TExtra               IntBool              `xml:"TExtra"`
}

func (f *File) UnmarshalXML(dec *xml.Decoder, start xml.StartElement) error {
	type Alias File
	var alias Alias

	// dec.Entity = xml.HTMLEntity
	err := dec.DecodeElement(&alias, &start)
	if err == nil {
		alias.Name = strings.TrimSpace(alias.Name)
		alias.MOD = strings.TrimSpace(alias.MOD)
		alias.XName = strings.TrimSpace(alias.XName)
		*f = File(alias)
	}
	return err
}

func (f *File) ToXML() ([]byte, error) {
	return xml.Marshal(f)
}

func (f *File) AllRecipes() Recipes {
	recipes := Recipes{}
	for _, recipe := range f.Recipes {
		recipes = append(recipes, recipe)
	}

	recipes = append(recipes, f.Tables.GetRecipes()...)
	return recipes
}

func Open(name string) (*File, error) {
	f, err := os.Open(name)
	if err != nil {
		return nil, err
	}

	return Read(f)
}

func ReadBytes(b []byte) (*File, error) {
	r := bytes.NewReader(b)
	return Read(r)
}

func ReadString(s string) (*File, error) {
	file := File{}
	err := xml.Unmarshal([]byte(s), &file)
	return &file, err
}

func Read(r io.Reader) (*File, error) {
	dec := xml.NewDecoder(r)
	dec.Entity = xml.HTMLEntity
	file := File{}
	err := dec.Decode(&file)
	return &file, err
}
