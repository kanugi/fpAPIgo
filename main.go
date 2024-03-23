package main

import (
	"fpGOWebScalable/database"
	"fpGOWebScalable/router"
)

func main() {
	database.StartDB()
	r := router.StartApp()
	r.Run(":8080")
}
