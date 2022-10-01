package servers

import "github.com/matjam/gobbs/internal/servers/telnet"

type Server interface {
	Start()
}

func Start() {
	telnetServer := telnet.TelnetServer
}
