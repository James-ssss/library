package dbModel

type Book struct {
	Title       string `db:"title"`
	Description string `db:"description"`
	ImageURL    string `db:"image"`
	Genre       string `db:"genre"`
	Year        int    `db:"year"`
	Pages       int    `db:"pages"`
	Author      string `db:"author"`
}
