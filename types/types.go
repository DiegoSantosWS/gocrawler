package types

type Data struct {
	Link        string `json:"link" bson:"link"`
	Image       string `json:"image" bson:"image"`
	Title       string `json:"title" bson:"title"`
	Date        string `json:"date" bson:"date"`
	Description string `json:"description" bson:"description"`
}
