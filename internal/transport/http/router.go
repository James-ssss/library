package http

import (
	"crud/internal/core/interface/service"
	"crud/internal/transport/handler"
	"crud/internal/transport/middleware"
	"github.com/gin-gonic/gin"
)

func InitRoutes(service service.AuthService, postService service.PostService, bookService service.BookService) *gin.Engine {
	router := gin.New()

	router.POST("/register", handler.RegisterUser(service))
	router.POST("/login", handler.RegisterUser(service))

	api := router.Group("/api", middleware.AuthMiddleware)
	{
		api.POST("/post", handler.CreatePost(postService))
		api.GET("/post/:id", handler.GetPost(postService))

		api.GET("/books", handler.GetBooks(bookService))
		api.GET("/books/:id", handler.GetBook(bookService))
		api.POST("/books", handler.CreateBook(bookService))
		api.PUT("/books/:id", handler.UpdateBook(bookService))
		api.DELETE("/books/:id", handler.DeleteBook(bookService))

		api.GET("/books/title", handler.GetBooksByCondition(bookService, 1))
		api.GET("/books/genre", handler.GetBooksByCondition(bookService, 2))
		api.GET("/books/author", handler.GetBooksByCondition(bookService, 3))
	}
	return router
}
