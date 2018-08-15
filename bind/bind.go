package bind

import (
	"fmt"
	"sort"
	"sync"
	"time"

	"github.com/go-lego/engine/log"
	eproto "github.com/go-lego/engine/proto"
	"github.com/micro/go-micro/broker"
	"github.com/micro/go-micro/client"
)

var (
	// Topic default top
	Topic = "go.lego.topic.bind.srv"

	services = map[string]eproto.EventService{}

	mapping = map[string][]*Handler{}
)

// Handler struct
type Handler struct {
	ID       string
	Service  eproto.EventService
	Priority int
}

// Element binding element
type Element struct {
	ID       string `json:"id"` // event id
	Priority int
}

// Watch service booting to update bindings (for API)
func Watch() {
	log.Debug("Watching event bindings ...")
	buildBindings()
	go func() {
		for {
			if err := broker.Init(); err != nil {
				log.Info("Failed to initialize broker")
			} else if err := broker.Connect(); err != nil {
				log.Info("Failed to connect to broker")
			} else if _, err := broker.Subscribe(Topic, func(pub broker.Publication) error {
				ns := string(pub.Message().Body)
				if _, ok := services[ns]; !ok {
					log.Info("Got new service(%s), trying to build bindings...", ns)
					buildBindings()
				}
				return nil
			}); err != nil {
				log.Info("Failed to subscribe broker topic:%s", Topic)
			} else {
				break
			}
			log.Debug("Broker failed, will retry in 10 seconds")
			time.Sleep(10 * time.Second)
		}
	}()
}

// Report bindings (for SRV)
func Report(ns string, data []*Element) {
	log.Debug("Reporting service event binding ...")
	registry.Add(ns, data)
	log.Debug("Notifying API server with srv namespace: %s", ns)
	if err := broker.Publish(Topic, &broker.Message{Body: []byte(ns)}); err != nil {
		log.Error("Failed to publish srv namespace:%s", err)
	}
}

// Registry binding registry
type Registry interface {
	// GetAll get all binding elements.
	// srv-name => [e1, e2, ...]
	GetAll() map[string][]*Element

	Add(ns string, els []*Element)
}

var (

	// DefaultRegistry registry
	registry Registry = NewConsulRegistry("127.0.0.1:8500", "event/binding/") // NewLocalRegistry()

	lock = new(sync.Mutex)
)

// SetRegistry set registry
func SetRegistry(r Registry) {
	registry = r
}

func buildBindings() {
	dummy := map[string][]*Handler{}
	raw := registry.GetAll()
	for ns, els := range raw {
		s, ok := services[ns]
		if !ok {
			s = eproto.NewEventService(ns, client.DefaultClient)
			log.Debug("Add event service: %s", ns)
			services[ns] = s
		}
		for _, el := range els {
			arr, ok := dummy[el.ID]
			if !ok {
				arr = []*Handler{}
			}
			arr = append(arr, &Handler{ID: ns, Service: s, Priority: el.Priority})
			dummy[el.ID] = arr
		}
	}

	for id, arr := range dummy {
		sort.Slice(arr, func(i, j int) bool { return arr[i].Priority < arr[j].Priority })
		dummy[id] = arr
	}

	// dump mapping
	if log.GetLevel() >= log.LevelDebug {
		log.Debug("Binding mapping >>>>>>>>>>>>>>>>>>>>> ")
		for id, arr := range dummy {
			fmt.Printf("%s => [", id)
			for _, a := range arr {
				fmt.Printf("%s,", a.ID)
			}
			fmt.Printf("]\n")
		}
		log.Debug("Binding mapping <<<<<<<<<<<<<<<<<<<<< ")
	}

	lock.Lock()
	mapping = dummy
	lock.Unlock()
}

// GetMapping get bind mapping
func GetMapping() map[string][]*Handler {
	lock.Lock()
	defer lock.Unlock()
	return mapping
}
