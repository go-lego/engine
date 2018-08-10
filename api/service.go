package api

import restful "github.com/emicklei/go-restful"

// Route config element
type Route struct {
	Method  string
	Path    string
	Handler restful.RouteFunction
}

// Service API service
type Service interface {
	Name() string
	Routes() []*Route
}
