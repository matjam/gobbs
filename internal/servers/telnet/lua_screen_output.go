package telnet

import (
	"github.com/Shopify/go-lua"
	"github.com/rs/zerolog/log"
)

func (tc TelnetConnection) doMore() lua.Function {
	return func(state *lua.State) int {
		tc.w.WriteString("More [Y,n]")
		var b []byte
		tc.conn.Read(b)

		return 0 // number of results
	}
}

func (tc TelnetConnection) print() lua.Function {
	return func(state *lua.State) int {
		str, ok := state.ToString(1)
		if !ok {
			log.Error().Msgf("error calling send(msg), incorrect number of arguments.")
			return 0
		}
		tc.w.WriteString(str)
		return 0 // number of results
	}
}
