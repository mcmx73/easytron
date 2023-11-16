package frontrpc

import "github.com/mcmx73/easytron/wallet"

type WithServerOption func(*Server)

func NewServer(options ...WithServerOption) *Server {
	s := &Server{}
	for _, opt := range options {
		opt(s)
	}
	return s
}

type Server struct {
	walletManager *wallet.Manager
}

func (s *Server) Start() (err error) {
	return nil
}
