package debugserver

type Option func(*BebugServer)

func WithAddress(address string) Option {
	return func(server *BebugServer) {
		server.server.Addr = address
	}
}
