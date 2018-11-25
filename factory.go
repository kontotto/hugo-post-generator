package main

type Factory interface {
	CreateProvider() (Provider, error)
}

func NewFactory(meta *Meta) (Factory, error) {
	return &niconicoFactory{
		Meta: meta,
	}, nil
}

type niconicoFactory struct {
	Meta *Meta
}

func (f *niconicoFactory) CreateProvider() (Provider, error) {
	return nil, nil
}
