package api

import (
	"context"
	"transaction-api-server/internal/domain/entity"
	pb "transaction-api-server/proto/pb"

	"google.golang.org/protobuf/types/known/timestamppb"
)

// Listar contas
func (s *Server) ListAccounts(ctx context.Context, req *pb.AccountEmpty) (*pb.AccountListResponse, error) {
	accountsData, err := s.AccountUseCase.ListAccounts()
	if err != nil {
		return nil, err
	}

	var accounts []*pb.AccountResponse
	for _, acc := range accountsData {
		accounts = append(accounts, &pb.AccountResponse{
			AccountId: acc.ID,
			Name:      acc.Name,
			Document:  acc.Document,
			CreatedAt: timestamppb.New(acc.CreatedAt),
			UpdatedAt: timestamppb.New(acc.UpdatedAt),
		})
	}

	return &pb.AccountListResponse{Accounts: accounts}, nil
}

// Obter conta por ID
func (s *Server) GetAccountById(ctx context.Context, req *pb.AccountId) (*pb.AccountResponse, error) {
	accountsData, err := s.AccountUseCase.GetAccountById(req.AccountId)
	if err != nil {
		return nil, err
	}

	return &pb.AccountResponse{
		AccountId: accountsData.ID,
		Name:      accountsData.Name,
		Document:  accountsData.Document,
		CreatedAt: timestamppb.New(accountsData.CreatedAt),
		UpdatedAt: timestamppb.New(accountsData.UpdatedAt),
	}, nil
}

// Criar conta
func (s *Server) CreateAccount(ctx context.Context, req *pb.AccountRequest) (*pb.AccountIdResponse, error) {
	account := entity.Account{
		Name:     req.Name,
		Document: req.Document,
	}

	accountResponse, err := s.AccountUseCase.CreateAccount(account)
	if err != nil {
		return nil, err
	}

	return &pb.AccountIdResponse{AccountId: accountResponse.ID}, nil
}
