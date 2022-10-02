package servers

import (
	"sync"

	"github.com/matjam/gobbs/internal/servers/server"
	"github.com/matjam/gobbs/internal/servers/telnet"
)

func Start() {
	var wg sync.WaitGroup

	serverDefinitions := []server.Server{
		telnet.NewTelnetServer(),
	}

	wg.Add(len(serverDefinitions))
	for _, s := range serverDefinitions {
		go func(s server.Server) {
			s.Start()
			wg.Done()
		}(s)
	}
	wg.Wait()
}
