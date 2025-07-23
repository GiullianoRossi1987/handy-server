package main

import (
	"fmt"
	"pkg"
	routes "routes"
	usr "routes/users"
	types "types/config"

	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	config := types.PsConfig{}
	config.FromEnv()
	fmt.Println(config.Db)
	pool, err := pkg.GeneratePool(config)
	if err != nil {
		panic(err)
	}
	router := gin.Default()
	routes.SetRouter(router)
	router.POST("/test", routes.TestResponse)

	usr.RouteUsers(router, pool)
	router.Run("localhost:8080")
}
