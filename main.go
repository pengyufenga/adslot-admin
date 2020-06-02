package main

import (
	"adslot-admin/config"
	"adslot-admin/router"
	"log"
)

func main() {
	log.Println("init gin server")

	router := router.InitRouter()

	router.Run(":"+config.Config.GetString("port"))
}
