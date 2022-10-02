package telnet

import (
	"fmt"
	"net"

	"github.com/Shopify/go-lua"
	"github.com/gookit/config/v2"
	"github.com/matjam/gobbs/internal/servers/server"
	"github.com/pborman/ansi"
	"github.com/rs/zerolog/log"
)

type TelnetConnection struct {
	conn        net.Conn
	user        *string
	accessLevel int
	state       *lua.State
}

type TelnetServer struct {
	connections map[string]TelnetConnection
}

func NewTelnetServer() server.Server {
	return TelnetServer{
		connections: make(map[string]TelnetConnection),
	}
}

func (t TelnetServer) Start() {
	telnetServerPort := config.Int("telnet.port", 2125)
	telnetServerAddress := config.String("telnet.address", "127.0.0.1")

	log.Info().Msgf("telnet service starting on %v:%v ...", telnetServerAddress, telnetServerPort)
	ln, err := net.Listen("tcp", fmt.Sprintf("%v:%v", telnetServerAddress, telnetServerPort))
	if err != nil {
		log.Fatal().Msgf("couldn't listen to telnet port %v", err.Error())
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatal().Msgf("couldn't accept connection from %v: %v", conn.RemoteAddr(), err.Error())
		}
		go t.handleTelnetConnection(conn)
	}
}

func (t TelnetServer) handleTelnetConnection(conn net.Conn) {
	user := ""
	tc := TelnetConnection{
		conn:        conn,
		user:        &user,
		accessLevel: -1,
	}

	// add the connection to the map of current connections. We use the remote host:port
	// as the key because that's all we got right now.
	t.connections[tc.conn.RemoteAddr().String()] = tc
	log.Info().Msgf("[%v] accepted connection via telnet", tc.conn.RemoteAddr())

	// General operation:
	//
	// * Call the LUA function handler registered for the entrypoint to the BBS
	// * The LUA function will have access to functions that provide all the
	//   necessary features for things like access list of connected users, accessing
	//   the database, etc. That function will call a function that will wait for input
	//   and then it will handle the input all in LUA. It will
	// * When the LUA function ends, that means the connection can be closed.

	tc.state = lua.NewState()
	lua.OpenLibraries(tc.state)

	tc.state.Register("send", luaSend)

	if err := lua.DoFile(tc.state, "lua/connect.lua"); err != nil {
		panic(err)
	}

	c := ansi.NewWriter(conn)
	c.Colorize()
	c.ForceReset()
	c.Blue().Bold().WriteString(config.String("telnet.banner", "GOBBS"))
	conn.Close()
}

func luaSend(state *lua.State) int {
	return 0
}
