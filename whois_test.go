package ripego

import (
	"reflect"
	"testing"
)

func TestNicProvider(t *testing.T) {

	w := getNicProvider("1.5.5.1")
	t.Log(reflect.TypeOf(w).Name())

	if reflect.TypeOf(w).Name() != "apnic" {
		t.Fatal("Invalid type: apnic")
	}

	w = getNicProvider("177.148.56.7")
	t.Log(reflect.TypeOf(w).Name())

	if reflect.TypeOf(w).Name() != "lacnic" {
		t.Fatal("Invalid type: lacnic")
	}

	w = getNicProvider("51.2.25.4")
	t.Log(reflect.TypeOf(w).Name())

	if reflect.TypeOf(w).Name() != "ripe" {
		t.Fatal("Invalid type: ripe")
	}

	w = getNicProvider("154.125.148.148")
	t.Log(reflect.TypeOf(w).Name())

	if reflect.TypeOf(w).Name() != "afrinic" {
		t.Fatal("Invalid type: afrinic")
	}

	w = getNicProvider("184.12.95.8")
	t.Log(reflect.TypeOf(w).Name())

	if reflect.TypeOf(w).Name() != "arin" {
		t.Fatal("Invalid type: arin")
	}
}