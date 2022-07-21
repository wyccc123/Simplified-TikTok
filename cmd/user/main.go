package main

import (
	"net"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	"github.com/cloudwego/simplified-tiktok/cmd/user/dal"
	userdemo "github.com/cloudwego/simplified-tiktok/kitex_gen/userdemo/userservice"
	constants "github.com/cloudwego/simplified-tiktok/pkg/constants"
	"github.com/cloudwego/simplified-tiktok/pkg/middleware"
	tracer2 "github.com/cloudwego/simplified-tiktok/pkg/tracer"
	etcd "github.com/kitex-contrib/registry-etcd"
	trace "github.com/kitex-contrib/tracer-opentracing"
)

func Init() {
	tracer2.InitJaeger(constants.UserServiceName)
	dal.Init()
}

func main() {
	r, err := etcd.NewEtcdRegistry([]string{constants.EtcdAddress})
	if err != nil {
		panic(err)
	}
	addr, err := net.ResolveTCPAddr("tcp", "127.0.0.1:8889")
	if err != nil {
		panic(err)
	}
	Init()

	svr := userdemo.NewServer(new(UserServiceImpl),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: constants.UserServiceName}), // server name
		server.WithMiddleware(middleware.CommonMiddleware),                                             // middleware
		server.WithMiddleware(middleware.ServerMiddleware),
		server.WithServiceAddr(addr),                    // address
		server.WithMuxTransport(),                       // Multiplex
		server.WithSuite(trace.NewDefaultServerSuite()), // tracer
		server.WithRegistry(r),                          // registry
	)

	err = svr.Run()

	if err != nil {
		klog.Fatal(err)
	}
}
