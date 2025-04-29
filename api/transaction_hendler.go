package api

import (
	"context"
	"time"
	"transaction-api-server/internal/domain/entity"
	pb "transaction-api-server/proto/pb"

	"google.golang.org/protobuf/types/known/timestamppb"
)

// Listar transactions
func (s *Server) ListTransactions(ctx context.Context, req *pb.TransactionEmpty) (*pb.TransactionListResponse, error) {
	transactionsData, err := s.TransactionUseCase.ListTransactions()
	if err != nil {
		return nil, err
	}

	var transactions []*pb.TransactionResponse
	for _, tsc := range transactionsData {
		transactions = append(transactions, &pb.TransactionResponse{
			TransactionId:     tsc.ID,
			AccountId:         tsc.AccountId,
			TransactionTypeId: tsc.TransactionTypeId,
			Amount:            tsc.Amount,
			TransactionDate:   timestamppb.New(tsc.TransactionDate),
		})
	}
	return &pb.TransactionListResponse{Transactions: transactions}, nil
}

// Criar transaction
func (s *Server) CreateTransaction(ctx context.Context, req *pb.TransactionRequest) (*pb.TransactionIdResponse, error) {
	transaction := entity.Transaction{
		AccountId:         req.AccountId,
		TransactionTypeId: req.TransactionTypeId,
		Amount:            req.Amount,
		TransactionDate:   time.Now(),
	}

	transactionResponse, err := s.TransactionUseCase.CreateTransaction(transaction)
	if err != nil {
		return nil, err
	}

	return &pb.TransactionIdResponse{TransactionId: transactionResponse.ID}, nil
}

// Listar montante transacoes por account ID
func (s *Server) ListTransactionsAmount(ctx context.Context, req *pb.Document) (*pb.TransactionAmountListResponse, error) {
	transactionsData, err := s.TransactionUseCase.ListTransactionsAmount(req.Document)
	if err != nil {
		return nil, err
	}

	var transactions []*pb.TransactionAmountResponse
	for _, tsc := range transactionsData {
		transactions = append(transactions, &pb.TransactionAmountResponse{
			Name:     tsc.Name,
			Document: tsc.Document,
			Amount:   tsc.Amount,
		})
	}
	return &pb.TransactionAmountListResponse{TransactionsAmount: transactions}, nil
}

// Listar transactionAcount por account ID
func (s *Server) ListTransactionsByAccountId(ctx context.Context, req *pb.AccountId) (*pb.TransactionJoinAccountListResponse, error) {
	transactionsData, err := s.TransactionUseCase.ListTransactionsByAccountId(req.AccountId)
	if err != nil {
		return nil, err
	}

	var transactions []*pb.TransactionJoinAccountResponse
	for _, tsc := range transactionsData {
		transactions = append(transactions, &pb.TransactionJoinAccountResponse{
			TransactionId:   tsc.ID,
			Name:            tsc.Name,
			Document:        tsc.Document,
			Description:     tsc.Description,
			Amount:          tsc.Amount,
			TransactionDate: timestamppb.New(tsc.TransactionDate),
		})
	}
	return &pb.TransactionJoinAccountListResponse{TransactionsByAccount: transactions}, nil
}

// Listar transactionAcount por document
func (s *Server) ListTransactionsByDocument(ctx context.Context, req *pb.Document) (*pb.TransactionJoinAccountListResponse, error) {
	transactionsData, err := s.TransactionUseCase.ListTransactionsByDocument(req.Document)
	if err != nil {
		return nil, err
	}

	var transactions []*pb.TransactionJoinAccountResponse
	for _, tsc := range transactionsData {
		transactions = append(transactions, &pb.TransactionJoinAccountResponse{
			TransactionId:   tsc.ID,
			Name:            tsc.Name,
			Document:        tsc.Document,
			Description:     tsc.Description,
			Amount:          tsc.Amount,
			TransactionDate: timestamppb.New(tsc.TransactionDate),
		})
	}
	return &pb.TransactionJoinAccountListResponse{TransactionsByAccount: transactions}, nil
}
