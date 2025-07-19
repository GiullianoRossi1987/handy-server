package main

import (
	"pkg"
	// types "types/config"
	"fmt"
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
	router.POST("/user/add", usr.AddUserHandler(pool))
	router.GET("/user/get-login/:login", usr.GetUserByLogin(pool))
	router.Run("localhost:8080")
}
