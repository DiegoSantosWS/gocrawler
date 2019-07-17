package types

type Data struct {
	ID          int64  `json:"id" db:"pid"`
	Link        string `json:"link" db:"link"`
	Image       string `json:"image" db:"image"`
	Title       string `json:"title" db:"title"`
	Date        string `json:"date" db:"date"`
	Description string `json:"description" db:"description"`
}
