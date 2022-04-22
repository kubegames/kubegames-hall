package nats

import (
	"github.com/nats-io/nats.go"
)

type (
	NatsCfg struct {
		Nats string `ini:"Nats"`
	}

	Nats interface {
		Engine() *nats.Conn
	}

	natsImpl struct {
		conn *nats.Conn
	}
)

func NewNats(cfg *NatsCfg) Nats {
	conn, err := nats.Connect(cfg.Nats)
	if err != nil {
		panic(err)
	}
	return &natsImpl{
		conn: conn,
	}
}

func (impl *natsImpl) Engine() *nats.Conn {
	return impl.conn
}
