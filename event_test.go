package engine

import (
	"testing"

	eproto "github.com/go-lego/engine/proto"
)

func TestNewEvent(t *testing.T) {
	e := NewEvent("test", 10, nil)
	if e.Id != "test" {
		t.Fatal("Event id was expected to be 'test', but:", e.Id)
	}
	if e.Sender != 10 {
		t.Fatal("Event sender was expected to be 10, but:", e.Sender)
	}
	if e.Parent != nil {
		t.Fatal("Event parent was expected to be nil")
	}

	se := NewEvent("test2", 11, e)
	if se.Id != "test2" {
		t.Fatal("Event id was expected to be 'test2', but:", se.Id)
	}
	if se.Sender != 11 {
		t.Fatal("Event sender was expected to be 11, but:", se.Sender)
	}
	if se.Parent == nil {
		t.Fatal("Event parent was expected to be non-nil")
	}
	if se.Parent.Id != "test" {
		t.Fatal("Parent event id was expected to be 'test', but:", e.Parent.Id)
	}
	if se.GetParent().Id != "test" {
		t.Fatal("GetParent() event id was expected to be 'test', but:", e.Parent.Id)
	}
}

func TestNewEventFromProto(t *testing.T) {
	te := &eproto.Event{Id: "test", Sender: 10}
	e := NewEventFromProto(te)
	if e.Id != "test" {
		t.Fatal("Event id was expected to be 'test', but:", e.Id)
	}
	if e.Sender != 10 {
		t.Fatal("Event sender was expected to be 10, but:", e.Sender)
	}
	if e.Parent != nil {
		t.Fatal("Event parent was expected to be nil")
	}
	ste := &eproto.Event{Id: "test2", Sender: 11, Parent: te}
	se := NewEventFromProto(ste)
	if se.Id != "test2" {
		t.Fatal("Event id was expected to be 'test2', but:", se.Id)
	}
	if se.Sender != 11 {
		t.Fatal("Event sender was expected to be 11, but:", se.Sender)
	}
	if se.Parent == nil {
		t.Fatal("Event parent was expected to be non-nil")
	}
	if se.Parent.Id != "test" {
		t.Fatal("Parent event id was expected to be 'test', but:", e.Parent.Id)
	}
}

func TestSetDataAndGetData(t *testing.T) {
	type tt struct {
		Name string
		Age  int
	}
	e := NewEvent("test", 10, nil)
	e.SetData("intval", 10)
	e.SetData("floatval", 10.1)
	e.SetData("strval", "hello")
	e.SetData("objval", &tt{Name: "tt", Age: 100})
	if e.Data["intval"] != "10" {
		t.Fatal("SetData with intval incorrect:", e.Data["intval"])
	}
	if e.Data["floatval"] != "10.100000" {
		t.Fatal("SetData with floatval incorrect", e.Data["floatval"])
	}
	if e.Data["strval"] != "hello" {
		t.Fatal("SetData with strval incorrect", e.Data["strval"])
	}
	if e.Data["objval"] != `{"Name":"tt","Age":100}` {
		t.Fatal("SetData with intval incorrect:", e.Data["objval"])
	}

	if e.GetDataAsInt("intval") != 10 {
		t.Fatal("GetDataAsInt incorrect:", e.GetDataAsInt("intval"))
	}
	if e.GetDataAsFloat32("floatval") != 10.1 {
		t.Fatal("GetDataAsInt incorrect:", e.GetDataAsInt("floatval"))
	}
	if e.GetDataAsString("strval") != "hello" {
		t.Fatal("GetDataAsInt incorrect:", e.GetDataAsInt("strval"))
	}
	o := &tt{}
	if err := e.GetDataAsObject("objval", o); err != nil {
		t.Fatal("GetDataAsObject should returns nil")
	}
	if o.Name != "tt" || o.Age != 100 {
		t.Fatal("GetDataAsObject incorrect:", o)
	}
}

func TestEventGetParent(t *testing.T) {
	e := NewEvent("test", 0, nil)
	if e.Parent != nil {
		t.Fatal("e.Parent should be nil")
	}
	if e.GetParent() != nil {
		t.Fatal("e.GetParent() should be nil")
	}
}
