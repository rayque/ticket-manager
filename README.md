# Sistema de Gerenciamento de Remessas (Shipment Management)

Este projeto consiste em uma API para gerenciamento de remessas, permitindo a criação de pacotes, cotação de entregas, contratação de transportadoras e atualização de status de entrega.

## Estrutura do Projeto

O projeto segue uma arquitetura limpa (Clean Architecture) e está organizado da seguinte forma:

```
shipment-management/
├── cmd/                    # Pontos de entrada da aplicação
│   ├── api/                # API principal
│   └── seed/               # Utilitário para seed do banco de dados
├── docs/                   # Documentação (OpenAPI e coleção Postman)
├── internal/               # Código interno da aplicação
│   ├── application/        # Regras de aplicação (casos de uso)
│   ├── domain/             # Regras de negócio (entidades e interfaces)
│   └── infrastructure/     # Implementações concretas (adaptadores, repositórios)
├── mocks/                  # Mocks para testes
└── scripts/                # Scripts utilitários
```

## Pré-requisitos

Para executar este projeto localmente, você precisará:

- Docker e Docker Compose

## Configuração e Execução Local

### 1. Clone o repositório

```bash
git clone https://github.com/rayque/shipment-management.git
cd shipment-management
```

### 2. Iniciar o ambiente com Docker Compose

```bash
make install
```

Este comando irá iniciar:
- PostgreSQL - banco de dados principal
- MongoDB - banco de dados para trasnportadoras
- API do sistema - em modo de desenvolvimento
- Execução da seed de trasportadoras

A API estará disponível em `http://localhost:8080`

### 3. Executar testes

```bash
# Executar todos os testes
make test
```

#### Executar testes com cobertura

```bash
# Executar todos os testes
make test-command
```

## Utilizando a API

A API oferece os seguintes endpoints principais:

- `POST /package` - Criar um novo pacote
- `GET /package/{uuid}` - Obter detalhes de um pacote
- `GET /package/quotation/{quotation}` - Obter cotação para entrega de um pacote
- `POST /package/hire/carrier` - Contratar transportadora para entrega
- `PUT /package/update/status` - Atualizar status de entrega

Para detalhes completos da API, consulte a documentação OpenAPI em `docs/open-api.yml` ou importe a coleção Postman disponível em `docs/shipment-management.postman_collection.json`.

## Usando o Makefile

O projeto inclui um Makefile para facilitar operações comuns:

```bash
# Iniciar todos os serviços
make up

# Parar todos os serviços
make down

# Executar testes
make test

# Gerar mocks
make generate-mocks
```

## Desenvolvimento

### Gerar Mocks

Para gerar/atualizar os mocks para testes:

```bash
./scripts/generate-mocks.sh
```

### Adicionar Novas Migrações

Para adicionar uma nova migração:

```bash
# Criar arquivo de migração (substitua MIGRATION_NAME pelo nome descritivo)
migrate create -ext sql -dir internal/infrastructure/database/migrations -seq MIGRATION_NAME
```

