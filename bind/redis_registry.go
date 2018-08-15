package bind

import (
	"time"

	"github.com/garyburd/redigo/redis"
	"github.com/go-lego/engine/log"
	"github.com/go-lego/engine/util"
)

// RedisRegistry redis registry
type RedisRegistry struct {
	prefix string
	pool   *redis.Pool
}

// NewRedisRegistry create redis registry
func NewRedisRegistry(address, prefix string) *RedisRegistry {
	log.Debug("Trying to create redis registry for binding ... (%s)", address)
	p := &redis.Pool{
		MaxIdle:     5,
		MaxActive:   0,
		IdleTimeout: 2 * time.Minute,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", address)
		},
	}
	c := p.Get()
	defer c.Close()
	if _, err := c.Do("PING"); err != nil {
		log.Fatal("Failed to create redis regitry for binding: %s", err)
	}
	return &RedisRegistry{
		prefix: prefix,
		pool:   p,
	}
}

// GetAll get all binding elements
func (r *RedisRegistry) GetAll() map[string][]*Element {
	c := r.pool.Get()
	defer c.Close()
	ret := map[string][]*Element{}
	kvs, err := redis.StringMap(c.Do("HGETALL", r.prefix))
	if err != nil {
		log.Error("Failed to get binding elements from redis: %s", err)
		return ret
	}
	for k, v := range kvs {
		es := []*Element{}
		util.Str2Obj(v, &es)
		ret[k] = es
	}
	return ret
}

// Add add binding elements
func (r *RedisRegistry) Add(ns string, els []*Element) {
	c := r.pool.Get()
	defer c.Close()
	if _, err := c.Do("HSET", r.prefix, ns, util.Obj2Str(els)); err != nil {
		log.Error("Failed to set binding elements to redis for: %s, error: %s", ns, err)
	}
}
