package service_bus

import (
	"github.com/nats-io/nats.go"
	"strings"
	"time"
)

func SingleNodeConnect(url, name string) (*nats.Conn, error) {
	nc, err := nats.Connect(
		url,
		nats.Name(name),
		nats.Timeout(10*time.Second),
		nats.PingInterval(20*time.Second),
		nats.MaxPingsOutstanding(5),
		nats.MaxReconnects(10),
		nats.ReconnectWait(10*time.Second),
		nats.ReconnectBufSize(5*1024*1024))
	if err != nil {
		return nil, err
	}
	return nc, nil
}

func ClusterConnect(urls []string, name string) (*nats.Conn, error) {
	nc, err := nats.Connect(
		strings.Join(urls, ","),
		nats.Name(name),
		nats.Timeout(10*time.Second),
		nats.PingInterval(20*time.Second),
		nats.MaxPingsOutstanding(5),
		nats.MaxReconnects(10),
		nats.ReconnectWait(10*time.Second),
		nats.ReconnectBufSize(5*1024*1024))
	if err != nil {
		return nil, err
	}
	return nc, nil
}

func JSONEncodedSingleConnect(url, name string) (*nats.EncodedConn, error) {
	nc, err := nats.Connect(
		url,
		nats.Name(name),
		nats.Timeout(10*time.Second),
		nats.PingInterval(20*time.Second),
		nats.MaxPingsOutstanding(5),
		nats.MaxReconnects(10),
		nats.ReconnectWait(10*time.Second),
		nats.ReconnectBufSize(5*1024*1024))
	if err != nil {
		return nil, err
	}
	ec, err := nats.NewEncodedConn(nc, nats.JSON_ENCODER)
	if err != nil {
		return nil, err
	}
	return ec, nil
}

func JSONEncodedClusterConnect(urls []string, name string) (*nats.EncodedConn, error) {
	nc, err := nats.Connect(
		strings.Join(urls, ","),
		nats.Name(name),
		nats.Timeout(10*time.Second),
		nats.PingInterval(20*time.Second),
		nats.MaxPingsOutstanding(5),
		nats.MaxReconnects(10),
		nats.ReconnectWait(10*time.Second),
		nats.ReconnectBufSize(5*1024*1024))
	if err != nil {
		return nil, err
	}
	ec, err := nats.NewEncodedConn(nc, nats.JSON_ENCODER)
	if err != nil {
		return nil, err
	}
	return ec, nil
}
