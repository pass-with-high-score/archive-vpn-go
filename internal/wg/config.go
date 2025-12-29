package wg

import "fmt"

type ClientConfig struct {
	PrivateKey string
	ClientIP   string

	ServerPublicKey string
	ServerEndpoint  string
	ServerPort      int

	DNS []string
}

func BuildClientConfig(c ClientConfig) string {
	dns := ""
	for i, d := range c.DNS {
		if i > 0 {
			dns += ", "
		}
		dns += d
	}

	return fmt.Sprintf(`[Interface]
PrivateKey = %s
Address = %s
DNS = %s

[Peer]
PublicKey = %s
Endpoint = %s:%d
AllowedIPs = 0.0.0.0/0, ::/0
PersistentKeepalive = 25
`,
		c.PrivateKey,
		c.ClientIP,
		dns,
		c.ServerPublicKey,
		c.ServerEndpoint,
		c.ServerPort,
	)
}
