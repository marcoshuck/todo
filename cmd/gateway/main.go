package main

import (
	"context"
	"github.com/marcoshuck/todo/pkg/conf"
	"github.com/marcoshuck/todo/pkg/gateway"
	"log"
)

func main() {
	ctx := context.Background()

	cfg, err := conf.ReadGatewayConfig()
	if err != nil {
		log.Fatalln("Failed to read client config:", err)
	}

	gw, err := gateway.Setup(ctx, cfg)
	if err != nil {
		log.Fatalln("Failed to initialize gateway:", err)
	}
	if err := gateway.Run(gw); err != nil {
		log.Fatalln("Failed to run gateway:", err)
	}
}
