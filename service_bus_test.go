package service_bus

import (
	"github.com/nats-io/nats.go"
	"testing"
)

func TestMessage_Publish(t *testing.T) {
	nc, err := SingleNodeConnect("localhost:4222", "test-conn")
	if err != nil {
		t.Log("Error in MessagePublish Function")
		t.Error(err)
	}
	msg := Message{
		Connection:        nc,
		EncodedConnection: nil,
		Message:           nil,
		Type:              "test-message",
		Subject:           "test.package",
		Payload:           []byte("test-message"),
	}
	status, err := msg.Publish()
	if err != nil || status == false {
		t.Log("Error in MessagePublish Function")
		t.Error(err)
	}
	nc.Close()
}

func TestMessage_PublishJSON(t *testing.T) {
	ec, err := JSONEncodedSingleConnect("localhost:4222", "test-conn")
	if err != nil {
		t.Log("Error in MessagePublishJSON Function")
		t.Error(err)
	}
	payload := make(map[string]interface{})
	payload["test_string"] = "test"
	payload["test_integer"] = 1
	payload["test_bool"] = true
	payload["test_float"] = 1.25
	msg := Message{
		Connection:        nil,
		EncodedConnection: ec,
		Message:           nil,
		Type:              "test-message",
		Subject:           "test.package",
		Payload:           nil,
	}
	status, err := msg.PublishJSON()
	if err != nil || status == false {
		t.Log("Error in MessagePublishJSON Function")
		t.Error(err)
	}
	ec.Close()
}

func TestMessage_SynchronousSubscribe(t *testing.T) {
	nc, err := SingleNodeConnect("localhost:4222", "test-conn")
	if err != nil {
		t.Log("Error in SynchronousSubscribe Function")
		t.Error(err)
	}
	var r Subscriber
	r = Message{
		Connection:        nc,
		EncodedConnection: nil,
		Message:           nil,
		Type:              "test-message",
		Subject:           "test.package",
		Payload:           nil,
	}
	sub, err := r.SynchronousSubscribe()
	if err != nil {
		t.Log("Error in SynchronousSubscribe Function")
		t.Error(err)
	}
	if sub.Subject != "test.package" || !sub.IsValid() {
		t.Log("Error in SynchronousSubscribe Function")
		t.Error("Subject is wrong or invalid!")
	}
	_ = sub.Unsubscribe()
	nc.Close()
}

func TestMessage_AsynchronousSubscribe(t *testing.T) {
	nc, err := SingleNodeConnect("localhost:4222", "test-conn")
	if err != nil {
		t.Log("Error in AsynchronousSubscribe Function")
		t.Error(err)
	}
	var r Subscriber
	r = Message{
		Connection:        nc,
		EncodedConnection: nil,
		Message:           nil,
		Type:              "test-message",
		Subject:           "test.package",
		Payload:           nil,
	}
	var h nats.MsgHandler
	h = func(msg *nats.Msg) {
		println("Message: ", msg.Data)
	}
	sub, err := r.AsynchronousSubscribe(h)
	if err != nil {
		t.Log("Error in AsynchronousSubscribe Function")
		t.Error(err)
	}
	if sub.Subject != "test.package" || !sub.IsValid() {
		t.Log("Error in AsynchronousSubscribe Function")
		t.Error("Subject is wrong or invalid!")
	}
	_ = sub.Unsubscribe()
	nc.Close()
}

func TestMessage_JSONEncodedAsynchronousSubscribe(t *testing.T) {
	ec, err := JSONEncodedSingleConnect("localhost:4222", "test-conn")
	if err != nil {
		t.Log("Error in JSONEncodedAsynchronousSubscribe Function")
		t.Error(err)
	}
	var r Subscriber
	r = Message{
		Connection:        nil,
		EncodedConnection: ec,
		Message:           nil,
		Type:              "test-message",
		Subject:           "test.package",
		Payload:           nil,
	}
	var h nats.MsgHandler
	h = func(msg *nats.Msg) {
		println("Message: ", msg.Data)
	}
	sub, err := r.JSONEncodedAsynchronousSubscribe(h)
	if err != nil {
		t.Log("Error in JSONEncodedAsynchronousSubscribe Function")
		t.Error(err)
	}
	if sub.Subject != "test.package" || !sub.IsValid() {
		t.Log("Error in JSONEncodedAsynchronousSubscribe Function")
		t.Error("Subject is wrong or invalid!")
	}
	_ = sub.Unsubscribe()
	ec.Close()
}

func TestMessage_QueueSubscribe(t *testing.T) {
	nc, err := SingleNodeConnect("localhost:4222", "test-conn")
	if err != nil {
		t.Log("Error in QueueSubscribe Function")
		t.Error(err)
	}
	var r Subscriber
	r = Message{
		Connection:        nc,
		EncodedConnection: nil,
		Message:           nil,
		Type:              "test-message",
		Subject:           "test.package",
		Payload:           nil,
	}
	var h nats.MsgHandler
	h = func(msg *nats.Msg) {
		println("Message: ", msg.Data)
	}
	sub, err := r.QueueSubscribe(h, "test-queue")
	if err != nil {
		t.Log("Error in QueueSubscribe Function")
		t.Error(err)
	}
	if sub.Subject != "test.package" || !sub.IsValid() {
		t.Log("Error in QueueSubscribe Function")
		t.Error("Subject is wrong or invalid!")
	}
	_ = sub.Unsubscribe()
	nc.Close()
}
