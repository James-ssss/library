package http

import (
	"crud/internal/core/interface/service"
	"crud/internal/transport/handler"
	"crud/internal/transport/middleware"
	"github.com/gin-gonic/gin"
)

func InitRoutes(service service.AuthService, bookService service.BookService, favouriteService service.FavouritesService) *gin.Engine {
	router := gin.New()

	router.POST("/register", handler.RegisterUser(service))
	router.POST("/login", handler.RegisterUser(service))

	api := router.Group("/api", middleware.AuthMiddleware)
	{
		api.GET("/books/title", handler.GetBooksByCondition(bookService, 1))
		api.GET("/books/genre", handler.GetBooksByCondition(bookService, 2))
		api.GET("/books/author", handler.GetBooksByCondition(bookService, 3))

		api.GET("/favourites", handler.GetFavourites(favouriteService))
		api.POST("/favourites/:id", handler.CreateFavourite(favouriteService))
		api.DELETE("/favourites/:id", handler.DeleteFavourite(favouriteService))
	}

	adminApi := router.Group("/admin", middleware.AuthMiddleware, middleware.RoleMiddleware)
	{
		adminApi.GET("/books", handler.GetBooks(bookService))
		adminApi.GET("/books/:id", handler.GetBook(bookService))
		adminApi.POST("/books", handler.CreateBook(bookService))
		adminApi.PUT("/books/:id", handler.UpdateBook(bookService))
		adminApi.DELETE("/books/:id", handler.DeleteBook(bookService))
	}
	return router
}
