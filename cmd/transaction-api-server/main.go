package main

import (
	"net"
	"transaction-api-server/api"

	"google.golang.org/grpc"

	"github.com/sirupsen/logrus"
)

func main() {
	logrus.Info("Iniciando container ...")
	container := api.BuildContainer()

	err := container.Invoke(func(s *api.Server) {
		listener, err := net.Listen("tcp", ":50051")
		if err != nil {
			logrus.Error("Falha ao iniciar o servidor: " + err.Error())
		}

		grpcServer := grpc.NewServer()
		s.Register(grpcServer)

		logrus.Info("Servidor gRPC ouvindo na porta :50051...")
		if err := grpcServer.Serve(listener); err != nil {
			logrus.Error("Falha ao iniciar o servidor: " + err.Error())
		}
	})

	if err != nil {
		logrus.Error("Falha ao iniciar o servidor: " + err.Error())
	}
}
