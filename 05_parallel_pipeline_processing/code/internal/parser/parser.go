package parser

import (
	"bytes"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"unicode"

	"github.com/PuerkitoBio/goquery"
	"github.com/hahaclassic/algorithm-analysis/05_parallel_pipeline/internal/models"
)

// HTML elements
// url: <meta property="og:url" content="https://www.edimdoma.ru/retsepty/121730-salat-simvol-goda">
//
// name: <h1 class="recipe-header__name" data-promo-recipe="false">Соленые зеленые помидоры</h1>
//
// img: <img alt="Соленые зеленые помидоры" title="Соленые зеленые помидоры"
//     src="https://e1.edimdoma.ru/data/recipes/0012/0003/120003-ed4_wide.jpg?1628775312">
//
// ingredient title: <span class="recipe_ingredient_title">помидоры зеленые</span>
//
// ingredient quantity + unit<input type="hidden" name="753297_amount" id="753297_amount"
//     value="0.375" data-declensions="[&quot;кг&quot;,&quot;кг&quot;,&quot;кг&quot;]">
//
// steps: <div data-module="step_hint" class="moduleinited">Помидоры и петрушку помыть, чеснок почистить.</div>

var (
	ErrNoExist = errors.New("html element does not exist")
)

type Parser struct{}

func (p Parser) ParseHTMLToRecipe(html []byte) (*models.Recipe, error) {
	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(html))
	if err != nil {
		return nil, fmt.Errorf("failed to load HTML: %w", err)
	}

	recipe := &models.Recipe{
		Title:       doc.Find("h1.recipe-header__name").Text(),
		ImageURL:    p.extractImageURL(doc),
		Ingredients: p.extractIngredients(doc),
		URL:         p.extractURL(doc),
	}

	doc.Find("div[data-module='step_hint']").Each(func(i int, s *goquery.Selection) {
		step := removeNonPrintable(strings.TrimSpace(s.Text()))
		if step != "" {
			recipe.Steps = append(recipe.Steps, step)
		}
	})

	if recipe.Title == "" || len(recipe.Ingredients) == 0 || len(recipe.Steps) == 0 {
		return nil, ErrNoExist
	}

	return recipe, nil
}

func (p Parser) extractIngredients(doc *goquery.Document) []*models.Ingredient {
	var ingredients []*models.Ingredient

	seen := make(map[string]struct{})
	doc.Find("input[type='hidden'][data-declensions]").Each(func(i int, s *goquery.Selection) {
		ingredientName := s.Parent().Find("span.recipe_ingredient_title").Text()

		if _, ok := seen[ingredientName]; ingredientName != "" && !ok {
			seen[ingredientName] = struct{}{}

			ingredients = append(ingredients, &models.Ingredient{
				Name:     ingredientName,
				Unit:     p.extractUnit(s),
				Quantity: p.extractQuantity(s),
			})
		}
	})

	return ingredients
}

func (Parser) extractImageURL(doc *goquery.Document) string {
	imgURL, exists := doc.Find("div.thumb-slider__slide img").Attr("src")
	if exists {
		return imgURL
	}
	return ""
}

func (Parser) extractURL(doc *goquery.Document) string {
	url, exists := doc.Find("meta[property='og:url']").Attr("content")
	if exists {
		return url
	}
	return ""
}

func (Parser) extractQuantity(s *goquery.Selection) float64 {
	quantityStr, exists := s.Attr("value")
	var quantity float64
	if exists {
		parsedQuantity, err := strconv.ParseFloat(quantityStr, 64)
		if err == nil {
			quantity = parsedQuantity
		}
	}

	return quantity
}

func (Parser) extractUnit(s *goquery.Selection) string {
	unitRaw, exists := s.Attr("data-declensions")
	var unit string
	if exists {
		unitRaw = strings.Trim(unitRaw, "[]")
		units := strings.Split(unitRaw, ",")
		if len(units) > 0 {
			unit = strings.Trim(units[0], `"'`)
		}
	}

	return unit
}

func removeNonPrintable(s string) string {
	return strings.Map(func(r rune) rune {
		if unicode.IsPrint(r) {
			return r
		}
		return -1 // Удаляем символ
	}, s)
}
