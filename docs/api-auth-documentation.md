# Documentação da API - Autenticação JWT

## Visão Geral

A API de gerenciamento de envios agora inclui autenticação JWT (JSON Web Token) para proteger as rotas. Todas as rotas de pacotes e usuários agora requerem autenticação, exceto as rotas de login e registro.

## Rotas de Autenticação

### 1. Registro de Usuário
**POST** `/auth/register`

Cria um novo usuário no sistema.

**Headers:**
- `Content-Type: application/json`

**Body:**
```json
{
  "name": "João Silva",
  "email": "joao@exemplo.com", 
  "password": "minhasenha123",
  "phone": "11999999999",
  "address": "Rua das Flores, 123"
}
```

**Exemplo com curl:**
```bash
curl -X POST http://localhost:8080/auth/register \
  -H "Content-Type: application/json" \
  -d '{
    "name": "João Silva",
    "email": "joao@exemplo.com",
    "password": "minhasenha123",
    "phone": "11999999999",
    "address": "Rua das Flores, 123"
  }'
```

**Resposta de Sucesso (201):**
```json
{
  "id": 1,
  "uuid": "550e8400-e29b-41d4-a716-446655440000",
  "name": "João Silva",
  "email": "joao@exemplo.com",
  "phone": "11999999999",
  "address": "Rua das Flores, 123",
  "created_at": "2025-01-07T10:30:00Z",
  "updated_at": "2025-01-07T10:30:00Z"
}
```

**Possíveis Erros:**
- `400 Bad Request`: Dados inválidos ou validação falhou
- `409 Conflict`: Email já está em uso
- `500 Internal Server Error`: Erro interno do servidor

---

### 2. Login de Usuário
**POST** `/auth/login`

Autentica um usuário e retorna um token JWT.

**Headers:**
- `Content-Type: application/json`

**Body:**
```json
{
  "email": "joao@exemplo.com",
  "password": "minhasenha123"
}
```

**Exemplo com curl:**
```bash
curl -X POST http://localhost:8080/auth/login \
  -H "Content-Type: application/json" \
  -d '{
    "email": "joao@exemplo.com",
    "password": "minhasenha123"
  }'
```

**Resposta de Sucesso (200):**
```json
{
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
  "user": {
    "id": 1,
    "uuid": "550e8400-e29b-41d4-a716-446655440000",
    "name": "João Silva",
    "email": "joao@exemplo.com",
    "phone": "11999999999",
    "address": "Rua das Flores, 123",
    "created_at": "2025-01-07T10:30:00Z",
    "updated_at": "2025-01-07T10:30:00Z"
  }
}
```

**Possíveis Erros:**
- `400 Bad Request`: Dados inválidos ou validação falhou
- `401 Unauthorized`: Email ou senha incorretos
- `500 Internal Server Error`: Erro interno do servidor

---

## Rotas Protegidas (Requerem Autenticação)

Para acessar as rotas protegidas, você deve incluir o token JWT no header `Authorization`:

**Header obrigatório:**
- `Authorization: Bearer <seu_token_jwt>`

### Rotas de Pacotes

#### 1. Criar Pacote
**POST** `/package`

**Exemplo com curl:**
```bash
curl -X POST http://localhost:8080/package \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxLCJ1c2VyX3V1aWQiOiJjMjAxZWY4ZS1mZTE0LTRlMTYtOTdhOS0yZmZkZGQ1ZTE1MjgiLCJlbWFpbCI6ImpvYW9AZXhlbXBsby5jb20iLCJpc3MiOiJzaGlwcGluZy1tYW5hZ2VtZW50IiwiZXhwIjoxNzUxNDk2MTE3LCJuYmYiOjE3NTE0MDk3MTcsImlhdCI6MTc1MTQwOTcxN30.2tMPAm4ZzvWpcwcZaFbawakC4RlTTQMuHNRA996dUAk" \
  -d '{
    "product": "Smartphone",
    "weight": 0.5,
    "destination": "SP"
  }'
```

#### 2. Buscar Pacote
**GET** `/package/{uuid}`

**Exemplo com curl:**
```bash
curl -X GET http://localhost:8080/package/550e8400-e29b-41d4-a716-446655440000 \
  -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
```

#### 3. Atualizar Status do Pacote
**PATCH** `/package/update/status`

**Exemplo com curl:**
```bash
curl -X PATCH http://localhost:8080/package/update/status \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..." \
  -d '{
    "uuid": "550e8400-e29b-41d4-a716-446655440000",
    "status": "DELIVERED"
  }'
```

#### 4. Contratar Transportadora
**POST** `/package/hire/carrier`

