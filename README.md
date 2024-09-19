# Books API

## Como rodar localmente

1. Clone o repositório.
2. Configure as variáveis de ambiente no arquivo `.env`.
3. Rode `go run main.go`.

## Como fazer o deploy no Render

1. Crie uma conta no [Render](https://render.com).
2. Configure o banco de dados PostgreSQL no Render.
3. Adicione as variáveis de ambiente (DB_HOST, DB_USER, DB_PASSWORD, DB_NAME, DB_PORT).
4. Conecte o repositório e clique em "Deploy".

## Estrutura da Aplicação

go-books-api/
│
├── config/
│   └── config.go
├── controllers/
│   └── book_controller.go
├── models/
│   └── book.go
├── routes/
│   └── routes.go
├── main.go
├── go.mod
├── go.sum
├── .env
└── swagger.yaml