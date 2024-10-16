package config

func (s *Session) SetRoute(route, subRoute string) {
	s.LastRoute = route
	if len(subRoute) > 0 {
		s.LastSub[route] = subRoute
	}
	_ = s.Save()
}
