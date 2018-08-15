package srv

import (
	"github.com/go-lego/engine/bind"
	eproto "github.com/go-lego/engine/proto"
	micro "github.com/micro/go-micro"
)

// Initializer of srv
type Initializer interface {
	InitServices()
}

// Init srv
func Init(s micro.Service, z Initializer) error {
	c := NewContainer()
	z.InitServices()

	eproto.RegisterEventServiceHandler(s.Server(), c)

	bind.Report(s.Server().Options().Name, GetFocusedEvents())
	return nil
}
