package bsmx_test

import (
	"io/ioutil"
	"log"
	"path/filepath"
	"strings"
	"testing"

	"github.com/go-xmlfmt/xmlfmt"
	"github.com/leonb/go-beersmith/bsmx"
	"github.com/mitchellh/go-homedir"
)

func TestOpenSaveShouldResultInSameFile(t *testing.T) {
	names, err := filepath.Glob("test_assets/*.bsmx")
	if err != nil {
		t.Error(err)
		return
	}

	for _, name := range names {
		// read .bsmlx file
		b, err := ioutil.ReadFile(name)
		if err != nil {
			t.Error(err)
			return
		}

		// convert file contents to string
		xml1 := string(b)
		// trim spaces and newlinews (formatxml doesn't do that)
		xml1 = strings.TrimSpace(xml1)
		// format xml
		xml1 = xmlfmt.FormatXML(xml1, "\t", "  ")

		// read file contents and convert to struct
		// f, err := bsmx.ReadString(xml1)
		f, err := bsmx.ReadBytes(b)
		if err != nil {
			t.Error(err)
			return
		}

		// convert struct back to bytes
		b, err = f.ToXML()
		if err != nil {
			t.Error(err)
			return
		}

		// convert bytes to string
		xml2 := string(b)
		// trim spaces and newlinews (formatxml doesn't do that)
		xml2 = strings.TrimSpace(xml2)
		// format xml
		xml2 = xmlfmt.FormatXML(xml2, "\t", "  ")
		xml2 = strings.ReplaceAll(xml2, "&#xA;", "\n")
		xml2 = strings.ReplaceAll(xml2, "–", "&ndash;")
		xml2 = strings.ReplaceAll(xml2, "“", "&ldquo;")
		xml2 = strings.ReplaceAll(xml2, "”", "&rdquo;")
		xml2 = strings.ReplaceAll(xml2, "&#34;", "&quot;")
		xml2 = strings.ReplaceAll(xml2, "©", "&copy;")
		xml2 = strings.ReplaceAll(xml2, "®", "&reg;")
		xml2 = strings.ReplaceAll(xml2, "™", "&trade;")
		// ®™
		// &reg;&trade;
		// xml2 = html.UnescapeString(xml2)

		// compare original contents with unmarshalled + marshalled contents
		if strings.Compare(xml1, xml2) != 0 {
			err = ioutil.WriteFile("xml1.xml", []byte(xml1), 0600)
			if err != nil {
				t.Error(err)
				return
			}

			err = ioutil.WriteFile("xml2.xml", []byte(xml2), 0600)
			if err != nil {
				t.Error(err)
				return
			}

			log.Fatalf("%s doesnt's marshal back to the same as the original", name)
			// t.Errorf("%s doesnt's marshal back to the same as the original", name)
		}
	}
}

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

	// file, err := bsmx.Open(home + "/.beersmith3/Yeast.bsmx")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// file, err := bsmx.Open(home + "/.beersmith3/Style.bsmx")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// file, err := bsmx.Open(home + "/.beersmith3/Mash.bsmx")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	file, err := bsmx.Open(home + "/.beersmith3/Water.bsmx")
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

	for _, s := range file.Styles {
		log.Println(s.Name)
	}

	for _, s := range file.Styles {
		log.Println(s.Name)
	}

	for _, m := range file.MashProfiles {
		log.Println(m.Name)
	}

	for _, w := range file.WaterProfiles {
		log.Println(w.Name)
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
}
