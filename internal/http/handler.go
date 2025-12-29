package http

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"archive-vpn/internal/wg"
)

func CreateVPNConfig(c *gin.Context) {
	priv, _, err := wg.GenerateKeyPair()
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	clientIP := "10.10.0.2/32"

	cfg := wg.BuildClientConfig(wg.ClientConfig{
		PrivateKey:      priv,
		ClientIP:        clientIP,
		ServerPublicKey: "SERVER_PUBLIC_KEY_BASE64",
		ServerEndpoint:  "vpn.example.com",
		ServerPort:      51820,
		DNS:             []string{"1.1.1.1", "8.8.8.8"},
	})

	c.Data(http.StatusOK, "text/plain; charset=utf-8", []byte(cfg))
}
