package main

import (
	"net"

	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	"github.com/cloudwego/simplified-tiktok/cmd/comment/dal"
	commentdemo "github.com/cloudwego/simplified-tiktok/kitex_gen/commentdemo/commentservice"
	constants "github.com/cloudwego/simplified-tiktok/pkg/constants"
	"github.com/cloudwego/simplified-tiktok/pkg/middleware"
	tracer2 "github.com/cloudwego/simplified-tiktok/pkg/tracer"
	etcd "github.com/kitex-contrib/registry-etcd"
	trace "github.com/kitex-contrib/tracer-opentracing"
	"github.com/u2takey/go-utils/klog"
)

func Init() {
	tracer2.InitJaeger(constants.VideoServiceName)
	dal.Init()
}

func main() {
	r, err := etcd.NewEtcdRegistry([]string{constants.EtcdAddress}) // r should not be reused.
	if err != nil {
		panic(err)
	}
	addr, err := net.ResolveTCPAddr("tcp", "127.0.0.1:8888")
	if err != nil {
		panic(err)
	}
	Init()

	svr := commentdemo.NewServer(new(CommentServiceImpl),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: constants.CommentServiceName}), // server name
		server.WithMiddleware(middleware.CommonMiddleware),                                                // middleWare
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
