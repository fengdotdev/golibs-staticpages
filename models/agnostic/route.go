package agnostic

type Route struct {
	route string
}

func NewRoute(route string) *Route {

	//somechecks here

	//todo
	validRoute := route

	return &Route{
		route: validRoute,
	}
}

func (r *Route) GetRoute() string {
	return r.route
}
