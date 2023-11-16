package frontrpc

import "github.com/mcmx73/easytron/wallet"

func WithWalletManager(wm *wallet.Manager) WithServerOption {
	return func(c *Server) {
		c.walletManager = wm
	}
}
