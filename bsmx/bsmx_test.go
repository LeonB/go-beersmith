package bsmx_test

import (
	"encoding/xml"
	"log"
	"os"
	"testing"

	"github.com/leonb/go-beersmith/bsmx"
)

func TestRecipe(t *testing.T) {
	f, _ := os.Open("/home/leon/.beersmith2/Recipe.bsmx")

	dec := xml.NewDecoder(f)
	// dec.Entity = xml.HTMLEntity
	recipes := bsmx.RecipesFile{}
	err := dec.Decode(&recipes)
	log.Println(err)

	// recipes.AllRecipes()
	for _, recipe := range recipes.AllRecipes() {
		log.Println(recipe.Name)
	}
}
