package util

import "testing"

func TestObj2Str(t *testing.T) {
	obj := &struct {
		A int
		B string
		C []int
	}{
		A: 1,
		B: "OK",
		C: []int{1, 2, 3},
	}
	str := Obj2Str(obj)
	if str != `{"A":1,"B":"OK","C":[1,2,3]}` {
		t.Fatal("Obj2Str simple struct incorrect")
	}

	type inner struct {
		A int
	}
	obj2 := &struct {
		A *inner
		B []*inner
	}{
		A: &inner{A: 1},
		B: []*inner{
			&inner{A: 2},
			&inner{A: 3},
		},
	}
	str = Obj2Str(obj2)
	if str != `{"A":{"A":1},"B":[{"A":2},{"A":3}]}` {
		t.Fatal("Obj2Str embeded struct incorrect")
	}
	arr := []int64{1, 2, 3}
	str = Obj2Str(arr)
	if str != `[1,2,3]` {
		t.Fatal("Obj2Str with array incorrect")
	}
}

func TestStr2Obj(t *testing.T) {
	str := `{"A":1,"B":"OK","C":[1,2,3]}`
	obj := &struct {
		A int
		B string
		C []int
	}{}
	Str2Obj(str, obj)
	if obj.A != 1 || obj.B != "OK" {
		t.Fatal("Str2Obj simple struct incorrect")
	}

	str = `{"A":{"A":1},"B":[{"A":2},{"A":3}]}`
	type inner struct {
		A int
	}
	obj2 := &struct {
		A *inner
		B []*inner
	}{}
	Str2Obj(str, obj2)
	if obj2.A.A != 1 || len(obj2.B) != 2 || obj2.B[1].A != 3 {
		t.Fatal("Str2Obj embeded struct incorrect")
	}

	str = `[{"A":2},{"A":3}]`
	obj3 := []*inner{}
	Str2Obj(str, &obj3)
	if len(obj3) != 2 || obj3[1].A != 3 {
		t.Fatal("Str2Obj struct array incorrect")
	}
}
