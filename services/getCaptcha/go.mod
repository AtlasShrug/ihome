module getCaptcha

go 1.16

require (
	github.com/afocus/captcha v0.0.0-20191010092841-4bd1f21c8868
	github.com/asim/go-micro/plugins/registry/consul/v4 v4.0.0-20211016045651-29fefbad4e06
	github.com/golang/freetype v0.0.0-20170609003504-e2365dfdc4a0 // indirect
	github.com/gomodule/redigo v1.8.5
	go-micro.dev/v4 v4.1.0
	google.golang.org/protobuf v1.26.0
)

// This can be removed once etcd becomes go gettable, version 3.4 and 3.5 is not,
// see https://github.com/etcd-io/etcd/issues/11154 and https://github.com/etcd-io/etcd/issues/11931.
replace google.golang.org/grpc => google.golang.org/grpc v1.26.0
