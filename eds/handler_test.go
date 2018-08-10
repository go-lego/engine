package eds

import (
	"errors"
	"testing"

	eproto "github.com/go-lego/engine/eds/proto"
)

func TestActorContainerParseEventId(t *testing.T) {
	hi := &ActorContainer{}
	m, a := hi.parseEventID("account.register", false)
	if m != "account" {
		t.Fatal("Module id was incorrect")
	}
	if a != "OnRegister" {
		t.Fatal("Action name was incorrect")
	}
	m, a = hi.parseEventID("account.register", true)
	if m != "account" {
		t.Fatal("Module id was incorrect")
	}
	if a != "OnRegisterRollback" {
		t.Fatal("Action name was incorrect")
	}
}

type testHandlerActor struct {
}

func (k *testHandlerActor) ID() string {
	return "account"
}
func (k *testHandlerActor) OnRegister(req *eproto.EventRequest, rsp *eproto.EventResponse) error {
	req.Event.Data["counter"] = "1"
	return nil
}
func (k *testHandlerActor) OnRegisterRollback(req *eproto.RollbackRequest, rsp *eproto.RollbackResponse) error {
	req.Event.Data["counter"] = "0"
	return nil
}
func (k *testHandlerActor) OnRegistered(req *eproto.EventRequest, rsp *eproto.EventResponse) error {
	return errors.New("test")
}
func TestActorContainerOnEvent(t *testing.T) {
	hi := &ActorContainer{}
	hi.AddActor(&testHandlerActor{})
	e := NewEvent("account.register", 0, nil)
	e.Data["counter"] = "0"
	err := hi.OnEvent(nil, &eproto.EventRequest{Event: e.Event}, &eproto.EventResponse{})
	if err != nil {
		t.Fatal("no error was expected")
	}
	if e.Data["counter"] != "1" {
		t.Fatalf("counter was expected to 1, but:%s", e.Data["counter"])
	}
	e = NewEvent("account.registered", 0, nil)
	err = hi.OnEvent(nil, &eproto.EventRequest{Event: e.Event}, &eproto.EventResponse{})
	if err == nil {
		t.Fatalf("'test' error was expected, but nil")
	}
}

func TestActorContainerOnEventNotImplemented(t *testing.T) {
	hi := &ActorContainer{}
	hi.AddActor(&testHandlerActor{})
	e := NewEvent("account.test", 0, nil)
	e.Data["counter"] = "0"
	err := hi.OnEvent(nil, &eproto.EventRequest{Event: e.Event}, &eproto.EventResponse{})
	if err != nil {
		t.Fatal("no error was expected")
	}
}

func TestActorContainerOnRollback(t *testing.T) {
	hi := &ActorContainer{}
	hi.AddActor(&testHandlerActor{})
	e := NewEvent("account.register", 0, nil)
	e.Data["counter"] = "1"
	err := hi.OnRollback(nil, &eproto.RollbackRequest{Event: e.Event}, &eproto.RollbackResponse{})
	if err != nil {
		t.Fatal("no error was expected")
	}
	if e.Data["counter"] != "0" {
		t.Fatalf("counter was expected to 0, but:%s", e.Data["counter"])
	}
}

func TestActorContainerOnRollbackNotImplemented(t *testing.T) {
	hi := &ActorContainer{}
	hi.AddActor(&testHandlerActor{})
	e := NewEvent("account.test", 0, nil)
	e.Data["counter"] = "1"
	err := hi.OnRollback(nil, &eproto.RollbackRequest{Event: e.Event}, &eproto.RollbackResponse{})
	if err != nil {
		t.Fatal("no error was expected")
	}
}
