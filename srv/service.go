package srv

import (
	"fmt"
	"reflect"
	"regexp"
	"strings"
)

// Service interface.
// To handle event abc.def.ghi, need method OnDefGhi.
// To rollback event abc.def.ghi, need method RollbackDefGhi.
// Event function: func(ctx *engine.Context, e *event.Event) error {}
// Rollback function: func(e *event.Event, results map[string]string) {}
type Service interface {
	ID() string
}

// EventHandler service event handler entry
type EventHandler struct {
	ID         string // event ID
	Name       string
	Caller     reflect.Value
	Rollbacker reflect.Value
	Priority   int
	Async      bool
}

var camel = regexp.MustCompile("(^[^A-Z0-9]*|[A-Z0-9]*)([A-Z0-9][^A-Z]+|$)")

func methodToEventID(s string) string {
	var a []string
	for _, sub := range camel.FindAllStringSubmatch(s, -1) {
		if sub[1] != "" {
			a = append(a, sub[1])
		}
		if sub[2] != "" {
			a = append(a, sub[2])
		}
	}
	return strings.ToLower(strings.Join(a, "."))
}

// GetServiceEvents get event id list
func GetServiceEvents(group string, s Service) []*EventHandler {
	ret := []*EventHandler{}
	t := reflect.TypeOf(s)
	v := reflect.ValueOf(s)
	// log.Debug("Method number:%d", t.NumMethod())
	for i := 0; i < t.NumMethod(); i++ {
		m := t.Method(i)
		mn := m.Name
		if strings.Index(mn, "On") == 0 {
			// log.Debug("Method:%s", mn)
			eid := strings.Replace(methodToEventID(mn), "on", s.ID(), 1)
			rbn := fmt.Sprintf("Rollback%s", strings.Replace(mn, "On", "", 1))
			// log.Debug("Event:%s", eid)
			eh := &EventHandler{
				ID:         eid,
				Name:       fmt.Sprintf("%s.%s#%s", group, s.ID(), m.Name),
				Caller:     v.MethodByName(mn),
				Rollbacker: v.MethodByName(rbn),
			}
			// log.Debug("Handler name:%s", eh.Name)
			ret = append(ret, eh)
		}
	}
	return ret
}
