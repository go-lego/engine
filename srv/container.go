package srv

import (
	"context"
	"fmt"
	"reflect"
	"sort"
	"time"

	"github.com/go-lego/engine"
	eerr "github.com/go-lego/engine/error"
	"github.com/go-lego/engine/log"
	eproto "github.com/go-lego/engine/proto"
)

var (
	cachedEventhandlers = map[string][]*EventHandler{}

	// PriorityHigh high
	PriorityHigh = 0

	// PriorityMedium medium
	PriorityMedium = 10

	// PriorityLow low
	PriorityLow = 20
)

// Container service container.
// Container is an implementation of EventService interface.
type Container struct {
}

// NewContainer create new container
func NewContainer() *Container {
	return &Container{}
}

// AddServices add services
func AddServices(group string, priority int, async bool, ss ...Service) {
	log.Debug("Trying to add services: %s", group)
	for _, s := range ss {
		ehs := GetServiceEvents(group, s)
		for _, eh := range ehs {
			eh.Priority = priority
			eh.Async = async
			arr, ok := cachedEventhandlers[eh.ID]
			if !ok {
				arr = []*EventHandler{}
			}
			arr = append(arr, eh)
			sort.Slice(arr, func(i, j int) bool { return arr[i].Priority < arr[j].Priority })
			cachedEventhandlers[eh.ID] = arr
		}
	}

	// dump
	if log.GetLevel() >= log.LevelDebug {
		log.Debug("Cached event handlers: >>>>>>>>>>>>>>>>>>>>>>>")
		for eid, arr := range cachedEventhandlers {
			fmt.Printf("%s => [", eid)
			for _, a := range arr {
				fmt.Printf("%s,", a.Name)
			}
			fmt.Printf("]\n")
		}
		log.Debug("Cached event handlers: <<<<<<<<<<<<<<<<<<<<<<<")
	}
}

// OnEvent entrance of event request
func (c *Container) OnEvent(ctx context.Context, req *eproto.EventRequest, rsp *eproto.EventResponse) error {
	eid := req.Event.Id
	ehs, ok := cachedEventhandlers[eid]
	if !ok {
		log.Info("Cannot find event handler:%s", eid)
		return nil
	}
	rsp.Results = map[string]string{}
	ng := engine.NewEngine(NewDispatcher(rsp))
	e := engine.NewEventFromProto(req.Event)
	params := []reflect.Value{
		reflect.ValueOf(ng),
		reflect.ValueOf(e),
	}
	//ng.StartTransaction()
	for _, eh := range ehs {
		log.Debug("Trying to call: %s", eh.Name)
		if eh.Async { // async, ignore error and results
			go func(h *EventHandler, ps []reflect.Value) {
				h.Caller.Call(ps)
			}(eh, params)
			continue
		}
		t1 := time.Now()
		ret := eh.Caller.Call(params)
		log.Profiling(t1, eh.Name)
		err := ret[0].Interface()
		if err != nil {
			if er, ok := err.(*eerr.Error); ok {
				ng.AddError(er)
			} else {
				ng.AddError(eerr.NewSystemError(1, fmt.Sprintf("%s", err)))
			}
			break
		}
	}
	// ng.EndTransaction()
	return nil
}

// OnRollback entrance of rollback request
func (c *Container) OnRollback(ctx context.Context, req *eproto.RollbackRequest, rsp *eproto.RollbackResponse) error {
	eid := req.Event.Id
	ehs, ok := cachedEventhandlers[eid]
	if !ok {
		log.Info("Cannot find event handler:%s", eid)
		return nil
	}
	ng := engine.NewEngine(nil)
	e := engine.NewEventFromProto(req.Event)
	params := []reflect.Value{
		reflect.ValueOf(ng),
		reflect.ValueOf(e),
		reflect.ValueOf(req.Results),
	}
	//ng.StartTransaction()
	for _, eh := range ehs {
		if !eh.Rollbacker.IsValid() {
			continue
		}
		log.Debug("Trying to rollback: %s", eh.Name)
		ret := eh.Rollbacker.Call(params)
		err := ret[0].Interface()
		if err != nil {
			log.Error("Rollback %s failed:%s", eh.Name, err)
		}
	}
	// ng.EndTransaction()
	return nil
}
