package handler

import (
	"fmt"
	"github.com/Nigelmes/L0/internal/repository"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Handler struct {
	repo *repository.Repository
}

func NewHandler(repo *repository.Repository) *Handler {
	return &Handler{repo: repo}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	router.LoadHTMLGlob("../internal/template/index.html")
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
	router.POST("/", h.home)
	return router
}

func (h *Handler) home(c *gin.Context) {
	uuid := c.PostForm("id")
	fmt.Printf("%T   %s", uuid, uuid)
	c.JSON(http.StatusOK, gin.H{"uuid": uuid})
}
