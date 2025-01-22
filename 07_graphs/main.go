package main

import (
	"fmt"
)

// code for analysis
func (p Parser) extractIngredients(doc *goquery.Document) []*models.Ingredient {
	var ingredients []*models.Ingredient
	seen := make(map[string]struct{})

	// Находим все контейнеры с ингредиентами
	containers := doc.Find("div.ingredient-list")
	for i := 0; i < containers.Length(); i++ {
		container := containers.Eq(i)

		// Находим все скрытые элементы внутри контейнера
		hiddenInputs := container.Find("input[type='hidden'][data-declensions]")
		for j := 0; j < hiddenInputs.Length(); j++ {
			input := hiddenInputs.Eq(j)

			// Извлекаем название ингредиента
			ingredientName := input.Parent().Find("span.recipe_ingredient_title").Text()

			if _, ok := seen[ingredientName]; ingredientName != "" && !ok {
				seen[ingredientName] = struct{}{}
				unit := p.extractUnit(input)
				quantity := p.extractQuantity(input)

				// Проверка на валидность данных
				if unit == "" || quantity <= 0 {
					fmt.Printf("Warning: invalid ingredient data for %s\n", ingredientName)
				}

				ingredients = append(ingredients, &models.Ingredient{
					Name:     ingredientName,
					Unit:     unit,
					Quantity: quantity,
				})
			}
		}
	}

	return ingredients
}
