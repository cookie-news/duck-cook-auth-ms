package api

import (
	"duck-cook-auth/controller"
	"duck-cook-auth/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Server struct {
	controller controller.Controller
}

func NewServer(controller controller.Controller) *Server {
	return &Server{controller}
}

func (s *Server) Start(addr string) error {

	docs.SwaggerInfo.Title = "Duck Cook Auth"
	docs.SwaggerInfo.Description = "Duck Cook Auth"
	docs.SwaggerInfo.Version = "0.0.1"
	docs.SwaggerInfo.BasePath = "/v1"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	r := gin.Default()

	v1 := r.Group("/v1")
	{
		costumer := v1.Group("/customer")
		{
			costumer.POST("/", s.controller.CreateCustomerHandler)
		}
		auth := v1.Group("/auth")
		{
			auth.POST("/login", s.controller.LoginHandler)
		}
	}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return r.Run(":" + addr)
}
