package eds

import (
	"context"

	eproto "github.com/go-lego/engine/eds/proto"
	gerr "github.com/go-lego/engine/error"
	"github.com/go-lego/engine/log"
)

// EDS is Event-Driven-System

var (

	// PriorityLow low
	PriorityLow = 1

	// PriorityMedium medium
	PriorityMedium = 128

	// PriorityHigh high
	PriorityHigh = 255
)

type handlerEntry struct {
	id       string
	client   eproto.HandlerService
	priority int
	async    bool
}

type dispatchHistory struct {
	h *handlerEntry     // handler
	e *Event            // event
	r map[string]string // result
}

var (
	mapping = map[string][]*handlerEntry{}
)

// Bind handler to event
func Bind(eid string, hid string, hs eproto.HandlerService, p int, a bool) {
	handlers, ok := mapping[eid]
	if !ok {
		handlers = []*handlerEntry{}
	}
	handlers = append(handlers, &handlerEntry{id: hid, client: hs, priority: p, async: a})
	mapping[eid] = handlers
}

// UnbindAll unbind all
func UnbindAll() {
	mapping = map[string][]*handlerEntry{}
}

// Dispatcher event dispatcher
type Dispatcher struct {
	events    []*Event
	results   map[string]string
	running   bool
	histories []*dispatchHistory
}

// NewDispatcher create new dispatcher
func NewDispatcher() *Dispatcher {
	return &Dispatcher{
		events:    []*Event{},
		results:   map[string]string{},
		running:   false,
		histories: []*dispatchHistory{},
	}
}

func (d *Dispatcher) handleEvent(e *Event) error {
	handlers, ok := mapping[e.Id]
	if !ok {
		log.Debug("No handler found for event:%s", e.Id)
		return nil
	}
	ctx := context.Background()
	req := &eproto.EventRequest{
		Event: e.Event,
	}
	for _, h := range handlers {
		log.Debug("Handler(%s) is handling event:%s", h.id, e.Id)
		if h.async { // async handler, do not need the response
			go func(ih *handlerEntry, ireq *eproto.EventRequest) {
				ih.client.OnEvent(context.Background(), ireq)
				log.Debug("Handler(%s) handle event completed:%s", ih.id, ireq.Event.Id)
			}(h, req)
			continue
		}
		rsp, err := h.client.OnEvent(ctx, req)
		if err != nil {
			log.Debug("Handler(%s) got system error:%s", h.id, err)
			return gerr.NewSystemError(1, err.Error())
		}
		if rsp.Code != 0 {
			log.Debug("Handler(%s) got logical error:code=%d, message=%s", h.id, rsp.Code, rsp.Message)
			return gerr.New(rsp.Code, rsp.Message)
		}
		// record history(for roll back)
		d.histories = append(d.histories, &dispatchHistory{
			h: h,
			e: e,
			r: rsp.Results,
		})
		// merge results
		for k, v := range rsp.Results {
			d.results[k] = v
		}
		// dispatch raised events
		for _, t := range rsp.Events {
			d.Dispatch(NewEventFromProto(t))
		}
		log.Debug("Handler(%s) handle event completed:%s", h.id, e.Id)
	}
	return nil
}

func (d *Dispatcher) rollback() {
	ctx := context.Background()
	l := len(d.histories)
	for i := 0; i < l; i++ {
		t := d.histories[l-1-i]
		req := &eproto.RollbackRequest{
			Event:   t.e.Event,
			Results: t.r,
		}
		log.Debug("Handler(%s) is rolling back event:%s", t.h.id, req.Event.Id)
		t.h.client.OnRollback(ctx, req)
		log.Debug("Handler(%s) is roll back event completed:%s", t.h.id, req.Event.Id)
	}
}

// Dispatch event
func (d *Dispatcher) Dispatch(e *Event) error {
	log.Debug("Dispatching event (%s) ...", e.Id)
	log.Debug("%s", e)
	d.events = append(d.events, e)
	if !d.running {
		d.running = true
		for i := 0; i < len(d.events); i++ {
			if err := d.handleEvent(d.events[i]); err != nil {
				d.rollback()
				return err
			}
			// logger.Debug("Remove handled event:%s", d.events[i].Id)
			d.events = append(d.events[:i], d.events[i+1:]...)
			i--
			// logger.Debug("Event list count:%d", len(d.events))
		}
		d.running = false
	}
	return nil
}

// Result get result by key
func (d *Dispatcher) Result(key string) string {
	v, ok := d.results[key]
	if !ok {
		return ""
	}
	return v
}

// Results get all result
func (d *Dispatcher) Results() map[string]string {
	return d.results
}
