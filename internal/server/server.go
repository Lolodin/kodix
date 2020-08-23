package serverapi

import "net/http"

type Server struct {
	http.Server
}

func NewServer(Port string)  Server {
	s:= Server{}
	s.Addr = Port
	return s
}

