# Go CRUD API

Uma API REST construída em **Go**, utilizando **Clean Architecture**, **PostgreSQL**, **Docker**, **Migrations** e uma arquitetura modular pensada para escalar.  
Hoje a API realiza operações CRUD de *Tasks*, mas o projeto está sendo estruturado para futuramente se tornar uma **Movie API**, onde usuários poderão marcar filmes assistidos, dar notas, e gerenciar uma coleção pessoal.


## Getting Started

### Prerequisites
- Go 1.18+
- Docker & Docker Compose
- PostgreSQL

### Environment Variables
Set the following environment variables (or use a `.env` file):
- `DB_HOST`: Database host
- `DB_PORT`: Database port
- `DB_USER`: Database user
- `DB_PASSWORD`: Database password
- `DB_NAME`: Database name

### Running with Docker Compose
```sh
docker compose up -d
```

### Running Locally
1. Export environment variables or create a `.env` file.
2. Start PostgreSQL.
3. Run the application:
   ```sh
   go run main.go
   ```

## Project Structure & Architecture

This project follows Clean Architecture principles, inspired by Domain Driven Design (DDD). As o projeto cresce, você pode organizar os arquivos em camadas como:

- **domain/**: entidades e regras de negócio
- **repository/**: interfaces e implementações para acesso a dados
- **usecase/**: casos de uso (aplicação)
- **config/**: configuração de banco de dados e ambiente
- **main.go**: ponto de entrada da aplicação

Exemplo de estrutura:
```
├── domain/
│   └── movie.go
├── repository/
│   └── movie_repository.go
├── usecase/
│   └── movie_usecase.go
├── config/
│   └── database.go
├── main.go
├── go.mod
├── docker-compose.yml
```

