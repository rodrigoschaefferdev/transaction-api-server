# Transaction API Server 📊  

This is a **bank transaction API service**, developed in **Go**, using **gRPC** and **PostgreSQL**. It allows for account creation and transaction queries.  

## Technologies Used  
- **Go** 🐹  
- **gRPC** 🔌  
- **PostgreSQL** 🗄️  
- **Docker & Docker Compose** 🐳  

## Environment Setup  

Clone the repository:  
```sh
git clone 
cd transaction-api-server
Start the containers:

sh
docker-compose up -d
This will start a PostgreSQL server (postgres) and the transaction API (transaction-api-server) running on port 50051.

Check the service logs:

sh
docker-compose logs -f transaction-api-server
gRPC Endpoints
Create Account
Request:

json
{
  "name": "John",
  "document": "12345678900"
}
Response:

json
{
  "account_id": "1"
}
List Accounts
Response:

json
{
  "accounts": [
    {
      "account_id": "1",
      "name": "John",
      "document": "12345678900",
      "created_at": {
        "seconds": "-62135596800",
        "nanos": 0
      },
      "updated_at": {
        "seconds": "-62135596800",
        "nanos": 0
      }
    }
  ]
}
Get Account by ID
Request:

json
{ "account_id": 1 }
Response:

json
{
  "account_id": "1",
  "name": "John",
  "document": "12345678900",
  "created_at": {
    "seconds": "-62135596800",
    "nanos": 0
  },
  "updated_at": {
    "seconds": "-62135596800",
    "nanos": 0
  }
}
Create Transaction - transaction_type_id = 1 (CASH_IN)/transaction_type_id = 2 (CASH_OUT)
Request:

json
{
  "account_id": 1,
  "transaction_type_id": 1,
  "amount": 1.45
}
Response:

json
{
  "transaction_id": "1"
}
List Transactions
Response:

json
{
  "transactions": [
    {
      "transaction_id": "1",
      "account_id": "1",
      "transaction_type_id": 1,
      "amount": 1.45,
      "transaction_date": {
        "seconds": "1745630561",
        "nanos": 535641000
      }
    }
  ]
}
List Transaction Amount by Document
Request:

json
{
  "document": "12345678900"
}
Response:

json
{
  "transactionsAmount": [
    {
      "name": "John",
      "document": "12345678900",
      "amount": -1.45
    }
  ]
}
List Transactions by Account ID
Request:

json
{
  "account_id": 1
}
Response:

json
{
  "transactionsByAccount": [
    {
      "transaction_id": "1",
      "name": "John",
      "document": "12345678900",
      "description": "CASH_IN",
      "amount": 1.45,
      "transaction_date": {
        "seconds": "1745630561",
        "nanos": 535641000
      }
    }
  ]
}
Docker Compose Structure
Your environment uses Docker Compose to orchestrate the services:

PostgreSQL (postgres): Database running on port 5432.

Transaction API (transaction-api-server): gRPC service running on port 50051.

To stop the containers, run:

sh
docker-compose down