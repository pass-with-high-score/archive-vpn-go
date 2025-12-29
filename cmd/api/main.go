package main

import (
	"log"

	"github.com/gin-gonic/gin"

	httpHandler "archive-vpn/internal/http"
)

func main() {
	r := gin.Default()

	r.POST("/vpn/config", httpHandler.CreateVPNConfig)

	addr := ":8080"
	log.Println("VPN backend listening on", addr)

	if err := r.Run(addr); err != nil {
		log.Fatal(err)
	}
}
