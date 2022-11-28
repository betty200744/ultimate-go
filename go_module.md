### [GO ENVIRON](https://pkg.go.dev/cmd/go#hdr-GOPATH_environment_variable) 
* https://pkg.go.dev/cmd/go#hdr-GOPATH_environment_variable
* GO ENVIRON
    * GOROOT, go install dir
        * "/usr/local/go", go sdk包目录，含toolchain ,build package...
    * GOBIN , go binary installed dir
        * `/usr/local/go/bin` , go 这个bin文件存放目录
    * GOPATH, used to resolve import statements, list places to look for go code
        * `/Users/betty/go`， default $HOME/go, 
        * GOPATH/bin,
        * GOPATH/src,
        * GOPATH/pkg, 
    * GO111MODULE
        * go 1.11
            * GOPATH, GO MODULE两种模式共存
            * GOPATH/src目录，默认off, 非GOPATH/src目录，又有go.mod,默认on
        * go 1.12
            * GOPATH, GO MODULE两种模式共存，
            * GOPATH/src目录，默认off, 非GOPATH/src目录，默认on
        * go 1.13
            * GO MODULE模式
    * GOCACHE
        * `/Users/betty/Library/Caches/go-build`
        * go build, test等的缓存，加快构建速度
    * GOPROXY
        * "https://goproxy.io,direct"， 设置代理
    * GOTOOLDIR , toolchain
        * /usr/local/go/pkg/tool/darwin_amd64
        * cgo, compile, doc, fix, link, pprof,trace, vet
    * GOSUMDB
        *  Go module checksum database, 校验代码库，避免下载的库出现不安全
### [Go Modules模式](https://golang.org/doc/code.html)
* https://blog.golang.org/using-go-modules
* module, 一堆package, 含go.mod, 含module gobyexample即module path
* submodule, module下可有子module，即子目录有go.mod
* package, 一个文件夹一个package, 同package都可见，不可重复定义
* import path, 即go.mod 的module path, 所有package的前缀，即路径为module path + package path
* go build, 默认当前目录下
* go install， 默认安装到$GOROOT/bin目录下
* $GOPATH/pkg/mod下存放所有go项目的不同版本的library, 因为不同项目不同版本
 
### [GOPATH模式](https://golang.org/doc/gopath_code.html) 
* 1.11及之前版本都用GOPATH
* 所有包都在GOPATH里面，无法对版本进行管理
* GO vendor, 解决版本不一致问题
* 可直接的版本，也可GOPATH的版本
* goverdor下载当前版本的依赖到git目录下，老是需要更新verdor下的文件，麻烦
* 不同项目用的依赖版本不一样导致强类型老报错
* 每个项目都再自己项目下有本地的govendor，vendor.json， 也很麻烦 
* git管理很不友好
* 特别是后来大项目， vendor同步太大，不同步大家又不一样
* GOPATH, 所有go代码都需要放到一个workspace下面
* $GOPATH/src, 所有代码都放到$GOPATH/src目录下
* $GOPATH/bin, 所有可执行文件， 也是go install的目录
* go build, 默认当前目录下
* go install， 默认安装到$GOPATH/bin目录下 , and $GOBIN/bin目录下

##  [Go Tools, Command Documentation](https://golang.org/cmd/go/)
* go help
    * go help 
* go bug, 
    * report a bug to go 
* go run 
    * Compile and run Go program
    * go run ./main.go
* go build
    * Compile packages and dependencies
    * 编译， 交叉编译用
    * `go build -o migrate_app -i ./ `
* go install
    * Compile and install packages and dependencies
    * 默认安装目录$GOROOT
* go test, 
    * `go test  --cover -v class_test.go`
* go clean, 
    * Remove object files and cached files
    * `go clean --testcache`
* go doc
    * Show documentation for package or symbol, 查看文档用
    * `go doc github.com/jinzhu/gorm.open`
* go env
    * Print Go environment information， 查看go环境变了
    * 很重要，获取当前go环境， 注意交叉编译时需要设置GO ENV
    * `go env`
* go fix
    * 一般不用，即查找旧api自动替换为新api
    * Update packages to use new APIs
* go generate
    * Generate Go files by processing source
    * `很重要， 增加接口时可以自动mock`
    * 命令行下运行： 
        * go generate
    * 文件中运行： 
        * `//go:generate mockgen -source=cloudconfig.go -destination=../mock/cloudconfig_mock.go`
        * https://github.com/golang/mock
* go get 
    * go get github.com/bazelbuild/bazel-gazelle/cmd/gazelle
    * 新增依赖， Add dependencies to current module and install them
    * `安装包，相当于yarn install, npm install`
* go list
    * List packages or modules
* go mod
    * Module maintenance
    * go mod tidy

### go build
* go build , 编译包及依赖，compile packages and dependencies
* go build [-o output] [-i] [build flags] [packages]
* CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o app cron/es-full-sync-shop/main.go
* CGO_ENABLED=1 GOOS=linux go build -tags netgo -mod=vendor -installsuffix cgo -o /go/src/whgo/iot/${DIR}/app /go/src/whgo/iot/${DIR}/main.go
* CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -a -installsuffix cgo -o knight main.go
* go build -o mango-store--srv main.go plugin.go
* CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o acu-linux cli.go
* 

### go run 
go run *.go
├── Main package is executed
├── All imported packages are initialized
|  ├── All imported packages are initialized (recursive definition)
|  ├── All global variables are initialized
|  └── init functions are called in lexical file name order
└── Main package is initialized
   ├── All global variables are initialized
   └── init functions are called in lexical file name order
*/


