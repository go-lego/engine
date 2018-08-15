package api

import (
	restful "github.com/emicklei/go-restful"
	"github.com/go-lego/engine"
	eerr "github.com/go-lego/engine/error"
	"github.com/go-lego/engine/log"
	"github.com/mohae/deepcopy"
)

// RouteHandler route handler function
type RouteHandler func(ctx *engine.Context, input interface{}) (interface{}, *eerr.Error)

// routeFunctionWrapper wrapper of restful.RouteFunction.
func routeFunctionWrapper(rh RouteHandler, input interface{}) restful.RouteFunction {
	return func(req *restful.Request, rsp *restful.Response) {
		ng := engine.NewEngine(NewDispatcher())
		ctx := ng.NewContext()
		var cpy interface{}
		if input != nil {
			cpy = deepcopy.Copy(input)
			// log.Debug("Copy input:%s", cpy)
			if err := _requestInput(req, cpy); err != nil {
				_error(rsp, eerr.New(101, "Bad input:"+err.Error()))
				return
			}
		}
		ret, err := rh(ctx, cpy)
		if err != nil {
			_error(rsp, err)
			return
		}
		_success(rsp, ret)
	}
}

// Route config element
type Route struct {
	Method  string
	Path    string
	Input   interface{}
	Handler RouteHandler
}

// Service API service
type Service interface {
	Name() string
	Routes() []*Route
}

// AddService add service to API
func AddService(s Service) {
	path := basePath + "/" + s.Name()
	ws := newWebService(path)
	for _, r := range s.Routes() {
		log.Info("Add route %s %s%s", r.Method, path, r.Path)
		switch r.Method {
		case MethodGet:
			ws.Route(ws.GET(r.Path).To(routeFunctionWrapper(r.Handler, r.Input)))
		case MethodPost:
			ws.Route(ws.POST(r.Path).To(routeFunctionWrapper(r.Handler, r.Input)))
		case MethodPut:
			ws.Route(ws.PUT(r.Path).To(routeFunctionWrapper(r.Handler, r.Input)))
		case MethodPatch:
			ws.Route(ws.PATCH(r.Path).To(routeFunctionWrapper(r.Handler, r.Input)))
		case MethodDelete:
			ws.Route(ws.DELETE(r.Path).To(routeFunctionWrapper(r.Handler, r.Input)))
		case MethodHead:
			ws.Route(ws.HEAD(r.Path).To(routeFunctionWrapper(r.Handler, r.Input)))
			// case MethodOptions:
			// 	ws.Route(ws.OPTIONS(r.Path).To(r.Handler))
		}
	}
	restful.Add(ws)
}
