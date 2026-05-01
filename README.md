# Primeira API em Go

Uma API RESTful desenvolvida em Go como projeto de aprendizado, com gerenciamento de usuários em memória e integração com a API do Google Books.

---

## Tecnologias Utilizadas

| Tecnologia | Versão | Finalidade |
|---|---|---|
| [Go](https://golang.org/) | 1.26.1 | Linguagem principal |
| [Gorilla Mux](https://github.com/gorilla/mux) | v1.8.1 | Roteamento HTTP |
| [godotenv](https://github.com/joho/godotenv) | v1.5.1 | Variáveis de ambiente |
| [bcrypt](https://pkg.go.dev/golang.org/x/crypto/bcrypt) | v0.50.0 | Hash de senhas |

---

## Estrutura do Projeto

```
Primeira-api-golang-main/
├── main.go                      # Ponto de entrada da aplicação
├── go.mod                       # Dependências do módulo
├── go.sum                       # Checksums das dependências
│
├── router/
│   ├── router.go                # Cria e retorna o roteador principal (Gorilla Mux)
│   └── routes/
│       ├── routes.go            # Registra todas as rotas na aplicação
│       ├── user.go              # Rotas de usuários
│       └── books.go             # Rotas de livros
│
├── controller/
│   ├── user.go                  # Handlers CRUD de usuários
│   └── books_handler.go         # Handler para busca de livros (Google Books API)
│
└── models/
    └── users.go                 # Estrutura User + validação + formatação
```

---

## Como Rodar o Projeto

### Pré-requisitos

- [Go](https://golang.org/dl/) instalado (versão 1.21+)
- Acesso à internet (para o endpoint de livros)

### Passo a passo

```bash
# 1. Clone ou extraia o projeto
git clone <url-do-repositorio>
cd Primeira-api-golang-main

# 2. Instale as dependências
go mod tidy

# 3. Execute a aplicação
go run main.go
```

O servidor irá iniciar em **http://localhost:8080** e você verá no terminal:

```
2024/xx/xx xx:xx:xx Servidor rodando em http://localhost:8080
```

---

## Endpoints Disponíveis

###  Usuários

> **Atenção:** os dados são armazenados **em memória**. Ao reiniciar o servidor, todos os usuários são perdidos.

---

#### `POST /users` — Criar usuário

Cria um novo usuário. A senha é automaticamente convertida para hash com **bcrypt**.

**Body (JSON):**
```json
{
  "name": "João Silva",
  "email": "joao@email.com",
  "cpf": "12345678900",
  "senha": "minhasenha123"
}
```

**Resposta (201 Created):**
```json
{
  "id": 1,
  "name": "João Silva",
  "email": "joao@email.com",
  "cpf": "12345678900",
  "senha": ""
}
```
> 🔒 A senha **nunca** é retornada na resposta.

---

#### `GET /users` — Listar usuários

Retorna todos os usuários cadastrados (sem as senhas).

**Resposta (200 OK):**
```json
[
  {
    "id": 1,
    "name": "João Silva",
    "email": "joao@email.com",
    "cpf": "12345678900",
    "senha": ""
  }
]
```

---

#### `PUT /users/{userID}` — Atualizar usuário

Atualiza os dados de um usuário existente pelo ID.

**Parâmetro de URL:** `userID` (número inteiro)

**Body (JSON):**
```json
{
  "name": "João Atualizado",
  "email": "novo@email.com",
  "cpf": "12345678900",
  "senha": "novasenha456"
}
```

**Resposta (200 OK):** retorna o usuário atualizado (sem a senha).

**Erros possíveis:**
- `400 Bad Request` — ID inválido ou dados malformados
- `404 Not Found` — usuário não encontrado

---

#### `DELETE /users/{userID}` — Deletar usuário

Remove um usuário pelo ID.

**Parâmetro de URL:** `userID` (número inteiro)

**Resposta (204 No Content):** sem corpo na resposta.

**Erros possíveis:**
- `400 Bad Request` — ID inválido
- `404 Not Found` — usuário não encontrado

---

### Livros

#### `GET /books` — Buscar livros

Realiza uma busca na **Google Books API** (`https://www.googleapis.com/books/v1/volumes`) com o texto enviado no corpo da requisição.

> ⚠️ Este endpoint requer uma **chave de API** do Google Books passada via query string, **direto pelo Postman**.

**Como obter uma chave:**
1. Acesse o Google Cloud Console
2. Crie um projeto e ative a **Books API**
3. Vá em **APIs & Services → Credentials → Criar credenciais → Chave de API**

**URL com a chave:**
```
GET http://localhost:8080/books?key=SUA_CHAVE_AQUI
```

**Body (texto puro — raw/Text no Postman):**
```
harry potter
```

**Resposta (200 OK):**
```json
{
  "kind": "books#volumes",
  "totalItems": 1024,
  "items": [
    {
      "id": "abc123",
      "volumeInfo": {
        "title": "Harry Potter and the Sorcerer's Stone",
        "authors": ["J.K. Rowling"],
        "publishedDate": "1997"
      }
    }
  ]
}
```

>  O endpoint usa um timeout de **10 segundos** para a requisição à API do Google.

---

## Segurança

- As senhas dos usuários são armazenadas com hash **bcrypt** (custo padrão) — nunca em texto puro.
- A senha é removida de **todas** as respostas da API, tanto no create quanto no fetch.
- A validação de e-mail verifica a presença do caractere `@`.

---

##  Validações do Modelo de Usuário

Ao criar ou atualizar um usuário, as seguintes regras são aplicadas:

| Campo | Regra |
|---|---|
| `name` | Obrigatório, não pode ser vazio |
| `email` | Obrigatório, deve conter `@` |
| `cpf` | Obrigatório, não pode ser vazio |
| `senha` | Obrigatório apenas na **criação** |

Após a validação, os campos de texto são tratados com `strings.TrimSpace()` para remover espaços desnecessários.

---

##  Como o Projeto Está Organizado (Arquitetura)

O projeto segue uma separação de responsabilidades em camadas:

```
Requisição HTTP
     │
     ▼
  Router (gorilla/mux)
     │  Direciona para o handler correto
     ▼
  Controller
     │  Lê o body, chama o model, monta a resposta
     ▼
  Model (models/users.go)
     │  Valida e formata os dados (regras de negócio)
     ▼
  "Banco" em Memória (slice []User)
```

- **`router/`** — responsável apenas por criar o roteador e registrar as rotas.
- **`controller/`** — responsável por receber a requisição HTTP, processar a lógica e retornar a resposta.
- **`models/`** — contém as structs e as regras de negócio (validação, formatação, hash de senha).

---

##  Observações e Limitações

- **Sem banco de dados:** todos os dados ficam em um slice Go (`[]models.User`) na memória. Ao reiniciar o servidor, os dados são perdidos.
- **Sem autenticação:** os endpoints não exigem token ou sessão.
- **Sem paginação:** o `GET /users` retorna todos os usuários de uma vez.
- **Busca de livros via body:** o endpoint `GET /books` recebe a query no corpo da requisição (texto puro), o que é incomum para métodos GET — o padrão seria usar query string (`?q=golang`).

---

## Possíveis Melhorias Futuras

- [ ] Integrar com banco de dados (PostgreSQL, SQLite, etc.)
- [ ] Adicionar autenticação com JWT
- [ ] Usar query string no endpoint de livros (`GET /books?q=golang`)
- [ ] Adicionar middleware de logging
- [ ] Implementar paginação no `GET /users`
- [ ] Adicionar testes unitários

---

## Licença

Projeto desenvolvido para fins educacionais. Sinta-se livre para estudar e modificar.
