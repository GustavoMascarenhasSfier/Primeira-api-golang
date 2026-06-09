# 📚 Treehouse Library API

API REST desenvolvida em Go para gerenciamento de uma biblioteca pessoal. Cada usuário pode criar sua conta, fazer login e montar sua própria coleção de livros — pesquisando pelo Google Books ou cadastrando manualmente.

---

## 🛠️ Tecnologias

- **Go** — linguagem principal
- **Gorilla Mux** — roteador HTTP
- **MySQL** — banco de dados
- **JWT** — autenticação via token
- **Google Books API** — pesquisa de livros
- **bcrypt** — hash de senhas

---

## 📁 Estrutura do Projeto

```
Primeira-api-golang/
├── auth/           # Geração e validação de JWT
├── config/         # Carregamento do .env e configuração do banco
├── controller/     # Handlers HTTP (lógica de cada rota)
├── middlewares/    # Logger e autenticação
├── models/         # Structs de User e Book
├── persistency/    # Conexão com o banco de dados
├── repository/     # Queries SQL (User e Book)
├── responses/      # Helpers para resposta JSON
├── router/         # Registro de rotas
├── security/       # Hash e validação de senha
├── sql/            # DDL do banco de dados
├── utils/          # Validador de CPF
└── main.go
```

---

## ⚙️ Como rodar o projeto

### 1. Pré-requisitos

- Go instalado (1.21+)
- MySQL rodando localmente
- Git

### 2. Clonar o repositório

```bash
git clone <url-do-repositorio>
cd Primeira-api-golang
```

### 3. Criar o banco de dados

Abra seu cliente MySQL e execute o arquivo:

```bash
mysql -u root -p < sql/ddl.sql
```

Isso vai criar o banco `treehousedb` com as tabelas `users` e `books`.

### 4. Configurar o `.env`

Edite o arquivo `config/.env` com suas credenciais:

```env
DB_USER=root
DB_PASSWORD=sua_senha
DB_ADDR=localhost:3306
DB_DATABASE=treehousedb

API_PORT=:8080

SECRET_KEY=sua_chave_secreta
```

### 5. Instalar dependências

```bash
go mod tidy
```

### 6. Rodar a API

```bash
go run main.go
```

A API estará disponível em `http://localhost:8080`.

---

## 🗄️ Banco de Dados

### Tabela `users`

| Coluna     | Tipo         | Descrição              |
|------------|--------------|------------------------|
| id         | INT PK AI    | Identificador único    |
| name       | VARCHAR(100) | Nome do usuário        |
| cpf        | VARCHAR(14)  | CPF (único)            |
| email      | VARCHAR(100) | E-mail (único)         |
| password   | VARCHAR(255) | Senha em bcrypt        |

### Tabela `books`

| Coluna      | Tipo         | Descrição                        |
|-------------|--------------|----------------------------------|
| id          | INT PK AI    | Identificador único              |
| user_id     | INT FK       | Referência ao dono do livro      |
| title       | VARCHAR(255) | Título do livro                  |
| author      | VARCHAR(255) | Autor                            |
| description | TEXT         | Descrição (opcional)             |
| publisher   | VARCHAR(255) | Editora (opcional)               |
| year        | INT          | Ano de publicação (opcional)     |
| created_at  | TIMESTAMP    | Data de adição à biblioteca      |

> `books.user_id` tem `FOREIGN KEY` para `users.id` com `ON DELETE CASCADE` — ao deletar um usuário, todos os seus livros são removidos automaticamente.

---

## 🔐 Autenticação

A API usa **JWT (JSON Web Token)**. Após o login, você recebe um token que deve ser enviado no header de todas as rotas protegidas:

```
Authorization: Bearer SEU_TOKEN_AQUI
```

O token expira em **6 horas**.

---

## 🚀 Rotas

### Resumo

| Método | Rota               | Auth | Descrição                        |
|--------|--------------------|------|----------------------------------|
| POST   | /users             | ❌   | Criar usuário                    |
| GET    | /users             | ✅   | Listar todos os usuários         |
| GET    | /users/{userID}    | ✅   | Buscar usuário por ID            |
| POST   | /login             | ❌   | Login e geração de token         |
| POST   | /books/search      | ❌   | Pesquisar livros no Google Books |
| POST   | /library           | ✅   | Adicionar livro à biblioteca     |
| GET    | /library           | ✅   | Ver toda a sua biblioteca        |
| GET    | /library/{bookID}  | ✅   | Ver um livro específico          |
| PUT    | /library/{bookID}  | ✅   | Editar um livro                  |
| DELETE | /library/{bookID}  | ✅   | Remover um livro                 |

---

## 📬 Testando no Postman

### Dica: salvar o token automaticamente

Na requisição de **Login**, vá na aba **Tests** e cole:

```javascript
pm.environment.set("token", pm.response.text());
```

Crie um **Environment** no Postman e em todas as rotas autenticadas use no header:

```
Authorization: Bearer {{token}}
```

---

### 👤 Usuários

