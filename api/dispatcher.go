package api

import (
	"context"

	"github.com/go-lego/engine"
	"github.com/go-lego/engine/bind"
	eerr "github.com/go-lego/engine/error"
	"github.com/go-lego/engine/log"
	eproto "github.com/go-lego/engine/proto"
)

// Dispatcher for API
type Dispatcher struct {
	events    []*engine.Event
	results   map[string]string
	running   bool
	histories []*dispatchHistory
	erro      *eerr.Error
	mapping   map[string][]*bind.Handler
}

type dispatchHistory struct {
	h *bind.Handler     // handler
	e *engine.Event     // event
	r map[string]string // result
}

// NewDispatcher create new dispatcher
func NewDispatcher() *Dispatcher {
	return &Dispatcher{
		events:  []*engine.Event{},
		results: map[string]string{},
		running: false,
		mapping: bind.GetMapping(),
	}
}

func (d *Dispatcher) handleEvent(e *engine.Event) *eerr.Error {
	mapping := d.mapping
	handlers, ok := mapping[e.Id]
	if !ok {
		log.Debug("No handler found for event: %s", e.Id)
		return nil
	}
	ctx := context.Background()
	req := &eproto.EventRequest{
		Event: e.Event,
	}
	for _, h := range handlers {
		log.Debug("Handler(%s) is handling event: %s", h.ID, e.Id)
		rsp, err := h.Service.OnEvent(ctx, req)
		if err != nil {
			log.Debug("Handler(%s) got system error: %s", h.ID, err)
			return eerr.NewSystemError(1, err.Error())
		}
		if rsp.Code != 0 {
			log.Debug("Handler(%s) got logical error: code=%d, message=%s", h.ID, rsp.Code, rsp.Message)
			return eerr.New(rsp.Code, rsp.Message)
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
			d.Dispatch(engine.NewEventFromProto(t))
		}
		log.Debug("Handler(%s) handle event completed: %s", h.ID, e.Id)
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
		log.Debug("Handler(%s) is rolling back event: %s", t.h.ID, req.Event.Id)
		t.h.Service.OnRollback(ctx, req)
		log.Debug("Handler(%s) is roll back event completed: %s", t.h.ID, req.Event.Id)
	}
}

// Dispatch event
func (d *Dispatcher) Dispatch(e *engine.Event) {
	log.Debug("Dispatching event (%s) ...", e.Id)
	log.Debug("%s", e)
	d.events = append(d.events, e)
	if !d.running {
		d.running = true
		for i := 0; i < len(d.events); i++ {
			if err := d.handleEvent(d.events[i]); err != nil {
				d.erro = err
				d.rollback()
				return
			}
			// logger.Debug("Remove handled event:%s", d.events[i].Id)
			d.events = append(d.events[:i], d.events[i+1:]...)
			i--
			// logger.Debug("Event list count:%d", len(d.events))
		}
		d.running = false
	}
}

// AddResult add result
func (d *Dispatcher) AddResult(key, value string) {
	d.results[key] = value
}

// Result get result with key
func (d *Dispatcher) Result(key string) string {
	return d.results[key]
}

// Results get all results
func (d *Dispatcher) Results() map[string]string {
	return d.results
}

// AddError add error
func (d *Dispatcher) AddError(e *eerr.Error) {
	d.erro = e
}

// HasError check if got an error
func (d *Dispatcher) HasError() bool {
	return d.erro != nil
}

// Error get single error by index
func (d *Dispatcher) Error() *eerr.Error {
	return d.erro
}
