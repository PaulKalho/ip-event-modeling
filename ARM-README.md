# Setup on ARM-based Systems

## Starting the Backend-Stack

```bash
docker compose up
```

## Manually creating the Postgres-DB (Wordaround)

> It should be possible to use `migrate.sh` or `migrate` directly but currently
> it does not work on my system. Therefore I am currently creating the
> tables manually

### Using `migrate`/`migrate.sh` 

Install `migrate` with Homebrew or [use other installation methods](https://github.com/golang-migrate/migrate/blob/master/cmd/migrate/README.md).

```bash
brew install golang-migrate
```

> [!IMPORTANT]
> The following steps to set up the postgres tables do not work for me currently

Using `migrate.sh` 

```bash
bash ./migrate.sh
```

Or directly though the cli.

```bash
migrate -verbose -path services/tasks-svc/migrations -database "postgres://postgres:postgres@localhost:5432/tasks-svc?sslmode=disable" up
```

### Starting the `task-svc` service

```bash
cd services/task-svc
go run main.go
```
