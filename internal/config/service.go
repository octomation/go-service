package config

import (
	"strings"

	"go.octolab.org/toolkit/config"
)

type Service struct {
	Name  string
	Build struct {
		Commit   string
		Date     string
		Version  string
		Features config.Features
	}
	Server Server
}

type Server struct {
	Connect Protocol
	Gateway Protocol
	GRPC    Protocol
	Profile Protocol
	Twirp   Protocol
}

type Protocol struct {
	Address string
	Secured bool
}

func (proto Protocol) BaseURL() string {
	base := proto.Address
	if !strings.Contains(base, "://") {
		scheme := "http://"
		if proto.Secured {
			scheme = "https://"
		}
		base = scheme + base
	}
	return base
}

func (proto Protocol) IsEnabled() bool {
	return proto.Address != ""
}
