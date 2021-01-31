package pprof

import (
	"os"
	"runtime/pprof"
	"time"
)

// 指定采集文件输出目录，和采集持续时间
func Dump2FileDuration(filePath string, d time.Duration) {
	// go tool pprof filePath
	go func() {
		Dump2FileStart(filePath)
		defer Dump2FileEnd()
		time.Sleep(d)
	}()
}

func Dump2FileStart(filePath string) {
	go func() {
		file, err := os.Create(filePath)
		if err != nil {
			panic(err)
		}
		// 获取运行时cpu信息
		err = pprof.StartCPUProfile(file)
		if err != nil {
			panic(err)
		}
	}()
}

func Dump2FileEnd() {
	pprof.StopCPUProfile()
}