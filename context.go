package engine

// Context engine context
type Context struct {
	engine *Engine
	values map[string]interface{}
}

// Engine get engine
func (c *Context) Engine() *Engine {
	return c.engine
}

// WithValue add value to context
func (c *Context) WithValue(key string, value interface{}) *Context {
	c.values[key] = value
	return c
}

// Value get value from context
func (c *Context) Value(key string) interface{} {
	return c.values[key]
}
