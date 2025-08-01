package main

import (
	"fmt"
	"pkg"
	routes "routes"
	operations "routes/operations"
	reports "routes/reports"
	satellites "routes/satellites"
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
	if err := router.SetTrustedProxies([]string{"127.0.0.1"}); err != nil {
		panic(err)
	}
	routes.SetRouter(router)
	router.POST("/test", routes.TestResponse)

	usr.RouteUsers(router, pool)
	usr.RouteCustomers(router, pool)
	usr.RouteWorkers(router, pool)

	satellites.RouteAddresses(router, pool)
	satellites.RouteEmails(router, pool)
	satellites.RoutePhones(router, pool)

	reports.RouteReports(router, pool)

	operations.RouteOrders(router, pool)
	operations.RoutePS(router, pool)
	router.Run("localhost:8080")
}
