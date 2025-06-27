package main

import (
	"config"
	"pkg"
)

func main() {
	db := config.GetConfigByEnv()
	println(db.Host)
	connection, err := pkg.GeneratePool(db)
	if err != nil {
		panic(err)
	}
	println("apparently connecting")
	defer connection.Close()
}
