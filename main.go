package main

import (
	"github.com/ngrnr/task-5-pbi-btpns-nugraharamadan/database"
	"github.com/ngrnr/task-5-pbi-btpns-nugraharamadan/router"
)

func main() {
	database.InitDB()
	database.MigrateDB()
	r := router.RouteInit()
	r.Run(":8080")
}
