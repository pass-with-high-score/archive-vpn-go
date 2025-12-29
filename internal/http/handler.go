package http

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"archive-vpn/internal/config"
	"archive-vpn/internal/wg"
)

type Handler struct {
	Server config.ServerConfig
	IPAM   *wg.IPAM
}

func (h *Handler) CreateVPNConfig(c *gin.Context) {
	priv, pub, err := wg.GenerateKeyPair()
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	clientIP := h.IPAM.Allocate()

	err = wg.AddPeer(
		h.Server.InterfaceName,
		pub,
		clientIP,
	)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	cfg := wg.BuildClientConfig(wg.ClientConfig{
		PrivateKey:      priv,
		ClientIP:        clientIP,
		ServerPublicKey: h.Server.ServerPublicKey,
		ServerEndpoint:  h.Server.ServerEndpoint,
		ServerPort:      h.Server.ServerPort,
		DNS:             []string{"1.1.1.1"},
	})

	c.Data(http.StatusOK, "text/plain; charset=utf-8", []byte(cfg))
}
