package pprof

import (
	"testing"
	"net/http"
	"github.com/handsomestWei/go-pprof-tool/test"
)

func TestListen(t *testing.T) {

	ipAddr := "172.16.13.59"
	RunAtAddrAndPort(ipAddr + ":8887")

	go test.BlockDefault()
	http.ListenAndServe(ipAddr + ":8888", nil)
}