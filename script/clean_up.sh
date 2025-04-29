#!/bin/bash

# Verifica se Docker está rodando antes de executar comandos
if ! docker info > /dev/null 2>&1; then
    echo "Erro: Docker Engine não está rodando. Inicie o Docker primeiro!"
    exit 1
fi

# Lista containers das aplicações
echo "Listando containers de postgres e transaction-api-server..."
docker ps -a --filter "name=postgres"
docker ps -a --filter "name=transaction-api-server"

# Para os containers específicos
echo "Parando containers..."
docker stop $(docker ps -q --filter "name=postgres") 2>/dev/null
docker stop $(docker ps -q --filter "name=transaction-api-server") 2>/dev/null

# Remove os containers específicos
echo "Removendo containers..."
docker rm $(docker ps -aq --filter "name=postgres") 2>/dev/null
docker rm $(docker ps -aq --filter "name=transaction-api-server") 2>/dev/null

# Remove volumes das aplicações
echo "Removendo volumes..."
docker volume prune -f

echo "Processo concluído!"