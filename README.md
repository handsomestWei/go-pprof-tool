# go-pprof-tool
golang pprof性能分析监控工具使用

# Usages
提供两种方式
+ http接口访问
+ 采集到文件

## http接口访问
```
import (
	"github.com/handsomestWei/go-pprof-tool/pprof"
)

// 暴露端口提供http接口访问
pprof.RunAtAddrAndPort("addr:port")
```
采集信息页面
```
http://addr:port/debug/pprof/
```
下载采集文件
```
go tool pprof http://addr:port/debug/pprof/profile?seconds=30
```
## 采集到文件
```
import (
	"github.com/handsomestWei/go-pprof-tool/pprof"
)

// 指定采集文件输出目录，和采集持续时间
pprof.Dump2FileDuration("./cpu.prof", 10*time.Second)
```
打开采集文件
```
go tool pprof [filePath]
```
## 采集文件解析
```
go tool pprof [filePath] [command]
```
### top命令
关注sum%列：给定函数累积使用cpu总比例
### list命令
定位到可能有问题的代码段
```
list [function name]
例：list BlockDefault
```
### web命令
[需要安装工具](http://www.graphviz.org/download/)
生成调用链
### exit命令
退出