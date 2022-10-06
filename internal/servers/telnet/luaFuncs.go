package telnet

import (
	"github.com/Shopify/go-lua"
	"github.com/gookit/config/v2"
	"github.com/matjam/gobbs/internal/servers/server"
	"github.com/rs/zerolog/log"
)

func (tc TelnetConnection) registerLuaFunctions() {
	tc.l.Register("do_more", tc.doMore())
	tc.l.Register("get_version", tc.getVersion())
	tc.l.Register("get_config", tc.getConfig())
	tc.l.Register("print", tc.print())
	tc.l.Register("log_info", tc.logInfo())
}

func (tc TelnetConnection) logInfo() lua.Function {
	return func(state *lua.State) int {
		str, ok := state.ToString(1)
		if !ok {
			log.Error().Msgf("error calling log_info(msg), incorrect number of arguments.")
			return 0
		}
		log.Info().Msgf("[%v] lua: %v", tc.conn.RemoteAddr().String(), str)
		return 0 // number of results
	}
}

func (tc TelnetConnection) getVersion() lua.Function {
	return func(state *lua.State) int {
		state.PushString(server.Version)
		return 1 // number of results
	}
}

func (tc TelnetConnection) getConfig() lua.Function {
	return func(state *lua.State) int {
		if state.Top() > 2 || state.Top() == 0 {
			log.Error().Msgf("error calling get_config(key, default), need 1 or 2 arguments.")
		}

		key, ok := state.ToString(1)
		if !ok {
			log.Error().Msgf("error calling get_config(key, default), incorrect argument 1.")
			return 0
		}
		def, _ := state.ToString(2)

		state.PushString(config.String(key, def))
		return 1 // number of results
	}
}
