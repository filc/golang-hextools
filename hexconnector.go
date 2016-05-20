package hextools

import (
// "fmt"
// "reflect"
)

type Connector interface {
	RegisterAdapter(Adapter, string)
	GetAdapter(string) Adapter
	Set(string, interface{})
	Get(string) interface{}
}

type HexConnector struct {
	adapters map[string]*AdapterStorageItem
	storage  map[string]interface{}
}

func (cn *HexConnector) RegisterAdapter(adapter Adapter, portname string) {
	if cn.adapters == nil {
		cn.adapters = make(map[string]*AdapterStorageItem)
	}

	adapter.SetConnector(cn)
	cn.adapters[portname] = &AdapterStorageItem{adapter, false}
}

func (cn *HexConnector) GetAdapter(portname string) Adapter {
	adapterItem := cn.adapters[portname]

	if adapterItem == nil {
		return nil
	}

	if !adapterItem.Initiated() {
		adapterItem.Init()
	}

	return adapterItem.GetAdapter()
}

func (cn *HexConnector) Get(key string) interface{} {
	if cn.storage == nil {
		return nil
	}

	return cn.storage[key]
}

func (cn *HexConnector) Set(key string, value interface{}) {
	if cn.storage == nil {
		cn.storage = make(map[string]interface{})
	}

	cn.storage[key] = value
}

var _ Connector = (*HexConnector)(nil)
