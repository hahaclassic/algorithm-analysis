package main

import (
	"fmt"
)

// code for analysis
func (p Parser) extractIngredients(doc *goquery.Document) []*models.Ingredient {
	var ingredients []*models.Ingredient          // 1
	seen := make(map[string]struct{})             // 2
	containers := doc.Find("div.ingredient-list") // 3

	for i := 0; i < containers.Length(); i++ { // 4
		container := containers.Eq(i)                                            // 5
		hiddenInputs := container.Find("input[type='hidden'][data-declensions]") // 6
		for j := 0; j < hiddenInputs.Length(); j++ {                             // 7
			input := hiddenInputs.Eq(j)                                                  // 8
			ingredientName := input.Parent().Find("span.recipe_ingredient_title").Text() // 9
			if _, ok := seen[ingredientName]; ingredientName != "" && !ok {              // 10
				seen[ingredientName] = struct{}{}    // 11
				unit := p.extractUnit(input)         // 12
				quantity := p.extractQuantity(input) // 13
				if unit == "" || quantity <= 0 {     // 14
					fmt.Printf("Warning: invalid ingredient data for %s\n", ingredientName) // 15
				}

				ingredients = append(ingredients, &models.Ingredient{ // 16
					Name:     ingredientName,
					Unit:     unit,
					Quantity: quantity,
				})
			}
		}
		fmt.Printf("Processed %d hidden inputs in container %d\n", hiddenInputs.Length(), i) // 17
	}

	return ingredients // 18
}
