# Making Bank

A backend focused web service designed with inspiration from the [Backend Master Class](https://www.udemy.com/course/backend-master-class-golang-postgresql-kubernetes/) Course on Udemy. Covers design, development, testing, and deployment of a real-world back-end service.

### API Capabilities

- Create and manage bank accounts
- Record balance changes to each of the accounts
- Perform money transfer between accounts

## Requirements

- [GO](https://go.dev/) version 1.18 or higher
- [Docker](<[Docker](https://www.docker.com/)>)
- [Golang Migrate](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate)
- [Sqlc](https://github.com/kyleconroy/sqlc#installation)
- [Gomock](https://github.com/golang/mock)

## Installation and set up

- Clone this repository
  ```
  git clone https://github.com/Smelton01/bank
  cd bank
  ```
- Start Postgres container
  ```
  make postgres
  ```
- Create bank database
  ```
  make createdb
  ```
- Run database migration up
  ```
  make migrateup
  ```
- Generate SQL Crud with sqlc
  ```
  make generate
  ```

## Running service

- Run server
  ```
  make server
  ```
- Run tests
  ```
  make test
  ```

## Deploy to kubernetes cluster

> ##### TODO

## License

This project is open source and available under the [MIT License](LICENSE)
