package main

import (
	"github.com/marcoshuck/todo/internal/conf"
	"github.com/marcoshuck/todo/internal/server"
	"log"
)

func main() {
	cfg, err := conf.ReadServerConfig()
	if err != nil {
		log.Fatalln("Failed to read configuration:", err)
	}
	app, err := server.Setup(cfg)
	if err != nil {
		log.Fatalln("Failed to initialize application:", err)
	}
	if err := server.Run(app); err != nil {
		log.Fatalln("Application exited abruptly:", err)
	}
	log.Println("Closing application...")
}
