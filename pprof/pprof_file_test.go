package pprof

import (
	"testing"
	"time"
	"github.com/handsomestWei/go-pprof-tool/test"
)

func TestDump2FileDuration_1(t *testing.T) {

	// go tool pprof filePath
	filePath := "F:\\cpu.prof"
	Dump2FileDuration(filePath, 10 *time.Second)

	go test.BlockDefault()
	time.Sleep(12 *time.Second)
}

func TestDump2FileDuration_2(t *testing.T) {

	// go tool pprof filePath
	filePath := "F:\\cpu.prof"
	Dump2FileDuration(filePath, 10 *time.Second)

	test.BlockNoneDefault()
	time.Sleep(12 *time.Second)
}

