package event

import (
	eerr "github.com/go-lego/engine/error"
)

// Dispatcher interface, used for dispatching event
type Dispatcher interface {
	// Dispatch event
	Dispatch(ent *Event)

	// AddResult add result
	AddResult(key, value string)

	// Result get result with key
	Result(key string) string

	// Results get all results
	Results() map[string]string

	// AddError add error
	AddError(e *eerr.Error)

	// HasError check if got an error
	HasError() bool

	// Error get single error by index
	Error(index int) *eerr.Error
}
