package main

import (
	"config"
	"context"
	"pkg"
)

func main() {
	db := config.GetConfigByEnv()
	println(db.Host)
	connection := pkg.GenerateConnection(db)
	println("apparently connecting")
	defer connection.Close(context.Background())
}
