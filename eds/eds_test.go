package eds

import (
	"context"
	"errors"
	"testing"

	event "github.com/go-lego/engine/eds/proto"
	"github.com/micro/go-micro/client"
)

type testHandler1 struct {
}

func (th *testHandler1) OnEvent(ctx context.Context, in *event.EventRequest, opts ...client.CallOption) (*event.EventResponse, error) {
	rsp := &event.EventResponse{
		Results: map[string]string{},
	}
	if in.Event.Id == "test" {
		if in.Event.Data["value"] == "1" || in.Event.Data["value"] == "2" {
			rsp.Results["test1.value"] = "test1"
			counter++
		}
	}

	return rsp, nil
}
func (th *testHandler1) OnRollback(ctx context.Context, in *event.RollbackRequest, opts ...client.CallOption) (*event.RollbackResponse, error) {
	counter--
	return nil, nil
}

type testHandler2 struct {
}

func (th *testHandler2) OnEvent(ctx context.Context, in *event.EventRequest, opts ...client.CallOption) (*event.EventResponse, error) {
	rsp := &event.EventResponse{
		Results: map[string]string{},
	}
	if in.Event.Id == "test" {
		if in.Event.Data["value"] == "1" {
			rsp.Results["test2.value"] = "test2"
			counter++
		}
		if in.Event.Data["value"] == "2" {
			return nil, errors.New("test")
		}
	} else if in.Event.Id == "test2" {
		t := NewEvent("test", in.Event.Sender, NewEventFromProto(in.Event))
		t.Data["value"] = "1"
		rsp.Results["test2.tmp"] = "hello world"
		rsp.Events = append(rsp.Events, t.Event)
	}
	return rsp, nil
}

func (th *testHandler2) OnRollback(ctx context.Context, in *event.RollbackRequest, opts ...client.CallOption) (*event.RollbackResponse, error) {
	return nil, nil
}

var (
	h1      = new(testHandler1)
	h2      = new(testHandler2)
	counter = 0
)

func setUp(t *testing.T) {
	Bind("test", "test1", h1, PriorityLow, false)
	Bind("test", "test2", h2, PriorityLow, false)
	Bind("test2", "test2", h2, PriorityLow, false)
	counter = 0
}

func tearDown(t *testing.T) {
	// UnBind("test", h1, h2)
	// UnBind("test2", h2)
	UnbindAll()
}

func TestBind(t *testing.T) {
	setUp(t)
	defer tearDown(t)
	if len(mapping) != 2 {
		t.Fatalf("event mapping length was expected to 2, but:%d", len(mapping))
	}
	hs, ok := mapping["test"]
	if !ok {
		t.Fatalf("'test' event not exists")
	}
	if len(hs) != 2 {
		t.Fatalf("'test' event handlers should be 2")
	}
	if hs[0].id != "test1" {
		t.Fatalf("'test' first handler id was expected to 'test1', but:%s", hs[0].id)
	}
	hs, ok = mapping["test2"]
	if !ok {
		t.Fatalf("'test2' event not exists")
	}
	if len(hs) != 1 {
		t.Fatalf("'test2' event handlers should be 1")
	}
	if hs[0].id != "test2" {
		t.Fatalf("'test' first handler id was expected to 'test2', but:%s", hs[0].id)
	}
}

// func TestUnBind(t *testing.T) {
// 	setUp(t)
// 	defer tearDown(t)
// 	UnBind("test", h1)
// 	hs, _ := mapping["test"]
// 	if len(hs) != 1 {
// 		t.Fatalf("'test' event handler length should be 1")
// 	}
// }

func TestDispatchSingleEvent(t *testing.T) {
	setUp(t)
	defer tearDown(t)

	e := NewEvent("test", 0, nil)
	e.Data["value"] = "1"
	d := NewDispatcher()
	err := d.Dispatch(e)
	// t.Error(d.Results())
	if err != nil {
		t.Fatalf("no error was expected, but:%s", err)
	}
	if d.Result("test1.value") != "test1" {
		t.Fatal("test1.value was incorrect")
	}
	if d.Result("test2.value") != "test2" {
		t.Fatal("test2.value was incorrect")
	}
	if counter != 2 {
		t.Fatalf("counter was expected to 2, but:%d", counter)
	}
}

func TestDispatchEventAndRaiseEvent(t *testing.T) {
	setUp(t)
	defer tearDown(t)

	e := NewEvent("test2", 0, nil)
	e.Data["value"] = "1"
	d := NewDispatcher()
	err := d.Dispatch(e)
	if err != nil {
		t.Fatalf("no error was expected, but:%s", err)
	}
	if d.Result("test1.value") != "test1" {
		t.Fatal("test1.value was unexpected")
	}
	if d.Result("test2.value") != "test2" {
		t.Fatal("test2.value was unexpected")
	}
	// t.Error(d.Results())
}

func TestDispatchEventWithError(t *testing.T) {
	setUp(t)
	defer tearDown(t)

	e := NewEvent("test", 0, nil)
	e.Data["value"] = "2"
	d := NewDispatcher()
	err := d.Dispatch(e)
	if err.Error() != `{"code":1,"message":"System error:test"}` {
		t.Fatalf("'test' error was expected, but:%s", err)
	}
	rs := d.Results()
	if len(rs) != 1 {
		t.Fatalf("result length was expected to 1, but:%d", len(rs))
	}
	if counter != 0 {
		t.Fatalf("counter was expected to 0, but:%d", counter)
	}
}

func TestNewEventWithParent(t *testing.T) {
	e1 := NewEvent("test", 0, nil)
	e1.SetMeta("m1", "hello")
	e1.SetMeta("m2", "world")
	e2 := NewEvent("child", 0, e1)
	if e2.Meta["m1"] != "hello" {
		t.Fatalf("Meta not inherit from parent")
	}
	if m := e2.Meta["none"]; m != "" {
		t.Fatalf("no meta not match:%s", m)
	}
}
