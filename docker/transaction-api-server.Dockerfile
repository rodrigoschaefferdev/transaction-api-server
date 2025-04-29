# Usar a última versão estável do Golang como base
FROM golang:latest

# Definir o diretório de trabalho dentro do contêiner
WORKDIR /app

# Copiar os arquivos do projeto para o contêiner
COPY . .

# Navegar para o diretório do comando principal
WORKDIR /app/cmd/transaction-api-server

# Baixar dependências e compilar o binário
RUN go mod tidy && go build -o transaction-api-server


# Comando para executar a aplicação
CMD ["./transaction-api-server"]