package api

import (
	"database/sql"
	"transaction-api-server/config"
	pb "transaction-api-server/proto/pb"
	"transaction-api-server/repository"
	"transaction-api-server/repository/postgres"
	"transaction-api-server/service"

	"go.uber.org/dig"
	"google.golang.org/grpc"
)

type Server struct {
	pb.UnimplementedAccountServiceServer
	pb.UnimplementedTransactionServiceServer
	AccountUseCase     service.AccountUseCase
	TransactionUseCase service.TransactionUseCase
}

func BuildContainer() *dig.Container {
	container := dig.New()

	// Banco de dados - Adicionar a conexão com o banco
	container.Provide(config.ConnectDB)

	// Dependências
	// Repositórios
	container.Provide(func(db *sql.DB) repository.AccountRepository {
		return postgres.NewAccountRepository(db)
	})
	container.Provide(func(db *sql.DB) repository.TransactionRepository {
		return postgres.NewTransactionRepository(db)
	})

	container.Provide(service.NewAccountUseCase)
	container.Provide(service.NewTransactionUseCase)

	// Criação do servidor
	container.Provide(func(accountUseCase service.AccountUseCase, transactionUseCase service.TransactionUseCase) *Server {
		return &Server{
			AccountUseCase:     accountUseCase,
			TransactionUseCase: transactionUseCase,
		}
	})

	return container
}

// Função para registrar serviços no gRPC
func (s *Server) Register(grpcServer *grpc.Server) {
	pb.RegisterAccountServiceServer(grpcServer, s)
	pb.RegisterTransactionServiceServer(grpcServer, s)
}