**Exemplo com curl:**
```bash
curl -X POST http://localhost:8080/package/hire/carrier \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..." \
  -d '{
    "package_uuid": "550e8400-e29b-41d4-a716-446655440000",
    "carrier_uuid": "650e8400-e29b-41d4-a716-446655440000"
  }'
```

#### 5. Obter Cotações
**GET** `/package/quotation/{uuid}`

**Exemplo com curl:**
```bash
curl -X GET http://localhost:8080/package/quotation/550e8400-e29b-41d4-a716-446655440000 \
  -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
```

### Rotas de Usuários

#### 1. Criar Usuário
**POST** `/users`

**Exemplo com curl:**
```bash
curl -X POST http://localhost:8080/users \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..." \
  -d '{
    "name": "Maria Santos",
    "email": "maria@exemplo.com",
    "phone": "11888888888",
    "address": "Av. Paulista, 1000"
  }'
```

#### 2. Listar Todos os Usuários
**GET** `/users`

**Exemplo com curl:**
```bash
curl -X GET http://localhost:8080/users \
  -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
```

#### 3. Buscar Usuário por UUID
**GET** `/users/{uuid}`

**Exemplo com curl:**
```bash
curl -X GET http://localhost:8080/users/550e8400-e29b-41d4-a716-446655440000 \
  -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
```

#### 4. Atualizar Usuário
**PUT** `/users/{uuid}`

**Exemplo com curl:**
```bash
curl -X PUT http://localhost:8080/users/550e8400-e29b-41d4-a716-446655440000 \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..." \
  -d '{
    "name": "João Silva Updated",
    "email": "joao.updated@exemplo.com",
    "phone": "11777777777",
    "address": "Nova Rua, 456"
  }'
```

#### 5. Deletar Usuário
**DELETE** `/users/{uuid}`

**Exemplo com curl:**
```bash
curl -X DELETE http://localhost:8080/users/550e8400-e29b-41d4-a716-446655440000 \
  -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
```

---

## Tratamento de Erros de Autenticação

### Possíveis Erros de Autenticação:

#### 401 Unauthorized - Token Ausente
```json
{
  "error": "Token de autorização é obrigatório"
}
```

#### 401 Unauthorized - Formato de Token Inválido
```json
{
  "error": "Formato de token inválido"
}
```

#### 401 Unauthorized - Token Inválido ou Expirado
```json
{
  "error": "Token inválido ou expirado"
}
```

---

## Configuração do Token JWT

- **Tempo de Expiração**: 24 horas
- **Algoritmo**: HS256
- **Emissor**: shipping-management

### Estrutura do Token JWT

O token JWT contém as seguintes informações no payload:

```json
{
  "user_id": 1,
  "user_uuid": "550e8400-e29b-41d4-a716-446655440000",
  "email": "joao@exemplo.com",
  "iss": "shipping-management",
  "exp": 1704629400,
  "iat": 1704543000,
  "nbf": 1704543000
}
```

---

## Fluxo de Autenticação Recomendado

1. **Registro**: Registre um novo usuário usando `/auth/register`
2. **Login**: Faça login com `/auth/login` para obter o token JWT
3. **Armazenar Token**: Guarde o token de forma segura (ex: localStorage, cookies httpOnly)
4. **Usar Token**: Inclua o token no header `Authorization` em todas as requisições para rotas protegidas
5. **Renovar Token**: Quando o token expirar (24h), faça login novamente para obter um novo token

---

## Exemplo de Uso Completo

```bash
# 1. Registrar usuário
TOKEN=$(curl -s -X POST http://localhost:8080/auth/register \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Test User",
    "email": "test@example.com",
    "password": "password123"
  }')

# 2. Fazer login e capturar o token
TOKEN=$(curl -s -X POST http://localhost:8080/auth/login \
  -H "Content-Type: application/json" \
  -d '{
    "email": "test@example.com",
    "password": "password123"
  }' | jq -r '.token')

# 3. Usar o token para criar um pacote
curl -X POST http://localhost:8080/package \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer $TOKEN" \
  -d '{
    "product": "Smartphone",
    "weight": 0.5,
    "destination": "São Paulo"
  }'
```

---

## Configuração de Segurança

### Variáveis de Ambiente

Certifique-se de configurar a variável de ambiente `JWT_SECRET` no arquivo `.env`:

```env
JWT_SECRET=your-super-secret-jwt-key-change-this-in-production
```

### Recomendações de Segurança

1. **Use HTTPS em produção** para proteger o token em trânsito
2. **Configure um JWT_SECRET forte** com pelo menos 32 caracteres aleatórios
3. **Implemente refresh tokens** para maior segurança
4. **Configure CORS adequadamente** para restringir origens
5. **Monitore tentativas de login inválidas** para detectar ataques
