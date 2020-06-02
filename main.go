package main

import (
	"adslot-admin/config"
	"adslot-admin/router"
	"log"
)

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)

	log.Println("init gin server")

	router := router.InitRouter()

	router.Run(":" + config.Config.GetString("port"))
}
