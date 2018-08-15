package engine

import (
	eerr "github.com/go-lego/engine/error"
	"github.com/go-lego/engine/event"
	"github.com/go-lego/engine/tx"
)

// Account interface
type Account interface {
	ID() int64
}

// Engine struct,
// Hold dispatcher and transaction
type Engine struct {
	dispatcher  event.Dispatcher
	transaction tx.Transaction
}

// NewEngine create new engine instance by specifying dispatcher
func NewEngine(d event.Dispatcher) *Engine {
	return &Engine{
		dispatcher: d,
	}
}

// NewContext create new context for the engine
func (e *Engine) NewContext() *Context {
	return &Context{
		engine: e,
		values: map[string]interface{}{},
	}
}

// RaiseEvent raise new event
func (e *Engine) RaiseEvent(ent *event.Event) {
	e.dispatcher.Dispatch(ent)
}

// AddResult add result
func (e *Engine) AddResult(key, value string) {
	e.dispatcher.AddResult(key, value)
}

// Result get result by key
func (e *Engine) Result(key string) string {
	return e.dispatcher.Result(key)
}

// Results get all results
func (e *Engine) Results() map[string]string {
	return e.dispatcher.Results()
}

// AddError add a new error
func (e *Engine) AddError(err *eerr.Error) {
	e.dispatcher.AddError(err)
}

// HasError check if engine has an error
func (e *Engine) HasError() bool {
	return e.dispatcher.HasError()
}

// Error get error by index
func (e *Engine) Error(index int) *eerr.Error {
	return e.dispatcher.Error(index)
}

// StartTransaction start engine transaction
func (e *Engine) StartTransaction() {
	if e.transaction != nil {
	}
}

// EndTransaction end engine transaction
func (e *Engine) EndTransaction() {
	if e.transaction != nil {
	}
}

// DB get database connection by specifying domain name
func (e *Engine) DB(domain string) {

}

// Cache get cache connection by specifying domain name
func (e *Engine) Cache(domain string) {

}
