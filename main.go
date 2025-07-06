package main

import (
	// "pkg"
	// types "types/config"
	"fmt"
	"handlers"
	types "types/config"

	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	config := types.PsConfig{}
	config.FromEnv()
	fmt.Println(config.Db)
	router := gin.Default()
	handlers.SetRouter(router)
	router.POST("/test", handlers.TestResponse)
	router.Run("localhost:8080")
}
