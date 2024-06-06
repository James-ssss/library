package handler

import (
	"crud/internal/core/interface/service"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
	"strconv"
)

type handlerFavourite struct {
	BookID int `json:"book_id"`
	UserID int `json:"user_id"`
}

func CreateFavourite(service service.FavouritesService) gin.HandlerFunc {
	return func(c *gin.Context) {
		bId := c.Param("id")
		bookId, err := strconv.Atoi(bId)
		login := c.GetString("user")

		id, err := service.CreateFavourite(c.Request.Context(), login, bookId)

		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err})
			return
		}

		c.JSON(http.StatusOK, gin.H{"book": id})

	}
}

func GetFavourites(service service.FavouritesService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var books []handlerBooks

		login := c.GetString("user")

		rBooks, ids, err := service.GetFavourites(c.Request.Context(), login)

		if err != nil {
			slog.Error(err.Error())

			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "ошибка получения книги"})

			return

		}

		for index, book := range rBooks {
			var b handlerBooks
			b.Id = ids[index]
			b.Title = book.Title
			b.Description = book.Description
			b.ImageURL = book.ImageURL
			b.Genre = book.Genre
			b.Pages = book.Pages
			b.Year = book.Year
			b.Author = book.Author
			books = append(books, b)
		}

		c.JSON(http.StatusOK, books)

	}
}

func DeleteFavourite(service service.FavouritesService) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		numberId, err := strconv.Atoi(id)

		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest,
				gin.H{"message": "неверно передан id книги"})

			return
		}

		err = service.DeleteFavourite(c.Request.Context(), numberId)

		if err != nil {
			slog.Error(err.Error())

			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "ошибка получения книги"})

			return

		}

		c.JSON(http.StatusOK, gin.H{"book deleted": id})

	}
}
