package main

import (
	"log"

	"github.com/YnMann/DtMS/services/sso/internal/config"
)

func main() {
	// Create new configuration obj
	if err := config.Init(); err != nil {
		log.Fatalf("%s", err.Error())
	}
}
