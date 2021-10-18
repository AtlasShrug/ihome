package main

import (
	"getCaptcha/handler"
	pb "getCaptcha/proto"

	"github.com/asim/go-micro/plugins/registry/consul/v4"
	"go-micro.dev/v4"
	log "go-micro.dev/v4/logger"
)

var (
	service = "getcaptcha"
	version = "latest"
)

func main() {
	reg := consul.NewRegistry()
	// Create service
	srv := micro.NewService(
		micro.Name(service),
		micro.Version(version),
		micro.Registry(reg),
		micro.Address(":12341"),
	)
	srv.Init()

	// Register handler
	pb.RegisterGetCaptchaHandler(srv.Server(), new(handler.GetCaptcha))

	// Run service
	if err := srv.Run(); err != nil {
		log.Fatal(err)
	}
}
