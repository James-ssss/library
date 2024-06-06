package dbModel

type Favourites struct {
	BookID int `db:"book_id"`
	UserID int `db:"user_id"`
}
