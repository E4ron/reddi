package handler

import (
	"github.com/gin-gonic/gin"
	"reddi/pkg/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h Handler) GetRouter() *gin.Engine {
	router := gin.Default()

	api := router.Group("/api")
	{
		posts := api.Group("/posts")
		{
			posts.GET("/:item_id")
		}
	}

	return router

}