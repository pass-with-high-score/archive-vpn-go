package main

import (
	"archive-vpn/internal/config"
	"archive-vpn/internal/wg"
	"log"

	"github.com/gin-gonic/gin"

	httpHandler "archive-vpn/internal/http"
)

func main() {
	serverCfg := config.Load()
	ipam := wg.NewIPAM(2)

	h := &httpHandler.Handler{
		Server: serverCfg,
		IPAM:   ipam,
	}

	r := gin.Default()
	r.POST("/vpn/config", h.CreateVPNConfig)

	addr := ":8080"
	log.Println("VPN backend listening on", addr)

	if err := r.Run(addr); err != nil {
		log.Fatal(err)
	}
}
