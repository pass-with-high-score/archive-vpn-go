package main

import (
	"log"

	httpHandler "archive-vpn/internal/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.POST("/vpn/config", httpHandler.CreateVPNConfig)

	addr := ":8089"
	log.Println("VPN backend listening on", addr)
	if err := r.Run(addr); err != nil {
		log.Fatal(err)
	}
}
