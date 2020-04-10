package service_bus

import (
	"github.com/nats-io/nats.go"
)

type (
	Publisher interface {
		Publish() (bool, error)
		PublishJSON() (bool, error)
	}

	Subscriber interface {
		SynchronousSubscribe() (*nats.Subscription, error)
		AsynchronousSubscribe(handler nats.MsgHandler) (*nats.Subscription, error)
		QueueSubscribe(handler nats.MsgHandler, queueName string) (*nats.Subscription, error)
		JSONEncodedQueueSubscribe(handler nats.MsgHandler, queueName string) (*nats.Subscription, error)
		JSONEncodedAsynchronousSubscribe(handler nats.MsgHandler) (*nats.Subscription, error)
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

func (m Message) SynchronousSubscribe() (*nats.Subscription, error) {
	sub, err := m.Connection.SubscribeSync(m.Subject)
	if err != nil {
		return nil, err
	}
	return sub, err
}

func (m Message) AsynchronousSubscribe(handler nats.MsgHandler) (*nats.Subscription, error) {
	sub, err := m.Connection.Subscribe(m.Subject, handler)
	if err != nil {
		return nil, err
	}
	return sub, nil
}

func (m Message) JSONEncodedAsynchronousSubscribe(handler nats.MsgHandler) (*nats.Subscription, error) {
	sub, err := m.EncodedConnection.Subscribe(m.Subject, handler)
	if err != nil {
		return nil, err
	}
	return sub, nil
}

func (m Message) QueueSubscribe(handler nats.MsgHandler, queueName string) (*nats.Subscription, error) {
	sub, err := m.Connection.QueueSubscribe(m.Subject, queueName, handler)
	if err != nil {
		return nil, err
	}
	return sub, nil
}

func (m Message) JSONEncodedQueueSubscribe(handler nats.MsgHandler, queueName string) (*nats.Subscription, error) {
	sub, err := m.EncodedConnection.QueueSubscribe(m.Subject, queueName, handler)
	if err != nil {
		return nil, err
	}
	return sub, nil
}
