package master

import (
	"net/url"
	"testing"

	"github.com/ayush18023/Load_balancer_Fyp/internal/sqlite"
)

func TestAddRecord(t *testing.T) {
	m := &Master{}
	var err error
	m.dbDns, err = sqlite.CreateConn()
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
	err = m.AddDnsRecord(task)
	if err != nil {
		t.Fatalf("didnt work")
	}
}
