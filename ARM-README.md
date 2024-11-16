# Setup on ARM-based Systems

This setup guide is a rough workaround inorder to get the project running on
ARM-based systems. At least some of the quality of life improvements that you
can find in the `*.sh`-files in the root of this project do not work with this
apporach.

All of the commands are run in your local enviourment. You have to have golang
installed on your enviourment. See the [golang-website](https://go.dev/) or your packagemanger
like [homebrew](https://brew.sh).

In order to start to run this project you need an `.env`-file. An example would
look something like this:
```env
# The mode used, options available
#  development, production
MODE=development

# The following levels are available:
#  trace, debug, info, warn, error, fatal, panic
LOG_LEVEL=debug

# The port on which the the service will listen on (default: 8080)
PORT=8081

# Database Related Configurations
POSTGRES_HOST=localhost
POSTGRES_PORT=5432
POSTGRES_USER=postgres
POSTGRES_PASSWORD=postgres
POSTGRES_DB=tasks-svc

# For development purposes
INSECURE_FAKE_TOKEN_ENABLE=true

# the provided oauth / oidc client id for this service
OAUTH_CLIENT_ID=425f8b8d-c786-4ff7-b2bf-e52f505fb588

FEATURE_FLAG_EVENT_SOURCING=true
EVENTSTORE_CS=esdb://admin:changeit@localhost:2113?tls=false
```

The service will listen on `localhost:8081` and expects `grpc` calls that can
be send via [postman](https://www.postman.com) or a cli like [grpc-client-cli](https://github.com/vadimi/grpc-client-cli.).

## Starting the Backend-Stack

```bash
docker compose up
```

## Manually creating the Postgres-DB (Wordaround)

> It should be possible to use `migrate.sh` or `migrate` directly but currently
> it does not work on my system. Therefore I am currently creating the
> tables manually

Using pgAdmin4 or psql run [these commands](./services/tasks-svc/migrations/001_init.down.sql) to create the tables in the database `tasks-svc`.

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
