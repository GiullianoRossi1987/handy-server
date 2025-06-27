package main

import (
	// "pkg"
	// types "types/config"
	"fmt"
	"utils"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	config, _ := utils.GenerateDatabaseConfig()
	fmt.Println(config.Db)
}
