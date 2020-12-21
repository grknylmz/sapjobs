package main

import (
	db "sapjobs.com/v1/conn"
	routes "sapjobs.com/v1/routes"
)

func main() {
	db.Init()
	routes.StartGin()
}
