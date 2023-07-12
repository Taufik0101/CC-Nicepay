package main

import (
	"CC-Nicepay/api"
	"CC-Nicepay/api/middleware"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	router := gin.Default()
	router.Use(middleware.CORSMiddleware())

	api.Routes(router)

	err := router.Run(":5050")
	if err != nil {
		log.Println("Failed To Start System")
	}
	log.Printf("connect for GraphQL playground 5050")
}
