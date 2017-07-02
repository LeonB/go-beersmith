package bsmx

import "github.com/leonb/go-beersmith/units"

type Table struct {
	_MOD_    units.Date `xml:"_MOD_"`
	Name     string     `xml:"Name"`
	Type     Type       `xml:"Type"`
	Dirty    units.Bool `xml:"Dirty"`
	Owndata  units.Bool `xml:"Owndata"`
	TID      int        `xml:"TID"`
	Size     int        `xml:"Size"`
	_XName   *_XName    `xml:"_XName"`
	Allocinc int        `xml:"Allocinc"`
	// Data       *Data      `xml:"Data"`
	Recipes    Recipes    `xml:"Data>Recipe"`
	Tables     Tables     `xml:"Data>Table"`
	_TExpanded units.Bool `xml:"_TExpanded"`
}

func (t Table) GetRecipes() Recipes {
	recipes := Recipes{}
	for _, recipe := range t.Recipes {
		recipes = append(recipes, recipe)
	}

	recipes = append(recipes, t.Tables.GetRecipes()...)
	return recipes
}
