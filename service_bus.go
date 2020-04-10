package service_bus

import (
	"github.com/nats-io/nats.go"
)

type (
	Publisher interface {
		Publish() (bool, error)
		PublishJSON() (bool, error)
	}

	Receiver interface {
		SynchronousSubscribe() (bool, error)
		AsynchronousSubscribe() (bool, error)
		AutoUnsubscribe() (bool, error)
		QueueSubscribe() (bool, error)
		Unsubscribe() (bool, error)
	}

	Message struct {
		Connection        *nats.Conn
		EncodedConnection *nats.EncodedConn
		Message           map[string]interface{}
		Type              string
		Subject           string
		Payload           []byte
	}
)

func (m Message) Publish() (bool, error) {
	if err := m.Connection.Publish(m.Subject, m.Payload); err != nil {
		return false, err
	}
	return true, nil
}

func (m Message) PublishJSON() (bool, error) {
	if err := m.EncodedConnection.Publish(m.Subject, &m.Message); err != nil {
		return false, err
	}
	return true, nil
}
