package server

type Server struct {
	host string
	port string
}

func New(host, port string) *Server {
	return &Server{host, port}
}

func (s Server) Serve()  {
	panic("not implemented")
}