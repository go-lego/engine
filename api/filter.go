package api

import (
	restful "github.com/emicklei/go-restful"
	eerr "github.com/go-lego/engine/error"
	"github.com/go-lego/engine/log"
)

// FilterExecutor filter handler function
type FilterExecutor func(req *restful.Request) *eerr.Error

// Filter API request filter
type Filter interface {
	// Execute do filtering
	Execute(req *restful.Request) *eerr.Error
}

// internalFilter prepare the environment
func internalFilter(req *restful.Request) *eerr.Error {
	log.Debug("Internal filter is executing ...")
	return nil
}

// filterFunctionWrapper wrapper of restful.FilterFunction
func filterFunctionWrapper(fh FilterExecutor) restful.FilterFunction {
	return func(req *restful.Request, rsp *restful.Response, chain *restful.FilterChain) {
		if err := fh(req); err != nil {
			_error(rsp, err)
			return
		}
		chain.ProcessFilter(req, rsp)
	}
}

// AddFilterExecutor add filter executor directly
func AddFilterExecutor(fe FilterExecutor) {
	restful.Filter(filterFunctionWrapper(fe))
}

// AddFilter add filter
func AddFilter(f Filter) {
	AddFilterExecutor(f.Execute)
}
