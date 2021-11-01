package config

import (
	"time"

	"github.com/multiformats/go-multiaddr"
	manet "github.com/multiformats/go-multiaddr/net"
)

const (
	defaultAdminServerAddr = "/ip4/127.0.0.1/tcp/3102"
	defaultReadTimeout     = Duration(30 * time.Second)
	defaultWriteTimeout    = Duration(30 * time.Second)
)

type AdminServer struct {
	// Admin is the admin API listen address
	ListenMultiaddr string
	ReadTimeout     Duration
	WriteTimeout    Duration
}

func (as *AdminServer) ListenNetAddr() (string, error) {
	maddr, err := multiaddr.NewMultiaddr(as.ListenMultiaddr)
	if err != nil {
		return "", err
	}

	netAddr, err := manet.ToNetAddr(maddr)
	if err != nil {
		return "", err
	}
	return netAddr.String(), nil
}
