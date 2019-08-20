package bsmx_test

import (
	"log"
	"os"
	"testing"

	"github.com/leonb/go-beersmith/bsmx"
	"github.com/mitchellh/go-homedir"
)

func TestRecipe(t *testing.T) {
	home, err := homedir.Dir()
	if err != nil {
		log.Fatal(err)
	}

	// file, err := bsmx.Open(home + "/.beersmith3/Carbonation.bsmx")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// file, err := bsmx.Open(home + "/.beersmith3/Age.bsmx")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	file, err := bsmx.Open(home + "/.beersmith3/Yeast.bsmx")
	if err != nil {
		log.Fatal(err)
	}

	// file, err := bsmx.Open(home + "/.beersmith3/Grain.bsmx")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// recipes.AllRecipes()
	// for _, recipe := range recipes.AllRecipes() {
	// 	log.Println(recipe.Name)
	// }

	for _, c := range file.CarbonationProfiles {
		log.Println(c.Name)
		log.Println(c.Type.String())
	}

	for _, f := range file.FermentationProfiles {
		log.Println(f.Name)
		log.Println(f.Type.String())
	}

	for _, g := range file.Grains {
		log.Println(g.Name)
	}

	for _, y := range file.Yeasts {
		log.Println(y.Name)
	}

	for _, t := range file.Tables {
		if t.Name == "Brew Log" {
			for _, recipe := range t.GetRecipes() {
				log.Println("STYLE:")
				log.Println(recipe.Style.Type.String())

				log.Println("GRAINS:")
				for _, g := range recipe.Ingredients.Grains {
					log.Println(g.Name)
				}

				log.Println("HOPS:")
				for _, h := range recipe.Ingredients.Hops {
					log.Println(h.Name)
				}

				log.Println("MISCS:")
				for _, m := range recipe.Ingredients.Miscs {
					log.Println(m.Name)
				}

				log.Println("YEASTS:")
				for _, y := range recipe.Ingredients.Yeasts {
					log.Println(y.Name)
				}
			}
		}
	}

	os.Exit(1)
}
