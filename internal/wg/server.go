package wg

import (
	"fmt"
	"os/exec"
)

func AddPeer(interfaceName, publicKey, clientIP string) error {
	cmd := exec.Command(
		"wg",
		"set",
		interfaceName,
		"peer",
		publicKey,
		"allowed-ips",
		clientIP,
	)

	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("wg set failed: %v - %s", err, string(output))
	}
	return nil
}
