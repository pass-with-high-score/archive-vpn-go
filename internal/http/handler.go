package http

import (
	"bytes"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

type CreateVPNRequest struct {
	Name string `json:"name" binding:"required"`
	DNS  string `json:"dns" binding:"required"`
}

type CreateVPNResponse struct {
	Config string `json:"config"`
}

func CreateVPNConfig(c *gin.Context) {
	var req CreateVPNRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	exePath, err := os.Executable()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	baseDir := filepath.Dir(exePath)

	scriptPath := filepath.Join(baseDir, "wireguard-install.sh")

	cmd := exec.Command(
		"sudo",
		scriptPath,
		"-create",
		"-name", req.Name,
		"-dns", req.DNS,
	)

	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	if err := cmd.Run(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":  err.Error(),
			"stderr": stderr.String(),
		})
		return
	}

	c.JSON(http.StatusOK, CreateVPNResponse{
		Config: stdout.String(),
	})
}
