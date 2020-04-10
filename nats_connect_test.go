package service_bus

import (
	"testing"
)

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

func TestClusterConnect(t *testing.T) {
	urls := make([]string, 2)
	urls = append(urls, "localhost:4222")
	urls = append(urls, "localhost:5222")
	nc, err := ClusterConnect(urls, "test-conn")
	if err != nil {
		t.Log("Error Occurred at SingleNodeConnect Function")
		t.Error(err)
	}
	if !nc.IsConnected() {
		t.Error("Not connected to server. SingleNodeConnect Function")
	}
}
