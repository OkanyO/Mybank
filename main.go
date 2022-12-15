package main

import (
	"nuel/go-bank/api"

	"nuel/go-bank/database"
)

func main() {

	database.InitDatabase()
	api.StartApi()
}
