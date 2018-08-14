package srv

import (
	"fmt"
	"testing"
)

type MyService struct {
}

func (s *MyService) ID() string {
	return "my"
}

func (s *MyService) OnTest1() {
	fmt.Println("OnTest1 is called")
}

func (s *MyService) OnTodoSomething() {
	fmt.Println("OnTodoSomething is called")
}

func (s *MyService) KKK() {

}

func TestGetServiceEvents(t *testing.T) {
	ehs := GetServiceEvents("test", &MyService{})
	if len(ehs) != 2 {
		t.Fatal("handler count was expected to 2")
	}
	if ehs[0].Name != "test.my#OnTest1" || ehs[1].Name != "test.my#OnTodoSomething" {
		t.Fatal("handler name incorrect")
	}
}
