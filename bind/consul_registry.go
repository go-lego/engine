package bind

import (
	"fmt"
	"strings"

	"github.com/go-lego/engine/log"
	"github.com/go-lego/engine/util"
	capi "github.com/hashicorp/consul/api"
)

// ConsulRegistry consul registry
type ConsulRegistry struct {
	client *capi.KV
	prefix string
}

// NewConsulRegistry create new registry
func NewConsulRegistry(address, prefix string) *ConsulRegistry {
	cfg := capi.DefaultConfig()
	cfg.Address = address
	c, err := capi.NewClient(cfg)
	if err != nil {
		log.Fatal("Failed to create consul registry for binding: %s", err)
	}
	return &ConsulRegistry{
		client: c.KV(),
		prefix: prefix,
	}
}

// GetAll get all binding elements
func (r *ConsulRegistry) GetAll() map[string][]*Element {
	ret := map[string][]*Element{}
	kvs, _, err := r.client.List(r.prefix, nil)
	if err != nil {
		log.Error("Failed to get binding elements from consul:%s", err)
		return ret
	}
	for _, kv := range kvs {
		name := strings.Replace(kv.Key, r.prefix, "", 1)
		value := string(kv.Value)
		es := []*Element{}
		util.Str2Obj(value, &es)
		ret[name] = es
	}
	return ret
}

// Add binding elements
func (r *ConsulRegistry) Add(ns string, els []*Element) {
	k := fmt.Sprintf("%s%s", r.prefix, ns)
	if _, err := r.client.Put(&capi.KVPair{Key: k, Value: []byte(util.Obj2Str(els))}, nil); err != nil {
		log.Error("Failed to set binding elements to consul for: %s, error: %s", ns, err)
	}
}