#### Criar usuário
```
POST http://localhost:8080/users
```
Body → raw → JSON:
```json
{
  "name": "João Silva",
  "email": "joao@email.com",
  "cpf": "529.982.247-25",
  "senha": "123456"
}
```
Resposta `201`:
```json
{
  "id": 1,
  "name": "joão silva",
  "cpf": "529.982.247-25",
  "email": "joao@email.com",
  "senha": "$2a$10$..."
}
```

---

#### Listar todos os usuários
```
GET http://localhost:8080/users
```
Header:
```
Authorization: Bearer {{token}}
```
Resposta `200`:
```json
[
  {
    "id": 1,
    "name": "joão silva",
    "cpf": "529.982.247-25",
    "email": "joao@email.com",
    "senha": ""
  }
]
```

---

#### Buscar usuário por ID
```
GET http://localhost:8080/users/1
```
Header:
```
Authorization: Bearer {{token}}
```
Resposta `200`:
```json
{
  "id": 1,
  "name": "joão silva",
  "cpf": "529.982.247-25",
  "email": "joao@email.com",
  "senha": ""
}
```

---

### 🔑 Login

#### Fazer login
```
POST http://localhost:8080/login
```
Body → raw → JSON:
```json
{
  "email": "joao@email.com",
  "senha": "123456"
}
```
Resposta `200`:
```
eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...
```
> O corpo da resposta é o token JWT em texto puro. Salve-o para usar nas próximas requisições.

---

### 📖 Pesquisa de Livros

#### Pesquisar no Google Books
```
POST http://localhost:8080/books/search?key=SUA_GOOGLE_API_KEY
```
Body → raw → **Text**:
```
Harry Potter
```
Resposta `200`: JSON completo da Google Books API com lista de livros encontrados.

> A chave da Google Books API é opcional — sem ela funciona, mas com limite menor de requisições. Gere a sua em: https://console.cloud.google.com/apis/credentials

---

### 📚 Biblioteca Pessoal

> Todas as rotas abaixo exigem o header:
> ```
> Authorization: Bearer {{token}}
> ```
> Os livros são sempre vinculados ao usuário dono do token — não é possível ver ou mexer na biblioteca de outro usuário.

---

#### Adicionar livro à biblioteca
```
POST http://localhost:8080/library
```
Body → raw → JSON:
```json
{
  "title": "Harry Potter e a Pedra Filosofal",
  "author": "J.K. Rowling",
  "description": "Primeiro livro da saga Harry Potter",
  "publisher": "Rocco",
  "year": 1997
}
```
> `title` e `author` são obrigatórios. Os demais campos são opcionais.

Resposta `201`:
```json
{
  "id": 1,
  "user_id": 1,
  "title": "Harry Potter e a Pedra Filosofal",
  "author": "J.K. Rowling",
  "description": "Primeiro livro da saga Harry Potter",
  "publisher": "Rocco",
  "year": 1997,
  "created_at": "2026-06-08T10:00:00Z"
}
```

---

#### Ver toda a biblioteca
```
GET http://localhost:8080/library
```
Resposta `200`:
```json
[
  {
    "id": 1,
    "user_id": 1,
    "title": "Harry Potter e a Pedra Filosofal",
    "author": "J.K. Rowling",
    "description": "Primeiro livro da saga Harry Potter",
    "publisher": "Rocco",
    "year": 1997,
    "created_at": "2026-06-08T10:00:00Z"
  }
]
```

---

#### Ver um livro específico
```
GET http://localhost:8080/library/1
```
Resposta `200`:
```json
{
  "id": 1,
  "user_id": 1,
  "title": "Harry Potter e a Pedra Filosofal",
  "author": "J.K. Rowling",
  "description": "Primeiro livro da saga Harry Potter",
  "publisher": "Rocco",
  "year": 1997,
  "created_at": "2026-06-08T10:00:00Z"
}
```

---

#### Editar um livro
```
PUT http://localhost:8080/library/1
```
Body → raw → JSON:
```json
{
  "title": "Harry Potter e a Pedra Filosofal",
  "author": "J.K. Rowling",
  "description": "Edição comemorativa 20 anos",
  "publisher": "Rocco",
  "year": 2017
}
```
Resposta `204` (sem body).

---

#### Remover um livro
```
DELETE http://localhost:8080/library/1
```
Resposta `204` (sem body).

---

## ❌ Respostas de erro

Todos os erros seguem o formato:

```json
{
  "error": "mensagem do erro"
}
```

| Status | Situação                                 |
|--------|------------------------------------------|
| 400    | Body inválido ou campo faltando          |
| 401    | Token ausente, inválido ou expirado      |
| 404    | Recurso não encontrado                   |
| 422    | Erro ao ler o body da requisição         |
| 500    | Erro interno do servidor                 |

---

## 🔄 Fluxo completo de uso

```
1. POST /users          → cria sua conta
2. POST /login          → recebe o token JWT
3. POST /books/search   → pesquisa livros no Google
4. POST /library        → adiciona um livro à sua biblioteca
5. GET  /library        → vê todos os seus livros
6. PUT  /library/{id}   → edita um livro
7. DELETE /library/{id} → remove um livro
```
