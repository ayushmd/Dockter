package master

import (
	"net/url"
	"testing"
)

func TestAddRecord(t *testing.T) {
	m := &Master{}
	task := Task{
		Subdomain: "test",
		URL: url.URL{
			Host: "127.0.0.1",
		},
		Hostport:    "3000",
		Runningport: "8000",
		ImageName:   "tyranthex/fyp_deps:node0",
		ContainerID: "acjahcuahcoash",
	}
	err := m.AddDnsRecord(task)
	if err != nil {
		t.Fatalf("didnt work")
	}
}
