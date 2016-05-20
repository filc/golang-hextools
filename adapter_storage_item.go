package hextools

type AdapterStorageItem struct {
	adapter   Adapter
	initiated bool
}

func (asi *AdapterStorageItem) Init() {
	asi.adapter.InitAdapter()
	asi.initiated = true
}

func (asi *AdapterStorageItem) Initiated() bool {
	return asi.initiated
}

func (asi *AdapterStorageItem) GetAdapter() Adapter {
	return asi.adapter
}
