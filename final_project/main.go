package main

import (
	"final_project/config"
	"final_project/router"
)

func main() {
	config.StartDB()
	r := router.StartApp()
	r.Run(":8081")
}
