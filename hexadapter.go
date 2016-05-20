package hextools

type Adapter interface {
	InitAdapter() (bool, error)
	SetConnector(*HexConnector)
	GetConnector() *HexConnector
}

type HexAdapter struct {
	cn *HexConnector
}

func (ha *HexAdapter) InitAdapter() (bool, error) {
	return true, nil
}

func (ha *HexAdapter) SetConnector(cn *HexConnector) {
	ha.cn = cn
}

func (ha *HexAdapter) GetConnector() *HexConnector {
	return ha.cn
}

var _ Adapter = (*HexAdapter)(nil)
