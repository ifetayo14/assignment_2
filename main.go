package main

import (
	"assignment_2/config"
	"assignment_2/route"

	_ "github.com/lib/pq"
)

func main() {
	config.StartDB()

	route.StartRoute().Run(":8080")
}
