package main

import "config"

func main() {
	db := config.GetConfigByEnv()
	println(db.Host)
}
