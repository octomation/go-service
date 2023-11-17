package config

import "strings"

type Server struct {
	Name  string
	Host  string
	Debug bool
}

func (cnf *Server) BaseURL() string {
	if !strings.HasPrefix(cnf.Host, "https://") &&
		!strings.HasPrefix(cnf.Host, "http://") {
		return "http://" + cnf.Host
	}
	return cnf.Host
}
