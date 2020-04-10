package service_bus

import "testing"

func TestSingleNodeConnect(t *testing.T) {
	nc, err := SingleNodeConnect("localhost:4222", "test-conn")
	if err != nil {
		t.Log("Error Occurred at SingleNodeConnect Function")
		t.Error(err)
	}
	if !nc.IsConnected() {
		t.Error("Not connected to server. SingleNodeConnect Function")
	}
}
