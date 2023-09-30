// Handler ENDPOINT API managing a user
package handlers

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

// Handler struct holds required services for handler to function

type Handler struct{}

type Config struct {
	Routes *gin.Engine
}

// NewHandler initializes the handler with required injected services along with http routes
// Does not return as it deals directly with a reference to the gin Engine

func NewHandler(config *Config) {

	handler := &Handler{}

	// Create a group, or base url for all routes
	apiRoutes := config.Routes.Group(os.Getenv("ACCOUNT_API_URL"))
	{
		apiRoutes.GET("/me", handler.Me)
		apiRoutes.POST("/signup", handler.Signup) // inscription
		apiRoutes.POST("/signin", handler.Signin) // login
		apiRoutes.POST("/signout", handler.Signout)
		apiRoutes.POST("/tokens", handler.Tokens)
		apiRoutes.POST("/image", handler.Image)
		apiRoutes.DELETE("/image", handler.DeleteImage)
		apiRoutes.PUT("/details", handler.Details)
	}

}

// Me handler calls services for getting
// a user's details
func (handler *Handler) Me(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"hello": "it's me",
	})
}

// Signup handler
func (handler *Handler) Signup(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"hello": "it's signup",
	})
}

// Signin handler
func (handler *Handler) Signin(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"hello": "it's signin",
	})
}

// Signout handler
func (handler *Handler) Signout(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"hello": "it's signout",
	})
}

// Tokens handler
func (handler *Handler) Tokens(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"hello": "it's tokens",
	})
}

// Image handler
func (handler *Handler) Image(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"hello": "it's image",
	})
}

// DeleteImage handler
func (handler *Handler) DeleteImage(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"hello": "it's deleteImage",
	})
}

// Details handler
func (handler *Handler) Details(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"hello": "it's details",
	})
}
