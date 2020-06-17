package main

import (
	"github.com/pianrspn/go-rest/database"
	"github.com/pianrspn/go-rest/routes"
)

func main() {
	database.Init()
	e := routes.Init()
	e.Logger.Fatal(e.Start(":1323"))
}
