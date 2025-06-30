package main

import (
	// "pkg"
	// types "types/config"
	"fmt"
	"handlers"
	"utils"

	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	config, _ := utils.GenerateDatabaseConfig()
	router := gin.Default()
	router.GET("/health", handlers.GetHealth)
	router.Run("localhost:8080")
	fmt.Println(config.Db)
}
