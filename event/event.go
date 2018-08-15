package event

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"

	"enjoypiano.cn/lynxpro/common/logger"
	"github.com/go-lego/engine/log"
	eproto "github.com/go-lego/engine/proto"
)

// Event event.
// Inherited from proto.Event and provides some utility functions.
type Event struct {
	*eproto.Event
}

// NewEvent create new event
func NewEvent(id string, sender int64, parent *Event) *Event {
	var pe *eproto.Event
	var md map[string]string
	if parent != nil {
		pe = parent.Event
		md = parent.Meta
	}
	return &Event{
		&eproto.Event{
			Id:     id,
			Sender: sender,
			Parent: pe,
			Data:   map[string]string{},
			Meta:   md,
		},
	}
}

// NewEventFromProto create new event from proto event
func NewEventFromProto(e *eproto.Event) *Event {
	return &Event{
		Event: e,
	}
}

// SetData set data
func (e *Event) SetData(key string, value interface{}) *Event {
	switch value.(type) {
	case string:
		e.Data[key] = value.(string)
	case int:
		e.Data[key] = fmt.Sprintf("%d", value.(int))
	case int32:
		e.Data[key] = fmt.Sprintf("%d", value.(int32))
	case int64:
		e.Data[key] = fmt.Sprintf("%d", value.(int64))
	case float32:
		e.Data[key] = fmt.Sprintf("%f", value.(float32))
	case float64:
		e.Data[key] = fmt.Sprintf("%f", value.(float64))
	default:
		// logger.Debug("Other type")
		data, err := json.Marshal(value)
		if err != nil {
			log.Error("Set event data failed to marshal event data (%s) to json string:%s", key, err)
			return e
		}
		e.Data[key] = string(data)
	}
	return e
}

// GetDataAsInt get int data
func (e *Event) GetDataAsInt(key string) int {
	v, ok := e.Data[key]
	if !ok {
		log.Info("Event data %s does not exist", key)
		return 0
	}
	i, err := strconv.Atoi(v)
	if err != nil {
		log.Info("Failed to convert event data %s=%s to int:%s", key, v, err)
	}
	return i
}

// GetDataAsInt32 get int32 data
func (e *Event) GetDataAsInt32(key string) int32 {
	v, ok := e.Data[key]
	if !ok {
		log.Info("Event data %s does not exist", key)
		return 0
	}
	i, err := strconv.Atoi(v)
	if err != nil {
		log.Info("Failed to convert event data %s=%s to int32:%s", key, v, err)
	}
	return int32(i)
}

// GetDataAsInt64 get int64 data
func (e *Event) GetDataAsInt64(key string) int64 {
	v, ok := e.Data[key]
	if !ok {
		log.Info("Event data %s does not exist", key)
		return 0
	}
	i, err := strconv.ParseInt(v, 0, 64)
	if err != nil {
		log.Info("Failed to convert event data %s=%s to int64:%s", key, v, err)
	}
	return i
}

// GetDataAsFloat32 get float32 data
func (e *Event) GetDataAsFloat32(key string) float32 {
	v, ok := e.Data[key]
	if !ok {
		log.Info("Event data %s does not exist", key)
		return 0
	}
	i, err := strconv.ParseFloat(v, 64)
	if err != nil {
		log.Info("Failed to convert event data %s=%s to float32:%s", key, v, err)
	}
	return float32(i)
}

// GetDataAsFloat64 get float64 data
func (e *Event) GetDataAsFloat64(key string) float64 {
	v, ok := e.Data[key]
	if !ok {
		log.Info("Event data %s does not exist", key)
		return 0
	}
	i, err := strconv.ParseFloat(v, 64)
	if err != nil {
		log.Info("Failed to convert event data %s=%s to float64:%s", key, v, err)
	}
	return i
}

// GetDataAsString get string data
func (e *Event) GetDataAsString(key string) string {
	v, ok := e.Data[key]
	if !ok {
		log.Info("Event data %s does not exist", key)
		return ""
	}
	return v
}

// GetDataAsObject get object data
func (e *Event) GetDataAsObject(key string, out interface{}) error {
	v, ok := e.Data[key]
	if !ok {
		logger.Info("Event data %s does not exist", key)
		return errors.New("not exist")
	}
	return json.Unmarshal([]byte(v), out)
}

// SetMeta set meta
func (e *Event) SetMeta(key string, value string) *Event {
	if e.Meta == nil {
		e.Meta = map[string]string{}
	}
	e.Meta[key] = value
	return e
}

// GetMeta get meta by key
func (e *Event) GetMeta(key string) string {
	if e.Meta == nil {
		return ""
	}
	return e.Meta[key]
}

// GetParent get parent event
func (e *Event) GetParent() *Event {
	if e.Parent == nil {
		return nil
	}
	return &Event{
		e.Parent,
	}
}
