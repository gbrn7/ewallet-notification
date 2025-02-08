package cmd

import (
	"ewallet-notification/cmd/proto/notification"
	"ewallet-notification/helpers"
	"ewallet-notification/internal/api"
	"ewallet-notification/internal/repository"
	"ewallet-notification/internal/services"
	"log"
	"net"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func ServeGRPC() {
	// init dependency
	d := dependencyInject()

	s := grpc.NewServer()
	notification.RegisterNotificationServiceServer(s, d.EmailAPI)
	// list method

	lis, err := net.Listen("tcp", ":"+helpers.GetEnv("GRPC_PORT", "7003"))
	if err != nil {
		log.Fatal("failed to listen grpc port: ", err)
	}

	logrus.Info("start listening grpc on port:" + helpers.GetEnv("GRPC_PORT", "7003"))
	if err := s.Serve(lis); err != nil {
		log.Fatal("failed to serve grpc port: ", err)
	}

}

type Dependency struct {
	EmailAPI *api.EmailAPI
}

func dependencyInject() *Dependency {
	emailRepo := &repository.EmailRepo{
		DB: helpers.DB,
	}

	emailSvc := &services.EmailService{
		EmailRepo: emailRepo,
	}

	emailAPI := &api.EmailAPI{
		EmailService: emailSvc,
	}

	return &Dependency{
		EmailAPI: emailAPI,
	}
}
