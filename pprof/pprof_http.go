package pprof

import (
	// 引入pprof
	_ "net/http/pprof"
	"net/http"
)

// 暴露端口提供http接口访问
func RunAtAddrAndPort(addrPort string) {
	// 会阻塞当前线程，因此放到go routinue里
	go func() {
		// 访问http://addr:port/debug/pprof/
		// go tool pprof http://addr:port/debug/pprof/profile?seconds=20
		err := http.ListenAndServe(addrPort, nil)
		if err != nil {
			panic(err)
		}
	}()
}
