package models

type Recipe struct {
	ID          string        `bson:"id"`
	IssueID     string        `bson:"issue_id"`
	URL         string        `bson:"url"`
	Title       string        `bson:"title"`
	Ingredients []*Ingredient `bson:"ingredients"`
	Steps       []string      `bson:"steps"`
	ImageURL    string        `bson:"image_url,omitempty"`
}

type Ingredient struct {
	Name     string  `bson:"name"`
	Unit     string  `bson:"unit"`
	Quantity float64 `bson:"quantity"`
}
