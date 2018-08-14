package bind

// LocalRegistry local registry
type LocalRegistry struct {
	data map[string][]*Element
}

// NewLocalRegistry create new local registry
func NewLocalRegistry() *LocalRegistry {
	return &LocalRegistry{
		data: map[string][]*Element{},
	}
}

// Add binding elements
func (r *LocalRegistry) Add(name string, els []*Element) {
	r.data[name] = els
}

// GetAll get all binding elements
func (r *LocalRegistry) GetAll() map[string][]*Element {
	return r.data
}
