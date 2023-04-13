package main

import (
	"github.com/johnsoncwb/mongo_gin/internal/initializer"
	"github.com/johnsoncwb/mongo_gin/internal/server"
	"log"
)

func init() {
	err := initializer.LoadEnvFile()
	if err != nil {
		log.Println("Error to load Env files:", err)
	}
}

func main() {
	server.Init().Start()
}
