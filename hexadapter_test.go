package hextools

import (
	// "fmt"
	// "reflect"
	"testing"
)

type Asd interface {
	GetNewVar() int
}

type NewHexAdapter struct {
	HexAdapter
	newVar int
}

func (ha *NewHexAdapter) InitAdapter() (bool, error) {
	ha.newVar = 3
	return true, nil
}

func (ha *NewHexAdapter) GetNewVar() int {
	return ha.newVar
}

type DoesNotOverwriteHexAdapter struct {
	HexAdapter
}

func getAdapter(adapter Adapter) (hc *HexConnector, ha Adapter) {
	hc = new(HexConnector)
	ha = adapter
	ha.SetConnector(hc)
	return
}

func TestHexAdapter(t *testing.T) {
	hc, ha := getAdapter(new(HexAdapter))

	if hc != ha.GetConnector() {
		t.Errorf("HexConnector is not a reference")
	}
}

func TestInitAdapterMethodDoesNotHaveToBeOverwritten(t *testing.T) {
	_, ha := getAdapter(new(DoesNotOverwriteHexAdapter))

	func(ha Adapter) {}(ha)
}

func TestInitAdapterMethodMustBeOverwrittenCorrectly(t *testing.T) {
	_, ha := getAdapter(new(NewHexAdapter))

	func(ha Adapter) {}(ha)
}

func TestInitAdapterMethodCanBeOverwritten(t *testing.T) {
	_, ha := getAdapter(new(NewHexAdapter))

	if ha.(*NewHexAdapter).GetNewVar() != 0 {
		t.Error("Default value of newVar is wrong")
	}

	if _, err := ha.InitAdapter(); err != nil {
		t.Errorf("An error happened during initAdapter: %s", err)
	}

	if ha.(*NewHexAdapter).GetNewVar() != 3 {
		t.Error("initAdapter method cannot be overwritten")
	}
}
