package main

import "github.com/voltgizerz/rest-api-go-firestore/config"

func main() {
	config.LoadENV()

	// initialize database
	_ = config.InitDB()
}
