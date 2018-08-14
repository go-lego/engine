package filter

import (
	"errors"
	"strings"

	restful "github.com/emicklei/go-restful"
	"github.com/go-lego/engine"
	eerr "github.com/go-lego/engine/error"
	"github.com/go-lego/engine/log"
)

// AuthCheckFunc function to get account by token
type AuthCheckFunc func(token string) (engine.Account, error)

// AuthExcludeFunc function returns auth check exclude list
type AuthExcludeFunc func() map[string]string

// Auth default auth filter
type Auth struct {
	Check   AuthCheckFunc
	Exclude AuthExcludeFunc
}

// retrieveAccount retrieve account data
func (a *Auth) retrieveAccount(req *restful.Request, must bool) error {
	re := errors.New("Unauthorized")
	token := req.HeaderParameter("Authorization")
	if must && token == "" {
		return re
	}
	if a.Check == nil {
		log.Info("Auth filter's check function is not provided")
		return re
	}
	acc, err := a.Check(token)
	if err != nil {
		if must {
			log.Debug("Failed to get account by token:%s", err)
			return re
		}
		return nil
	}
	req.SetAttribute("Account", acc)
	return nil
}

// Execute check auth
func (a *Auth) Execute(req *restful.Request) *eerr.Error {
	log.Debug("Auth filter is executing ...")
	excludes := a.Exclude()
	uri := strings.Split(req.Request.RequestURI, "?")[0]
	m := req.Request.Method
	ex, ok := excludes[uri]
	if err := a.retrieveAccount(req, !ok || (ex != "*" && ex != m)); err != nil {
		return eerr.New(403, "Forbidden")
	}
	return nil
}
