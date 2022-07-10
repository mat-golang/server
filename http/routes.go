package http

func (s Service) routes() {
	s.m.Get("/", s.handleIndex())
}
