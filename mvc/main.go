package main

import (
	"mvc/app"
	"mvc/db"
)
//GO RUN MAIN.GO
func main() {
	db.StartDbEngine()
	app.StartRoute()
}
