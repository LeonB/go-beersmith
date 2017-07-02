package bsmx

type Tables []Table

func (t Tables) GetRecipes() Recipes {
	recipes := Recipes{}
	for _, subt := range t {
		recipes = append(recipes, subt.GetRecipes()...)
	}
	return recipes
}
