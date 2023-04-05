package handler

import (
	"my_api/internal/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	UserService  model.UserService
	TokenService model.TokenService
}

type Config struct {
	R            *gin.Engine
	UserService  model.UserService
	TokenService model.TokenService
}

func NewHandler(c *Config) {
	h := &Handler{
		UserService:  c.UserService,
		TokenService: c.TokenService,
	}

	g := c.R.Group("/api/account")

	g.GET("/me", h.Me)
	g.POST("/singup", h.Singup)
	g.POST("/singin", h.Singin)
	g.POST("/singout", h.Singout)
	g.POST("/tokens", h.Tokens)
	g.POST("/image", h.Image)
	g.DELETE("/image", h.DeleteImage)
	g.PUT("/details", h.Details)

	//g.GET("/", func(c *gin.Context) {
	//	c.JSON(http.StatusOK, gin.H{
	//		"Hello": "people",
	//	})
	//})
} //

//func (h *Handler) Singup(c *gin.Context) {
//	c.JSON(http.StatusOK, gin.H{
//		"hello": "it's singup",
//	})
//}

func (h *Handler) Singin(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"hello": "it's singin",
	})
}

func (h *Handler) Singout(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"hello": "it's singout",
	})
}

func (h *Handler) Tokens(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"hello": "tokens",
	})
}

func (h *Handler) Image(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"hello": "it's image",
	})
}

func (h *Handler) DeleteImage(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"hello": "it's delete image",
	})
}

func (h *Handler) Details(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"hello": "it's details",
	})
}
