package service_bus

import "testing"

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
}
