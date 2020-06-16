package main

import "github.com/pianrspn/go-rest/routes"

func main() {
	e := routes.Init()
	e.Logger.Fatal(e.Start(":1323"))
}
