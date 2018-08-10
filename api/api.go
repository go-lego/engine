package api

import (
	"errors"
	"regexp"
	"strings"

	"github.com/emicklei/go-restful"
	eerr "github.com/go-lego/engine/error"
	"github.com/go-lego/engine/log"
	"github.com/go-playground/form"
	web "github.com/micro/go-web"
	validator "gopkg.in/go-playground/validator.v9"
)

// API is based on go-restful

var (
	basePath = "/" // service base path

	// ErrBadName bad API service name
	ErrBadName = errors.New("Bad API service name")
)

func convertNameToPath(name string) (string, error) {
	re := regexp.MustCompile("^v[0-9]+$")
	arr := strings.Split(name, ".")
	index := 0
	for i, a := range arr {
		if a == "api" {
			index = i
			break
		}
	}
	if index == 0 {
		return "", ErrBadName
	}
	parts := []string{""}
	parts = append(parts, arr[index+1:]...)
	l := len(parts)
	v := re.MatchString(parts[1])
	if (v && l == 3) || (!v && l == 2) {
		return strings.Join(parts, "/"), nil
	}
	return "", ErrBadName
}

// Initializer of API server
type Initializer interface {
	// InitFilters set go-restful filters
	InitFilters()

	// InitServices set services
	InitServices()
}

// Init API server
func Init(service web.Service, z Initializer) error {
	log.Info("API service is initializing ...")
	var err error
	if basePath, err = convertNameToPath(service.Options().Name); err != nil {
		return err
	}
	service.Handle("/", restful.DefaultContainer)

	log.Info("Enable content encoding")
	restful.DefaultContainer.EnableContentEncoding(true)

	log.Info("Set internal filter")
	restful.Filter(internalFilter)

	log.Info("Initialize customer filters")
	z.InitFilters()

	log.Info("Initialize customer services")
	z.InitServices()
	return nil
}

var (
	// MethodGet GET
	MethodGet = "GET"

	// MethodPost POST
	MethodPost = "POST"

	// MethodPut PUT
	MethodPut = "PUT"

	// MethodPatch PATCH
	MethodPatch = "PATCH"

	// MethodDelete DELETE
	MethodDelete = "DELETE"

	// MethodHead HEAD
	MethodHead = "HEAD"

	// MethodOptions OPTIONS
	MethodOptions = "OPTIONS"
)

// AddService add service to API
func AddService(s Service) {
	path := basePath + "/" + s.Name()
	ws := newWebService(path)
	for _, r := range s.Routes() {
		log.Info("Add route %s %s%s", r.Method, path, r.Path)
		switch r.Method {
		case MethodGet:
			ws.Route(ws.GET(r.Path).To(r.Handler))
		case MethodPost:
			ws.Route(ws.POST(r.Path).To(r.Handler))
		case MethodPut:
			ws.Route(ws.PUT(r.Path).To(r.Handler))
		case MethodPatch:
			ws.Route(ws.PATCH(r.Path).To(r.Handler))
		case MethodDelete:
			ws.Route(ws.DELETE(r.Path).To(r.Handler))
		case MethodHead:
			ws.Route(ws.HEAD(r.Path).To(r.Handler))
			// case MethodOptions:
			// 	ws.Route(ws.OPTIONS(r.Path).To(r.Handler))
		}
	}
	restful.Add(ws)
}

// NewWebService create new go-restful web service
func newWebService(path string) *restful.WebService {
	ws := new(restful.WebService)
	ws.Path(path).
		Consumes(restful.MIME_XML, restful.MIME_JSON, "application/x-www-form-urlencoded", "text/xml").
		Produces(restful.MIME_JSON, restful.MIME_XML)
	return ws
}

// AddFilter add filter to go-restful
func AddFilter(f Filter) {
	restful.Filter(f.Execute)
}

var (
	formDecoder   = form.NewDecoder()
	formValidator = validator.New()
)

// RequestInput get request input data as an entity.
// Make use of form & validator to check input automatically.
func RequestInput(req *restful.Request, e interface{}) error {
	values := req.Request.URL.Query()
	req.Request.ParseForm()
	for k, v := range req.Request.PostForm {
		values[k] = v
	}
	for k, v := range req.PathParameters() {
		values[k] = []string{v}
	}
	if err := formDecoder.Decode(e, values); err != nil {
		return err
	}
	return validator.New().Struct(e)
}

// Error output error response
func Error(rsp *restful.Response, err *eerr.Error) {
	rsp.WriteEntity(err)
}

// Success output success response
func Success(rsp *restful.Response, data interface{}) {
	d := map[string]interface{}{
		"code":    0,
		"message": "Success",
	}
	if data != nil {
		d["data"] = data
	}
	rsp.WriteEntity(d)
}
