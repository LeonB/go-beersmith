package bsmx

import (
	"encoding/xml"

	"github.com/leonb/go-beersmith/units"
)

type RecipesFile struct {
	XMLName xml.Name `xml:"Recipe"`

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
	Tables     Tables     `xml:"Data>Table"`
	Recipes    Recipes    `xml:"Data>Recipe"`
	_TExpanded units.Bool `xml:"_TExpanded,omitempty"`
}

func (r *RecipesFile) UnmarshalXML(dec *xml.Decoder, start xml.StartElement) error {
	type Alias RecipesFile
	var alias Alias

	dec.Entity = xml.HTMLEntity
	err := dec.DecodeElement(&alias, &start)
	if err == nil {
		*r = RecipesFile(alias)
	}
	return err
}

func (r *RecipesFile) AllRecipes() Recipes {
	recipes := Recipes{}
	for _, recipe := range r.Recipes {
		recipes = append(recipes, recipe)
	}

	recipes = append(recipes, r.Tables.GetRecipes()...)
	return recipes
}
