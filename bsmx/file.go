package bsmx

import (
	"encoding/xml"
	"os"

	"github.com/leonb/go-beersmith/units"
)

type File struct {
	XMLName xml.Name

	_MOD_    units.Date `xml:"_MOD_,omitempty"`
	Name     string     `xml:"Name,omitempty"`
	Type     Type       `xml:"Type,omitempty"`
	Dirty    units.Bool `xml:"Dirty,omitempty"`
	Owndata  units.Bool `xml:"Owndata,omitempty"`
	TID      int        `xml:"TID,omitempty"`
	Size     int        `xml:"Size,omitempty"`
	_XName   *_XName    `xml:"_XName,omitempty"`
	Allocinc int        `xml:"Allocinc,omitempty"`
	// Data       Data      `xml:"Data,omitempty"`
	CarbonationProfiles  CarbonationProfiles  `xml:"Data>Carbonation"`
	FermentationProfiles FermentationProfiles `xml:"Data>Age"`
	Grains               Grains               `xml:"Data>Grain"`
	Hops                 Hops                 `xml:"Data>Hop"`
	MashProfiles         MashProfiles         `xml:"Data>Mash"`
	Recipes              Recipes              `xml:"Data>Recipe"`
	Styles               Styles               `xml:"Data>Style"`
	Tables               Tables               `xml:"Data>Table"`
	WaterProfiles        WaterProfiles        `xml:"Data>Water"`
	Yeasts               Yeasts               `xml:"Data>Yeast"`
	_TExpanded           units.Bool           `xml:"_TExpanded,omitempty"`
}

func (f *File) UnmarshalXML(dec *xml.Decoder, start xml.StartElement) error {
	type Alias File
	var alias Alias

	dec.Entity = xml.HTMLEntity
	err := dec.DecodeElement(&alias, &start)
	if err == nil {
		*f = File(alias)
	}
	return err
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

	dec := xml.NewDecoder(f)
	// dec.Entity = xml.HTMLEntity
	file := File{}
	err = dec.Decode(&file)
	return &file, err
}
