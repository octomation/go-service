package config

import (
	"strings"

	"go.octolab.org/toolkit/config"
)

var Defaults = Service{
	Name: "service",
	Server: Server{
		Profile: Protocol{Host: "localhost:3360"},
		Twirp:   Protocol{Host: "localhost:8080"},
		Gateway: Protocol{Host: "localhost:8081"},
		Connect: Protocol{Host: "localhost:8890"},
		GRPC:    Protocol{Host: "localhost:8891"},
	},
}

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
	Host   string
	Secure bool
}

func (proto Protocol) BaseURL() string {
	base := proto.Host
	if !strings.Contains(base, "://") {
		scheme := "http://"
		if proto.Secure {
			scheme = "https://"
		}
		base = scheme + base
	}
	return base
}

func (proto Protocol) IsEnabled() bool {
	return proto.Host != ""
}
