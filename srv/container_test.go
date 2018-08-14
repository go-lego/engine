package srv

import (
	"fmt"
	"testing"
)

type NewService struct {
}

func (s *NewService) ID() string {
	return "my"
}

func (s *NewService) OnTest1() {
	fmt.Println("OnTest1 is called")
}

func TestAddServices(t *testing.T) {
	AddServices("test", 2, true, &MyService{})
	AddServices("test1", 1, true, &NewService{})
	arr, ok := cachedEventhandlers["my.test1"]
	if !ok || len(arr) != 2 {
		t.Fatal("my.test1 handler count incorrect")
	}
	// t.Error(arr[0].Name, arr[1].Name)
	if arr[0].Name != "test1.my#OnTest1" {
		t.Fatal("event handler sort incorrect")
	}
	arr, ok = cachedEventhandlers["my.todo.something"]
	if !ok || len(arr) != 1 {
		t.Fatal("my.todo.something handler count incorrect")
	}
}
