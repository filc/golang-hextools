package hextools

import (
	// "fmt"
	"reflect"
	"testing"
)

type FakeAdapter struct {
	HexAdapter
}

func TestCreateHexConnector(t *testing.T) {
	cn := new(HexConnector)

	if len(cn.adapters) != 0 {
		t.Errorf("HexConnector adapters are not empty")
	}
}

func TestImplementsTheInterfaceMethods(t *testing.T) {
	cn := new(HexConnector)

	if cn == nil {
		t.Error("Error in creating the connector")
	}
}

func TestGetFromEmptyStorage(t *testing.T) {
	cn := new(HexConnector)

	if cn.Get("test_key") != nil {
		t.Error("I got something instead of nil")
	}
}

func TestGetUnexistingValue(t *testing.T) {
	cn := new(HexConnector)
	cn.Set("test_key", 1)

	if cn.Get("test_key_2") != nil {
		t.Error("I got something instead of nil")
	}
}

func TestSetGet(t *testing.T) {
	cn := new(HexConnector)

	var testValue = map[string]int{"first": 1, "second": 2}
	cn.Set("test_key", testValue)

	if !reflect.DeepEqual(cn.Get("test_key"), testValue) {
		t.Error("Something wrong with setter/getter")
	}
}

func TestGetNonExistingAdapter(t *testing.T) {
	cn := new(HexConnector)
	if cn.GetAdapter("nonexisting") != nil {
		t.Error("Get a non existing adapter should return nil")
	}
}

func TestRegisterAndGetAdapter(t *testing.T) {
	cn := new(HexConnector)
	adapter := new(FakeAdapter)

	cn.RegisterAdapter(adapter, "testport")

	if cn.GetAdapter("testport") != adapter {
		t.Error("It's not the adapter that was given")
	}
}
