# Transaction API Server 📊  

This is a **bank transaction API service**, developed in **Go**, using **gRPC** and **PostgreSQL**. It allows for account creation and transaction queries.  

## Technologies Used  
- **Go** 🐹  
- **gRPC** 🔌  
- **PostgreSQL** 🗄️  
- **Docker & Docker Compose** 🐳  

## Environment Setup  

Clone the repository:  
```bash
git clone 
cd transaction-api-server
```

Start the containers:
```bash
docker-compose up -d
```

This will start a PostgreSQL server (postgres) and the transaction API (transaction-api-server) running on port 50051.

Check the service logs:

```bash
docker-compose logs -f transaction-api-server
gRPC Endpoints
```

Create Account

```bash
grpcurl -plaintext -d '{ "name": "John", "document": "12345678900" }' \
  localhost:50051 AccountService.CreateAccount
```

List Accounts

```bash
grpcurl -plaintext localhost:50051 AccountService.ListAccounts
```

Get Account by ID

```bash
grpcurl -plaintext -d '{ "account_id": 1 }' \
  localhost:50051 AccountService.GetAccountById
```

Create Transaction - transaction_type_id = 1 (CASH_IN)/transaction_type_id = 2 (CASH_OUT)

```bash
grpcurl -plaintext -d '{ "account_id": 1, "transaction_type_id": 1, "amount": 1.45 }' \
  localhost:50051 TransactionService.CreateTransaction
```

List Transactions

```bash
grpcurl -plaintext localhost:50051 TransactionService.ListTransactions
```

List Transaction Amount by Document

```bash
grpcurl -plaintext -d '{ "document": "12345678900" }' \
  localhost:50051 TransactionService.ListTransactionAmount
```

List Transactions by Account ID

```bash
grpcurl -plaintext -d '{ "account_id": 1 }' \
  localhost:50051 TransactionService.ListTransactionsByAccountId
```

List Transactions by Document

```bash
grpcurl -plaintext -d '{ "document": "12345678900" }' \
  localhost:50051 TransactionService.ListTransactionsByDocument
```

Docker Compose Structure
Your environment uses Docker Compose to orchestrate the services:

PostgreSQL (postgres): Database running on port 5432.

Transaction API (transaction-api-server): gRPC service running on port 50051.

To stop the containers, run:

```bash
sh script\clean_up.sh
```