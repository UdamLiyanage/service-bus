package service_bus

import (
	"github.com/nats-io/nats.go"
	"strings"
)

func SingleNodeConnect(url string) (*nats.Conn, error) {
	nc, err := nats.Connect(url)
	if err != nil {
		return nil, err
	}
	return nc, nil
}

func ClusterConnect(urls []string) (*nats.Conn, error) {
	nc, err := nats.Connect(strings.Join(urls, ","))
	if err != nil {
		return nil, err
	}
	return nc, nil
}
