package server

import (
	"go_clean_api/api/infra/http/routes"
	"log"

	"github.com/gin-gonic/gin"
)

type Server struct {
	port   string
	server *gin.Engine
}

func NewServer() Server {
	return Server{
		port:   "5000",
		server: gin.Default(),
	}
}

func (s *Server) Run() {
	router := routes.ConfigRoutes(s.server)
	router.SetTrustedProxies([]string{"127.0.0.1"})

	log.Fatal(router.Run(":" + s.port))
}
