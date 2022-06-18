package config

type Listener interface {
	Address() string
}

type listener struct {
	address string
}

func NewListener(address string) Listener {
	return &listener{
		address: address,
	}
}

func (l *listener) Address() string {
	return l.address
}
