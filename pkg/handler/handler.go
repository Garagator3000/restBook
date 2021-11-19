package handler

import (
	"github.com/Garagator3000/restBook/pkg/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (handler *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	rest := router.Group("/books")
	{
		rest.GET("/book/:id", handler.GetById)
		rest.GET("/all", handler.GetAll)
		rest.POST("/new", handler.CreateNew)
		rest.PUT("/update/:id", handler.UpdateById)
		rest.DELETE("/delete/:id", handler.DeleteById)
	}
	return router
}
