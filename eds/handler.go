package eds

import (
	"context"
	"reflect"
	"strings"
	"time"

	eproto "github.com/go-lego/engine/eds/proto"
	"github.com/go-lego/engine/log"
)

// Actor interface, used for handling concrete events.
// ID() get the module id, the first field of event ID.
// For example event ID "account.register" indicates that actor ID is "account".
// Actor functions must be started with "On" and hump event ID to be invoked.
// For example "OnRegister" will be invoked by "account.register",
// "OnRegisterSuccess" will be invoked by "account.register.success"
type Actor interface {
	ID() string
}

// ActorContainer contains actors
// It use reflection to dispatch event to concrete action
type ActorContainer struct {
	actors map[string]Actor
}

// AddActor add action to handler
func (hi *ActorContainer) AddActor(ha Actor) {
	if hi.actors == nil {
		hi.actors = map[string]Actor{}
	}
	hi.actors[ha.ID()] = ha
}

// GetActor get actor by id
func (hi *ActorContainer) GetActor(id string) Actor {
	return hi.actors[id]
}

// parseEventId parse event id to get module id and action name
func (hi *ActorContainer) parseEventID(id string, rb bool) (string, string) {
	fields := strings.Split(id, ".")
	mid := fields[0]
	aname := "On"
	for i := 1; i < len(fields); i++ {
		aname = aname + strings.Title(fields[i])
	}
	if rb {
		aname = aname + "Rollback"
	}
	return mid, aname
}

func (hi *ActorContainer) getActorMethod(m, a string) *reflect.Value {
	ha, ok := hi.actors[m]
	if !ok {
		log.Debug("Actor '%s' was not found", m)
		return nil
	}
	v := reflect.ValueOf(ha).MethodByName(a)
	if !v.IsValid() {
		log.Debug("Actor method '%s.%s' was not found:", m, a)
		return nil
	}
	return &v
}

func (hi *ActorContainer) invokeActorMethod(caller *reflect.Value, req interface{}, rsp interface{}) error {
	params := []reflect.Value{
		reflect.ValueOf(req),
		reflect.ValueOf(rsp),
	}
	ret := caller.Call(params)
	err := ret[0].Interface()
	if err != nil {
		return err.(error)
	}
	return nil
}

// OnEvent handle event
func (hi *ActorContainer) OnEvent(ctx context.Context, req *eproto.EventRequest, rsp *eproto.EventResponse) error {
	st := time.Now()
	// allocate map for result
	rsp.Results = map[string]string{}

	m, a := hi.parseEventID(req.Event.Id, false)
	defer log.Profiling(st, "OnEvent:%s.%s", m, a)
	log.Debug("Trying to call actor:%s.%s", m, a)
	caller := hi.getActorMethod(m, a)
	if caller != nil {
		return hi.invokeActorMethod(caller, req, rsp)
	}
	log.Debug("Actor %s.%s was not implemented yet", m, a)
	return nil
}

// OnRollback rollback event
func (hi *ActorContainer) OnRollback(ctx context.Context, req *eproto.RollbackRequest, rsp *eproto.RollbackResponse) error {
	st := time.Now()
	m, a := hi.parseEventID(req.Event.Id, true)
	defer log.Profiling(st, "OnRollback:%s.%s", m, a)
	log.Debug("Trying to call actor:%s.%s", m, a)
	caller := hi.getActorMethod(m, a)
	if caller != nil {
		return hi.invokeActorMethod(caller, req, rsp)
	}
	log.Debug("Actor %s.%s was not implemented yet", m, a)
	return nil
}
