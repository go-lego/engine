package srv

import (
	"github.com/go-lego/engine"
	eerr "github.com/go-lego/engine/error"
	eproto "github.com/go-lego/engine/proto"
)

// Dispatcher for srv
type Dispatcher struct {
	response *eproto.EventResponse
}

// NewDispatcher create new dispatcher for srv
func NewDispatcher(rsp *eproto.EventResponse) *Dispatcher {
	return &Dispatcher{
		response: rsp,
	}
}

// Dispatch event
func (d *Dispatcher) Dispatch(ent *engine.Event) {
	d.response.Events = append(d.response.Events, ent.Event)
}

// AddResult add result
func (d *Dispatcher) AddResult(key, value string) {
	d.response.Results[key] = value
}

// Result get result with key
func (d *Dispatcher) Result(key string) string {
	return d.response.Results[key]
}

// Results get all results
func (d *Dispatcher) Results() map[string]string {
	return d.response.Results
}

// AddError add error
func (d *Dispatcher) AddError(e *eerr.Error) {
	d.response.Code = e.Code
	d.response.Message = e.Message
}

// HasError check if got an error
func (d *Dispatcher) HasError() bool {
	return d.response.Code != 0
}

// Error get single error by index
func (d *Dispatcher) Error() *eerr.Error {
	return eerr.New(d.response.Code, d.response.Message)
}
