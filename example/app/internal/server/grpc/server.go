package grpc

import (
	"fmt"
	"net"

	"github.com/UnderTreeTech/waterdrop/utils/xnet"

	"github.com/UnderTreeTech/waterdrop/pkg/conf"

	"google.golang.org/grpc"

	"github.com/UnderTreeTech/protobuf/demo"
	"github.com/UnderTreeTech/waterdrop/example/app/internal/service"
	"github.com/UnderTreeTech/waterdrop/pkg/registry"
	"github.com/UnderTreeTech/waterdrop/pkg/server/rpc"
)

type ServerInfo struct {
	Server      *rpc.Server
	ServiceInfo *registry.ServiceInfo
}

func New() *ServerInfo {
	config := &rpc.ServerConfig{}
	if err := conf.Unmarshal("Server.RPC", config); err != nil {
		panic(fmt.Sprintf("unmarshal grpc server config fail, err msg %s", err.Error()))
	}
	fmt.Println("rpc config", config)
	server := rpc.NewServer(config)
	registerServers(server.Server(), &service.Service{})

	//server.Use()

	addr := server.Start()
	_, port, _ := net.SplitHostPort(addr.String())
	serviceInfo := &registry.ServiceInfo{
		Name:    "service.user.v1",
		Scheme:  "grpc",
		Addr:    fmt.Sprintf("%s://%s:%s", "grpc", xnet.InternalIP(), port),
		Version: "1.0.0",
	}

	return &ServerInfo{Server: server, ServiceInfo: serviceInfo}
}

func registerServers(g *grpc.Server, s *service.Service) {
	demo.RegisterDemoServer(g, s)
}
