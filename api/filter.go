package api

import (
	restful "github.com/emicklei/go-restful"
	"github.com/go-lego/engine/log"
)

// Filter API request filter
type Filter interface {
	// Execute do filtering
	Execute(req *restful.Request, rsp *restful.Response, chain *restful.FilterChain)
}

// internalFilter prepare the environment
func internalFilter(req *restful.Request, rsp *restful.Response, chain *restful.FilterChain) {
	log.Debug("Boot filter")
	chain.ProcessFilter(req, rsp)
}
