package api

import (
	"testing"
)

func TestConvertNameToPath(t *testing.T) {
	if _, err := convertNameToPath("go.lego.tt"); err == nil {
		t.Fatal("Name error was expected, but nil")
	}
	if p, _ := convertNameToPath("go.lego.api.mobile"); p != "/mobile" {
		t.Fatalf("Failed:go.lego.api.mobile !=> /mobile(%s)", p)
	}
	if _, err := convertNameToPath("go.lego.api.mobile.tt"); err == nil {
		t.Fatal("Name error was expected, but nil")
	}
	if p, _ := convertNameToPath("go.lego.api.v1.mobile"); p != "/v1/mobile" {
		t.Fatalf("Failed:go.lego.api.v1.mobile !=> /v1/mobile(%s)", p)
	}
}
