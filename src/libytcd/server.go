package libytcd

type Server struct {
}

func NewServer(ports []Port) (s *Server) {
	s = new(Server)
	return
}
