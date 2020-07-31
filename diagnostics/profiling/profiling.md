

### Processor and Thread and goroutine
* P - processor, a resource that is required to execute Go code. 一般一个进程
* M - worker thread, or machine. 根据cpu情况创建， 一般小于10个线程， 请求越多Thread越多， goroutine越多Thread也越多
* G - goroutine.
* Each goroutine (G) runs on an OS thread (M) that is assigned to a logical CPU (P)
* 



### flow
* cd pprof/http  or cd pprof/gin
* go build -o main // 编译程序
* ./main & // 运行程序
* go tool pprof main // 预设查看信息
    * http://127.0.0.1:6060/debug/pprof/ //查看全部支持的profile
    * go tool pprof main http://127.0.0.1:6060/debug/pprof/profile // 预设查看cpu profile
    * go tool pprof main http://127.0.0.1:6060/debug/pprof/heap // 预设查看heap profile
    * go tool pprof main http://127.0.0.1:6060/debug/pprof/threadcreate // 预设查看thread
    * go tool pprof main http://127.0.0.1:6060/debug/pprof/goroutine //预设查看goroutines信息
    * wget http://localhost:6060/debug/pprof/trace\?seconds\=10 -O ./trace // 预设查看trace profile
* ab 测试 // 请求程序
    * ab -k -c 8 -n 100000 "http://127.0.0.1:6060/concat/?str=test&count=50" &
    * ab -k -c 8 -n 100000 "http://127.0.0.1:6060/fib/?n=50&type=recursive" &
    * ab -k -c 8 -n 100000 "http://127.0.0.1:6060/fib/?n=50&type=iterative" &
* 用pprof visualization tool 可视化数据 
    * go tool pprof -http=:8080 ./pprof.pprof.samples.cpu.001.pb.gz
    * go tool trace ./trace

### profiling data
* [net/http/pprof收集](https://golang.org/pkg/net/http/pprof/)
* [go test收集]()


### Predefined profiles by runtime/pprof package
* cpu: 查看cpu profile
    * `go tool pprof main http://127.0.0.1:6060/debug/pprof/profile`
    * CPU profile determines where a program spends its time while actively consuming CPU cycles 
* heap: 
    * `go tool pprof main http://127.0.0.1:6060/debug/pprof/heap`
    * Heap profile reports memory allocation samples; used to monitor current and historical memory usage
* threadcreate: 
    * Thread creation profile reports the sections of the program that lead the creation of new OS threads.
* goroutine: 
    * `wget http://localhost:6060/debug/pprof/trace?seconds=5`
    * Goroutine profile reports the stack traces of all current goroutines.
* block: 
    * Block profile shows where goroutines block waiting on synchronization primitives. 
    * Block profile is not enabled by default; use runtime.SetBlockProfileRate to enable it.


### pprof visualization tool
* graphviz: brew install graphviz
* `go tool pprof -http=:8080 ./pprof.pprof.samples.cpu.001.pb.gz`

### gin pprof middleware
* go get github.com/gin-contrib/pprof




